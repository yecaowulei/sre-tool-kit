package check

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "check",
		Short: "check database sql",
		Long:  `检测数据库脚本规范性`,
	}
)
