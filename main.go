package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	userId := os.Args[1]

	// executeVulnerableQuery(userId)
	executeQuery(userId)
}

func executeQuery(userId string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	if !re.MatchString(userId) {
		e := errors.New("Invalid input")
		log.Println(e)
		return e
	}
	sql := fmt.Sprintf("SELECT * FROM Users WHERE UserId = %s", userId)
	fmt.Println(sql)
	return nil
}
