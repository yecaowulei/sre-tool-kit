package execute

import (
	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/internal/execute"
)

const (
	defaultPort = 3306
)

var (
	dbHost              string
	dbPort              int
	dbName              string
	dbUser              string
	dbPassword          string
	sqlFilename         string
	sqlType             string
	executeMysqlDataCmd = &cobra.Command{
		Use:     "mysql",
		Short:   "execute mysql sql",
		Long:    `执行mysql数据库脚本`,
		Example: `stk execute mysql --db-host xxx --db-port xxx --db-name xxx --db-user xxx --db-password xxx --sql-filename xxx --sql-type xxx`,
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			updater := execute.MysqlData{
				DbHost:      dbHost,
				DbPort:      dbPort,
				DbName:      dbName,
				DbUser:      dbUser,
				DbPassword:  dbPassword,
				SqlFilename: sqlFilename,
				SqlType:     sqlType,
			}
			updater.Run()
		},
	}
)

func init() {
	Cmd.AddCommand(executeMysqlDataCmd)
	executeMysqlDataCmd.PersistentFlags().StringVarP(&dbHost, "db-host", "", "", "mysql address")
	executeMysqlDataCmd.PersistentFlags().IntVarP(&dbPort, "db-port", "", defaultPort, "mysql port")
	executeMysqlDataCmd.PersistentFlags().StringVarP(&dbName, "db-name", "", "", "mysql db name")
	executeMysqlDataCmd.PersistentFlags().StringVarP(&dbUser, "db-user", "", "", "mysql db user")
	executeMysqlDataCmd.PersistentFlags().StringVarP(&dbPassword, "db-password", "", "", "mysql db password")
	executeMysqlDataCmd.PersistentFlags().StringVarP(&sqlFilename, "sql-filename", "", "", "sql filename need to execute")
	executeMysqlDataCmd.PersistentFlags().StringVarP(&sqlType, "sql-type", "", "", "sql type,select or other")
}
