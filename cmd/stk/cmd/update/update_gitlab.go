package update

import (
	"github.com/spf13/cobra"
	"github.com/yecaowulei/sre-tool-kit/internal/update"
)

var (
	gitlabAddr           string
	gitlabProjectId      string
	gitlabToken          string
	gitlabFileNameList   string
	gitlabProjectBranch  string
	gitlabCommitUsername string
	updateGitlabDataCmd  = &cobra.Command{
		Use:     "gitlab",
		Short:   "update gitlab file",
		Long:    `修改gitlab文件内容`,
		Example: `stk update gitlab --gitlab-addr http://gitlab.xxxx.com/ --gitlab-project-id 201 --gitlab-project-branch master --gitlab-token xxx --gitlab-filename-list test/test.yml --gitlab-commit-username xx`,
		Args:    cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			updater := update.GitlabData{
				GitlabAddr:           gitlabAddr,
				GitlabProjectId:      gitlabProjectId,
				GitlabToken:          gitlabToken,
				GitlabFileNameList:   gitlabFileNameList,
				GitlabProjectBranch:  gitlabProjectBranch,
				GitlabCommitUsername: gitlabCommitUsername,
			}
			updater.Run()
		},
	}
)

func init() {
	Cmd.AddCommand(updateGitlabDataCmd)
	updateGitlabDataCmd.PersistentFlags().StringVarP(&gitlabAddr, "gitlab-addr", "a", "", "gitlab address")
	updateGitlabDataCmd.PersistentFlags().StringVarP(&gitlabProjectId, "gitlab-project-id", "p", "", "gitlab project id")
	updateGitlabDataCmd.PersistentFlags().StringVarP(&gitlabProjectBranch, "gitlab-project-branch", "b", "", "gitlab project branch")
	updateGitlabDataCmd.PersistentFlags().StringVarP(&gitlabToken, "gitlab-token", "t", "", "gitlab token")
	updateGitlabDataCmd.PersistentFlags().StringVarP(&gitlabFileNameList, "gitlab-filename-list", "f", "", "gitlab file name list")
	updateGitlabDataCmd.PersistentFlags().StringVarP(&gitlabCommitUsername, "gitlab-commit-username", "u", "", "gitlab commit username")
}
