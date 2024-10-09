package update

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"log"
	"os"
	"strings"
)

type GitlabData struct {
	GitlabAddr           string
	GitlabProjectId      string
	GitlabToken          string
	GitlabFileNameList   string
	GitlabProjectBranch  string
	GitlabCommitUsername string
}

func (d GitlabData) Run() {
	if d.GitlabAddr == "" || d.GitlabProjectId == "" || d.GitlabProjectBranch == "" || d.GitlabToken == "" || d.GitlabFileNameList == "" {
		fmt.Println("入参异常")
		return
	}

	if d.GitlabCommitUsername == "" {
		d.GitlabCommitUsername = "None"
	}

	client, err := gitlab.NewClient(d.GitlabToken, gitlab.WithBaseURL(d.GitlabAddr))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fileNameList := strings.Split(d.GitlabFileNameList, ",")
	var commitActions []*gitlab.CommitActionOptions

	for _, fileName := range fileNameList {
		content, err := os.ReadFile(fileName)
		if err != nil {
			log.Fatalf("Failed to read file %s: %v", fileName, err)
		}

		// 判断文件在gitlab上是否已存在，存在就更新，不存在就新增
		_, _, err = client.RepositoryFiles.GetFile(d.GitlabProjectId, fileName, &gitlab.GetFileOptions{
			Ref: gitlab.Ptr(d.GitlabProjectBranch),
		})

		var fileModifyType string
		if err != nil {
			// 不存在
			fileModifyType = "create"
		} else {
			// 存在
			fileModifyType = "update"
		}

		action := &gitlab.CommitActionOptions{
			Action:   (*gitlab.FileActionValue)(gitlab.Ptr(fileModifyType)),
			FilePath: gitlab.Ptr(fileName),
			Content:  gitlab.Ptr(string(content)),
		}
		commitActions = append(commitActions, action)
	}

	// 创建提交选项，使用多个文件更新
	commitOptions := &gitlab.CreateCommitOptions{
		Branch:        gitlab.Ptr(d.GitlabProjectBranch),
		CommitMessage: gitlab.Ptr(fmt.Sprintf("fix(#110): update file %s[%s]", d.GitlabFileNameList, d.GitlabCommitUsername)),
		Actions:       commitActions,
	}

	_, _, err = client.Commits.CreateCommit(d.GitlabProjectId, commitOptions)
	if err != nil {
		log.Fatalf("Failed to update file: %v", err)
	}

	fmt.Printf("update file %s successfully.", d.GitlabFileNameList)
}
