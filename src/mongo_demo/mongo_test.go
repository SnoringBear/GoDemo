package mongo_demo

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 定义一个示例结构体
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Age      int                `bson:"age"`
	Email    string             `bson:"email"`
	CreateAt time.Time          `bson:"createAt"`
}

func TestMongo01(t *testing.T) {
	// 1. 连接MongoDB
	client, err := connectMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, context.TODO())

	// 获取数据库和集合
	db := client.Database("testdb")
	collection := db.Collection("users")

	// 2. 插入文档
	insertedID, err := insertDocument(collection)
	if err != nil {
		log.Println("插入失败:", err)
	} else {
		log.Println("插入成功，ID:", insertedID)
	}

	// 3. 查询文档
	err = findDocuments(collection)
	if err != nil {
		log.Println("查询失败:", err)
	}

	// 4. 更新文档
	err = updateDocument(collection, insertedID)
	if err != nil {
		log.Println("更新失败:", err)
	}

	// 5. 删除文档
	err = deleteDocument(collection, insertedID)
	if err != nil {
		log.Println("删除失败:", err)
	}
}

// 连接MongoDB
func connectMongoDB() (*mongo.Client, error) {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://150.158.46.207:27018")
	clientOptions.SetAuth(options.Credential{
		Username: "root",
		Password: "}aAW*Fd,G#TQu%2KUc~_",
	})

	// 连接到MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("成功连接到MongoDB")
	return client, nil
}

// 插入文档
func insertDocument(collection *mongo.Collection) (primitive.ObjectID, error) {
	// 创建一个用户
	user := User{
		Name:     "张三",
		Age:      30,
		Email:    "zhangsan@example.com",
		CreateAt: time.Now(),
	}

	// 插入单个文档
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	// 获取插入的ID
	insertedID := result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

// 批量插入文档
func insertManyDocuments(collection *mongo.Collection) ([]primitive.ObjectID, error) {
	// 创建多个用户
	users := []interface{}{
		User{Name: "李四", Age: 25, Email: "lisi@example.com", CreateAt: time.Now()},
		User{Name: "王五", Age: 35, Email: "wangwu@example.com", CreateAt: time.Now()},
	}

	// 批量插入文档
	result, err := collection.InsertMany(context.TODO(), users)
	if err != nil {
		return nil, err
	}

	// 转换为ObjectID切片
	var ids []primitive.ObjectID
	for _, id := range result.InsertedIDs {
		ids = append(ids, id.(primitive.ObjectID))
	}

	return ids, nil
}

// 查询文档
func findDocuments(collection *mongo.Collection) error {
	// 1. 查询单个文档
	var user User
	err := collection.FindOne(context.TODO(), bson.D{{"name", "张三"}}).Decode(&user)
	if err != nil {
		return err
	}
	fmt.Println("查询到的单个文档:", user)

	// 2. 查询多个文档
	// 设置查询条件
	filter := bson.D{{"age", bson.D{{"$gt", 20}}}} // 年龄大于20的

	// 执行查询
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cursor, context.TODO())

	// 遍历结果
	var users []User
	if err = cursor.All(context.TODO(), &users); err != nil {
		return err
	}

	fmt.Println("查询到的多个文档:")
	for _, u := range users {
		fmt.Println(u)
	}

	return nil
}

// 更新文档
func updateDocument(collection *mongo.Collection, id primitive.ObjectID) error {
	// 设置查询条件，根据ID查询
	filter := bson.D{{"_id", id}}

	// 设置更新内容
	update := bson.D{
		{"$set", bson.D{
			{"age", 31},
			{"email", "updated@example.com"},
		}},
	}

	// 执行更新
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	fmt.Printf("更新了 %v 个文档\n", result.ModifiedCount)
	return nil
}

// 删除文档
func deleteDocument(collection *mongo.Collection, id primitive.ObjectID) error {
	// 设置删除条件
	filter := bson.D{{"_id", id}}
	// bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
	// bson.A{"bar", "world", 3.14159, bson.D{{"qux", 12345}}}
	// bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}

	// 执行删除
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	fmt.Printf("删除了 %v 个文档\n", result.DeletedCount)
	return nil
}
