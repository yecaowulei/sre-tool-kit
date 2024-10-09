package transfer

import (
	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/internal/transfer"
)

const (
	defaultPort     = 3306
	defaultNeedData = "true"
)

var (
	sourceDbHost         string
	sourceDbPort         int
	sourceDbName         string
	sourceTableName      string
	ignoreTable          string
	sourceDbUser         string
	sourceDbPassword     string
	dstDbHost            string
	dstDbPort            int
	dstDbName            string
	dstDbUser            string
	dstDbPassword        string
	needData             string
	transferMysqlDataCmd = &cobra.Command{
		Use:     "mysql",
		Short:   "迁移Mysql指定库数据",
		Long:    `迁移Mysql指定库数据`,
		Example: `stk transfer mysql --source-db-host xxx --source-db-port xxx --source-db-name xxx --ignore-table=xx.xxx,xx.xxx --source-db-user xxx --source-db-password xxx --dst-db-host xxx --dst-db-port xxx --dst-db-name xxx --dst-db-user xxx --dst-db-password xxx`,
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			updater := transfer.MysqlData{
				SourceDbHost:     sourceDbHost,
				SourceDbPort:     sourceDbPort,
				SourceDbName:     sourceDbName,
				SourceTableName:  sourceTableName,
				IgnoreTable:      ignoreTable,
				SourceDbUser:     sourceDbUser,
				SourceDbPassword: sourceDbPassword,
				DstDbHost:        dstDbHost,
				DstDbPort:        dstDbPort,
				DstDbName:        dstDbName,
				DstDbUser:        dstDbUser,
				DstDbPassword:    dstDbPassword,
				NeedData:         needData,
			}
			updater.Run()
		},
	}
)

func init() {
	Cmd.AddCommand(transferMysqlDataCmd)
	transferMysqlDataCmd.PersistentFlags().StringVarP(&sourceDbHost, "source-db-host", "", "", "source mysql address")
	transferMysqlDataCmd.PersistentFlags().IntVarP(&sourceDbPort, "source-db-port", "", defaultPort, "source mysql port")
	transferMysqlDataCmd.PersistentFlags().StringVarP(&sourceDbName, "source-db-name", "", "", "source mysql db name")
	transferMysqlDataCmd.PersistentFlags().StringVarP(&sourceTableName, "source-table-name", "", "", "source mysql table name")
	transferMysqlDataCmd.PersistentFlags().StringVarP(&ignoreTable, "ignore-table", "", "", "ignore table list")
	transferMysqlDataCmd.PersistentFlags().StringVarP(&sourceDbUser, "source-db-user", "", "", "source mysql db user")
	transferMysqlDataCmd.PersistentFlags().StringVarP(&sourceDbPassword, "source-db-password", "", "", "source mysql db password")
	transferMysqlDataCmd.PersistentFlags().StringVarP(&dstDbHost, "dst-db-host", "", "", "dst mysql address")
	transferMysqlDataCmd.PersistentFlags().IntVarP(&dstDbPort, "dst-db-port", "", defaultPort, "dst mysql port")
	transferMysqlDataCmd.PersistentFlags().StringVarP(&dstDbName, "dst-db-name", "", "", "dst mysql db name")
	transferMysqlDataCmd.PersistentFlags().StringVarP(&dstDbUser, "dst-db-user", "", "", "dst mysql db user")
	transferMysqlDataCmd.PersistentFlags().StringVarP(&dstDbPassword, "dst-db-password", "", "", "dst mysql db password")
	transferMysqlDataCmd.PersistentFlags().StringVarP(&needData, "need-data", "", defaultNeedData, "whether data is needed")
}
