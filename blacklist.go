package main 

import "slices"

var blacklist []string = []string{"torvalds", "anthropics","nextcloud","microsoft","hyochan"}

func IsUserBlacklisted(username string) bool {
	return slices.Contains(blacklist, username)
}
