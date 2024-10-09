package check

import (
	"github.com/yecaowulei/sre-tool-kit/internal/common"
	"log"
	"os"
	"regexp"
	"strings"
)

type MysqlData struct {
	SqlFilename string
}

func checkStatementComment(stmt string) bool {
	lines := strings.Split(stmt, "\n")
	hasComment := true
	fieldRegex := regexp.MustCompile(`^` + "`" + `[A-Z]+` + "`" + `.*,$`)

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		// 检查字段定义行（以逗号结尾）是否包含COMMENT
		if strings.HasSuffix(trimmedLine, ";") {
			if !strings.Contains(trimmedLine, "COMMENT") {
				hasComment = false
				log.Println("存在建表定义行没有COMMENT注释：", trimmedLine)
			}
		} else if fieldRegex.MatchString(trimmedLine) {
			if !strings.Contains(trimmedLine, "COMMENT") {
				hasComment = false
				log.Println("存在字段定义行没有COMMENT注释：", trimmedLine)
				break
			}
		}
	}

	return hasComment
}

func (d MysqlData) Run() {
	if d.SqlFilename == "" {
		log.Println("入参异常")
		return
	}

	if _, err := os.Stat(d.SqlFilename); os.IsNotExist(err) {
		log.Fatalf("SQL文件%s不存在", d.SqlFilename)
	}

	statements, err := common.GetFileContents(d.SqlFilename)
	if err != nil {
		log.Fatal(err)
	}

	resultCheckStatementComment := true

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		stmtUpper := strings.ToUpper(stmt)
		addColumnRegex := regexp.MustCompile(`^ALTER TABLE [A-Z0-9]+ ADD COLUMN`)

		// 建表或者新增字段语句检测，是否包含COMMENT注释字段
		if addColumnRegex.MatchString(stmtUpper) {
			if !strings.Contains(stmtUpper, "COMMENT") {
				resultCheckStatementComment = false
				log.Println("存在字段定义行没有COMMENT注释：", stmtUpper)
			}
		} else if strings.HasPrefix(stmtUpper, "CREATE TABLE") {
			resultCheckStatementComment = checkStatementComment(stmtUpper)
		}

		if !resultCheckStatementComment {
			log.Fatalf("SQL脚本%s规范性检测异常，存在表或者字段缺少COMMENT注释，语句如下：\n%s", d.SqlFilename, stmt)
		}
	}

	log.Printf("SQL脚本%s规范性检测正常", d.SqlFilename)
}
