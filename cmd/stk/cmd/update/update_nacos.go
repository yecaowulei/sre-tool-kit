package update

import (
	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/internal/update"
)

const (
	defaultScheme     = "https"
	defaultNacosGroup = "DEFAULT_GROUP"
)

var (
	nacosAddr          string
	nacosAddrScheme    string
	nacosNsId          string
	nacosUsername      string
	nacosPasswd        string
	nacosNsGroup       string
	nacosFileNameList  string
	nacosPort          uint64
	updateNacosDataCmd = &cobra.Command{
		Use:     "nacos",
		Short:   "update nacos file",
		Long:    `修改nacos文件内容`,
		Example: `stk update nacos --nacos-addr nacos.xxx.com --nacos-addr-scheme https --nacos-ns-id test --nacos-username test --nacos-passwd test --nacos-filename-list test1.yml,test2.yml`,
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			updater := update.NacosData{
				NacosAddr:         nacosAddr,
				NacosAddrScheme:   nacosAddrScheme,
				NacosNsId:         nacosNsId,
				NacosUsername:     nacosUsername,
				NacosPasswd:       nacosPasswd,
				NacosNsGroup:      nacosNsGroup,
				NacosFileNameList: nacosFileNameList,
				NacosPort:         nacosPort,
			}
			updater.Run()
		},
	}
)

func init() {
	Cmd.AddCommand(updateNacosDataCmd)
	updateNacosDataCmd.PersistentFlags().StringVarP(&nacosAddr, "nacos-addr", "a", "", "nacos address")
	updateNacosDataCmd.PersistentFlags().StringVarP(&nacosAddrScheme, "nacos-addr-scheme", "s", defaultScheme, "nacos address scheme")
	updateNacosDataCmd.PersistentFlags().StringVarP(&nacosNsId, "nacos-ns-id", "n", "", "nacos namespace id")
	updateNacosDataCmd.PersistentFlags().StringVarP(&nacosUsername, "nacos-username", "u", "", "nacos username")
	updateNacosDataCmd.PersistentFlags().StringVarP(&nacosPasswd, "nacos-passwd", "p", "", "nacos passwd")
	updateNacosDataCmd.PersistentFlags().StringVarP(&nacosNsGroup, "nacos-ns-group", "g", defaultNacosGroup, "nacos namespace group")
	updateNacosDataCmd.PersistentFlags().StringVarP(&nacosFileNameList, "nacos-filename-list", "f", "", "nacos file name list")
	updateNacosDataCmd.PersistentFlags().Uint64VarP(&nacosPort, "nacos-port", "P", 80, "nacos server port")
}
