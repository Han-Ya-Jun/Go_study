package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"time"
)

/*
* @Author:hanyajun
* @Date:2019/6/21 10:56
* @Name:github_api
* @Function:
 */

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "*************************"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	c := "update readme.md"
	sha := "b43cb39ddc29f328888b5840582b8039a1ac0a6a"
	//auth:=&github.Contributor{
	//
	//}
	var cstZone = time.FixedZone("CST", 8*3600)
	t := time.Now().Add(time.Hour * 13).In(cstZone)
	content := &github.RepositoryContentFileOptions{
		Message: &c,
		SHA:     &sha,
		Committer: &github.CommitAuthor{
			Date:  &t,
			Name:  github.String("hanyajun"),
			Email: github.String("1581532052@qq.com"),
			Login: github.String("Han-Ya-Jun"),
		},
		Author: &github.CommitAuthor{
			Date:  &t,
			Name:  github.String("hanyajun"),
			Email: github.String("1581532052@qq.com"),
			Login: github.String("Han-Ya-Jun"),
		},
		Branch: github.String("master"),
	}
	op := &github.RepositoryContentGetOptions{}

	repo, _, _, er := client.Repositories.GetContents(ctx, "Han-Ya-Jun", "gocn_news_set", "README.md", op)
	if er != nil || repo == nil {
		fmt.Println(er)
		return
	}
	content.SHA = repo.SHA
	decodeBytes, err := base64.StdEncoding.DecodeString(*repo.Content)
	if err != nil {
		log.Fatalln(err)
	}
	//oldContentList := strings.Split(string(decodeBytes), "## gocn_news_set_2019")
	content.Content = []byte(string(decodeBytes) + "123")
	repos, _, err := client.Repositories.UpdateFile(ctx, "Han-Ya-Jun", "gocn_news_set", "README.md", content)
	if err != nil {
		println(err)
	}
	fmt.Println(repos.SHA)

}
