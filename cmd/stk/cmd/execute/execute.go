package execute

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "execute",
		Short: "execute database sql",
		Long:  `执行数据库脚本`,
	}
)
