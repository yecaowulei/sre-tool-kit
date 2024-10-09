package transfer

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

type MysqlData struct {
	SourceDbHost     string
	SourceDbPort     int
	SourceDbName     string
	SourceTableName  string
	IgnoreTable      string
	SourceDbUser     string
	SourceDbPassword string
	DstDbHost        string
	DstDbPort        int
	DstDbName        string
	DstDbUser        string
	DstDbPassword    string
	NeedData         string
}

var (
	ignoreCmd    string
	nodataCmd    string
	tableNameCmd string
)

func (d MysqlData) Run() {
	if d.SourceDbHost == "" || d.SourceDbName == "" || d.SourceDbUser == "" || d.SourceDbPassword == "" || d.DstDbHost == "" || d.DstDbName == "" || d.DstDbUser == "" || d.DstDbPassword == "" {
		log.Println("入参异常")
		return
	}

	if d.NeedData == "true" {
		nodataCmd = ""
	} else {
		nodataCmd = "--no-data"
	}

	if d.SourceTableName == "" {
		tableNameCmd = ""
		// 没有指定需要同步的表，--ignore-table可以生效
		if d.IgnoreTable == "" {
			ignoreCmd = ""
		} else {
			for _, tableName := range strings.Split(d.IgnoreTable, ",") {
				ignoreCmd += fmt.Sprintf("--ignore-table=%s.%s ", d.SourceDbName, tableName)
			}
		}
	} else {
		for _, tableName := range strings.Split(d.SourceTableName, ",") {
			tableNameCmd += fmt.Sprintf("%s ", tableName)
		}
		// 指定了需要同步的表，则--ignore-table不生效
		ignoreCmd = ""
	}

	exportCmd := fmt.Sprintf("mysqldump -h%s -P%d -u%s -p'%s' %s %s --set-gtid-purged=OFF --quick -c --skip-lock-tables %s %s| sed -e 's/ DEFINER=[^ ]* / /' >%s.sql", d.SourceDbHost, d.SourceDbPort, d.SourceDbUser, d.SourceDbPassword, d.SourceDbName, tableNameCmd, nodataCmd, ignoreCmd, d.SourceDbName)
	reExport := regexp.MustCompile(`-p'\S+`)
	newExportCmd := reExport.ReplaceAllString(exportCmd, "-pxxx")
	log.Printf("开始导出，导出命令为：%s\n", newExportCmd)

	execExportCmd := exec.Command("bash", "-c", exportCmd)
	exportOutput, err := execExportCmd.CombinedOutput()
	exportLowerOutput := strings.ToLower(string(exportOutput))
	if strings.Contains(exportLowerOutput, "error") || strings.Contains(exportLowerOutput, "command not found") {
		log.Fatalf("导出报错： %s\n", string(exportOutput))
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("导出完成")

	importCmd := fmt.Sprintf("mysql -h%s -P%d -u%s -p'%s' %s <%s.sql", d.DstDbHost, d.DstDbPort, d.DstDbUser, d.DstDbPassword, d.DstDbName, d.SourceDbName)
	reImport := regexp.MustCompile(`-p'\S+`)
	newImportCmd := reImport.ReplaceAllString(importCmd, "-pxxx")
	log.Printf("开始导入，导入命令为：%s\n", newImportCmd)

	execImportCmd := exec.Command("bash", "-c", importCmd)
	importOutput, err := execImportCmd.CombinedOutput()
	importLowerOutput := strings.ToLower(string(importOutput))
	if strings.Contains(importLowerOutput, "error") {
		log.Fatalf("导入报错： %s\n", string(importOutput))
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Println("导入完成")
}
