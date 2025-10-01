package main

import (
	"fmt"
	"flag"
)

func main() {
	websitePtr := flag.String("site", "github.com", "choose which website to get data from")
	usernamePtr := flag.String("user", "kaiberg", "choose which user to fetch")
	
	flag.Parse()

	if(*websitePtr != "github.com") {
		message := "that website is currently not supported"
		panic(message)
	}

	if IsUserBlacklisted(*usernamePtr) == true {
		message := "that user is currently not allowed to be fetched"
		panic(message)
	}

	fmt.Printf("Fetching %s from website %s\n", *usernamePtr, *websitePtr)
	if(*websitePtr == "github.com") {
		GetGithubUserInfo(*usernamePtr)
	}
}
