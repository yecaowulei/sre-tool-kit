package transfer

import (
	"github.com/spf13/cobra"
)

var (
	noReport bool
	Cmd      = &cobra.Command{
		Use:   "transfer",
		Short: "transfer mysql data",
		Long:  `从源mysql库同步表结构及数据到目标mysql库`,
	}
)
