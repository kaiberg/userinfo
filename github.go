package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const git_website = "https://api.github.com/users"

type GitHubUserInfo struct {
	Login     string `json:"login"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	PublicRepos int  `json:"public_repos"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	HTMLURL   string `json:"html_url"`
}

func GetGithubUserInfo(username string) {
	endpoint := fmt.Sprintf("%s/%s", git_website, username)
	res, err := http.Get(endpoint)
	if err != nil {
		message := fmt.Sprintf("Error fetching user: %s\n", err.Error())
		panic(message)
	}

	var user GitHubUserInfo
	if err := json.NewDecoder(res.Body).Decode(&user)
	err != nil {
		message := fmt.Sprintf("Error decoding: %s", err.Error())
		panic(message)
	}
	fmt.Println("GitHub User Info:")
	fmt.Printf("Login: %s\n", user.Login)
	fmt.Printf("Name: %s\n", user.Name)
	fmt.Printf("Bio: %s\n", user.Bio)
	fmt.Printf("Public Repos: %d\n", user.PublicRepos)
	fmt.Printf("Followers: %d\n", user.Followers)
	fmt.Printf("Following: %d\n", user.Following)
	fmt.Printf("Profile URL: %s\n", user.HTMLURL)
}
