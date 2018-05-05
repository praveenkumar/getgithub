package getgithub

import (
	"net/http"
	"os"
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"sync"
	"github.com/pkg/errors"
	"fmt"
	"path/filepath"
)

var (
	client *github.Client
	once   sync.Once
)

var tokenEnvVars = "GH_TOKEN"

// Client return github client.
func Client() *github.Client {
	once.Do(func() {
		token := os.Getenv(tokenEnvVars)
		var tc *http.Client
		if len(token) > 0 {
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{AccessToken: token},
			)
			tc = oauth2.NewClient(oauth2.NoContext, ts)
		}
		client = github.NewClient(tc)
	})
	return client
}

func GetRepositoryContentGetOptions(ref string) *github.RepositoryContentGetOptions {
	return &github.RepositoryContentGetOptions{Ref: ref}
}

// GetBranchOrTag verify if requested branch or tag present and return if there.
func GetBranchOrTag(client *github.Client, ctx context.Context, owner, repo, branch string) (string, error) {
	branchObject, resp, err := client.Repositories.GetBranch(ctx, owner, repo, branch)
	if err != nil {
		repoTags, resp, err := client.Repositories.ListTags(ctx, owner, repo, &github.ListOptions{Page:1, PerPage:100})
		if err != nil {
			return "", err
		}
		if resp.StatusCode != http.StatusOK {
			return "", errors.Errorf("Not able to connect Status code: %d\n", resp.StatusCode)
		}
		for _, tag := range(repoTags) {
			if *tag.Name == branch {
				return *tag.Name, nil
			}
		}
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.Errorf("Not able to connect Status code: %d\n", resp.StatusCode)
	}
	return *branchObject.Name, nil
}

// DownloadContent download the file or directory requested by user and put to destination path.
func DownloadContent(client *github.Client, ctx context.Context, owner, repo, path, branch, dest string)  error {
	fileContent, dirContent, _, err := client.Repositories.GetContents(ctx, owner, repo, path, GetRepositoryContentGetOptions(branch))
	if err != nil {
		return err
	}
	if len(dirContent) > 0 {
		for _, dir := range dirContent {
			if *dir.Type == "dir" {
				if err := os.MkdirAll(filepath.Join(dest, *dir.Path),0777); err != nil {
					return err
				}
			} else {
				if err := os.MkdirAll(filepath.Join(dest, path),0777); err != nil {
					return err
				}
			}
			if err := DownloadContent(client, ctx, owner, repo, *dir.Path, branch, dest); err != nil {
				return err
			}
		}
		}
	if fileContent != nil{
		file, err := os.Create(*fileContent.Path)
		if err != nil {
			return err
		}
		defer file.Close()

		content, err := fileContent.GetContent()
		if err != nil {
			return err
		}
		fmt.Fprintf(file, content)
	}
	return nil
}
