package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"testing"
	"time"
)

// ServiceRegister 服务注册
type ServiceRegister struct {
	cli     *clientv3.Client
	leaseID clientv3.LeaseID
	// 租约keepalive对应的chan
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
	key           string
	val           string
}

// NewServiceRegister 新建注册服务
func NewServiceRegister(endpoints []string, key, val string, lease int64) (*ServiceRegister, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	ser := &ServiceRegister{
		cli: cli,
		key: key,
		val: val,
	}

	// 申请租约设置时间keepalive
	if err := ser.putKeyWithLease(lease); err != nil {
		return nil, err
	}

	return ser, nil
}

// putKeyWithLease 设置key和租约
func (s *ServiceRegister) putKeyWithLease(lease int64) error {
	// 设置租约时间
	resp, err := s.cli.Grant(context.Background(), lease)
	if err != nil {
		return err
	}
	// 注册服务并绑定租约
	_, err = s.cli.Put(context.Background(), s.key, s.val, clientv3.WithLease(resp.ID))
	if err != nil {
		return err
	}
	// 设置续租 定期发送需求请求
	leaseRespChan, err := s.cli.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		return err
	}

	s.leaseID = resp.ID
	s.keepAliveChan = leaseRespChan
	log.Printf("Put key:%s  val:%s  success!", s.key, s.val)
	return nil
}

// ListenLeaseRespChan 监听续租情况
func (s *ServiceRegister) ListenLeaseRespChan() {
	for leaseKeepResp := range s.keepAliveChan {
		log.Println("续租成功", leaseKeepResp)
	}
	log.Println("关闭续租")
}

// Close 注销服务
func (s *ServiceRegister) Close() error {
	// 撤销租约
	if _, err := s.cli.Revoke(context.Background(), s.leaseID); err != nil {
		return err
	}
	log.Println("撤销租约")
	return s.cli.Close()
}

func TestRegister01(t *testing.T) {
	var endpoints = []string{"localhost:2379"}
	ser, err := NewServiceRegister(endpoints, "/web/node1", "localhost:8000", 5)
	if err != nil {
		log.Fatalln(err)
	}
	// 监听续租相应chan
	go ser.ListenLeaseRespChan()

	select {
	// case <-time.After(20 * time.Second):
	// 	ser.Close()
	}
}
