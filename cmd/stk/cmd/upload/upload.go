package upload

import (
	"github.com/spf13/cobra"
)

var (
	Cmd = &cobra.Command{
		Use:   "upload",
		Short: "upload local file to obs",
		Long:  `从本地上传文件到华为云对象存储`,
	}
)
