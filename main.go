package main

import (
	"flag"
	"fmt"
	"context"
	"github.com/praveenkumar/getgithub/pkg/getgithub"
	"os"
)

func main() {
	var repo, dir, owner, branch, dest, cwd string
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	flag.StringVar(&owner, "owner", "", "Name of the repo owner")
	flag.StringVar(&repo, "repo", "", "Name of the repo")
	flag.StringVar(&dir, "dir", "", "Directory or file to download")
	flag.StringVar(&branch, "branch", "master", "branch or tag")
	flag.StringVar(&dest, "dest", cwd , "Destination directory")
	flag.Parse()

	if repo == "" || dir == "" || owner == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	gitHubClient := getgithub.Client()
	ctx := context.Background()

	branchName, err := getgithub.GetBranchOrTag(gitHubClient, ctx, owner, repo, branch)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = getgithub.DownloadContent(gitHubClient, ctx, owner, repo, dir, branchName, dest)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

