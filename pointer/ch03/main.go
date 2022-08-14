package main

import (
	"fmt"
	"time"
)

type (
	GitLogin struct {
		Account   string
		Address   string
		LoginTime time.Time
		Project   []struct {
			Name   string
			ExTime time.Time
			Auther string
		}
	}
)

func main() {
	str := "Felix"
	str1 := "gitlab"
	gitLogin := &GitLogin{
		Account:   "Felix",
		Address:   "github",
		LoginTime: time.Now(),
		Project: []struct {
			Name   string
			ExTime time.Time
			Auther string
		}{{Name: str, ExTime: time.Now(), Auther: "test"}},
	}

	gitLogin1 := gitLogin
	gitLogin1.Project[0].Name = str1
	fmt.Println(gitLogin.Project[0].Name)
	fmt.Println(gitLogin1.Project[0].Name)
}
