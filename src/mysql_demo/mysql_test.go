package mysql_demo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"testing"
)

// 在 Go 中连接 MySQL 数据库，可以使用 github.com/go-sql-driver/mysql 包，它是官方推荐的 MySQL 驱动。
// 导入驱动的方法  _ "github.com/go-sql-driver/mysql"
func TestMysql01(t *testing.T) {
	// 创建数据库连接字符串
	// 格式: username:password@protocol(address)/mysql?param=value
	dsn := "root:root@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4"

	// Go使用SQL与类SQL数据库的惯例是通过标准库database/sql。这是一个对关系型数据库的通用抽象，它提供了标准的、轻量的、面向行的接口。
	// Go本身不提供具体数据库驱动，只提供驱动接口和管理，要使用数据库，除了database/sql包本身，还需要引入想使用的特定数据库驱动

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal().Msgf("Error opening database: %v", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal().Msgf("Error closing database: %v", err)
		}
	}(db)

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		log.Fatal().Msgf("Error pinging database: %v", err)
	}
	log.Info().Msgf("Successfully connected to the database!")

	// 执行查询
	rows, err := db.Query("SELECT Host, User FROM user")
	if err != nil {
		log.Fatal().Msgf("Error executing query: %v", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal().Msgf("Error closing rows: %v", err)
		}
	}(rows)

	// 读取查询结果
	for rows.Next() {
		var host string
		var user string
		if err := rows.Scan(&host, &user); err != nil {
			log.Fatal().Msgf("Error scanning row: %v", err)
		}
		log.Info().Msgf("User: %s, %s\n", host, user)
	}

	// 检查查询中的错误
	if err := rows.Err(); err != nil {
		log.Fatal().Msgf("Error iterating rows: %v", err)
	}
}
