package execute

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yecaowulei/sre-tool-kit/internal/common"
	"log"
	"os"
	"strings"
)

type MysqlData struct {
	DbHost      string
	DbPort      int
	DbName      string
	DbUser      string
	DbPassword  string
	SqlFilename string
	SqlType     string
}

func queryDatabase(db *sql.DB, query string) ([][]interface{}, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var results [][]interface{}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	numColumns := len(columns)

	for rows.Next() {
		// 创建一个interface{}切片来存储列的值
		values := make([]interface{}, numColumns)
		valuePtrs := make([]interface{}, numColumns)
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// Scan将行的列值复制到传递的切片
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		// 将这一行的值追加到结果切片
		results = append(results, values)
	}

	return results, nil
}

func (d MysqlData) Run() {
	if d.DbHost == "" || d.DbUser == "" || d.DbPassword == "" || d.SqlFilename == "" || d.SqlType == "" {
		fmt.Println("入参异常")
		return
	}

	if _, err := os.Stat(d.SqlFilename); os.IsNotExist(err) {
		log.Fatalf("SQL文件%s不存在", d.SqlFilename)
	}

	var dbMsg string
	if d.DbName == "" {
		dbMsg = ""
	} else {
		dbMsg = fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", d.DbName)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.DbUser, d.DbPassword, d.DbHost, d.DbPort, dbMsg)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("无法连接到数据库: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	statements, err := common.GetFileContents(d.SqlFilename)
	if err != nil {
		log.Fatal(err)
	}

	if d.SqlType == "select" {
		stmt := statements[0]
		//fmt.Printf("准备执行的SQL语句如下：\n%s;\n", stmt)
		results, err := queryDatabase(db, stmt)
		if err != nil {
			log.Fatal(err)
		}

		for _, row := range results {
			for _, value := range row {
				fmt.Printf("%v  ", value)
			}
			fmt.Println()
		}
	} else {
		tx, err := db.Begin()
		if err != nil {
			log.Fatalf("无法开始事务: %v", err)
		}

		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}
			fmt.Printf("准备执行的SQL语句如下：\n%s\n", stmt)
			_, err := tx.Exec(stmt)
			if err != nil {
				log.Printf("SQL语句执行失败：%v\n", err)
				err := tx.Rollback()
				if err != nil {
					log.Fatalf("执行SQL语句失败后回滚事务失败，请人工介入检查: %v", err)
				}
				log.Fatalf("事务回滚成功")
			} else {
				log.Println("SQL语句执行成功")
			}
		}

		if err := tx.Commit(); err != nil {
			log.Fatalf("无法提交事务: %v", err)
		}

		log.Printf("SQL脚本%s成功执行并提交", d.SqlFilename)
	}
}
