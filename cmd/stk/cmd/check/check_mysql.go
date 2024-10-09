package check

import (
	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/internal/check"
)

var (
	sqlFilename       string
	checkMysqlDataCmd = &cobra.Command{
		Use:     "mysql",
		Short:   "check mysql sql",
		Long:    `检测mysql数据库脚本规范性`,
		Example: `stk check mysql --sql-filename xxx`,
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			updater := check.MysqlData{
				SqlFilename: sqlFilename,
			}
			updater.Run()
		},
	}
)

func init() {
	Cmd.AddCommand(checkMysqlDataCmd)
	checkMysqlDataCmd.PersistentFlags().StringVarP(&sqlFilename, "sql-filename", "", "", "sql filename need to check")
}
