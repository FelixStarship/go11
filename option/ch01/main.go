package main

import (
	"fmt"
	"log"
)

type User struct {
	ID     string
	Name   string
	Age    int
	Email  string
	Phone  string
	Gender string
}

type Option func(*User)

func WithAge(age int) Option {
	return func(u *User) {
		u.Age = age
	}
}

func WithEmail(email string) Option {
	return func(u *User) {
		u.Email = email
	}
}

func WithPhone(phone string) Option {
	return func(u *User) {
		u.Phone = phone
	}
}

func WithGender(gender string) Option {
	return func(u *User) {
		u.Gender = gender
	}
}

func NewUser(id string, name string, options ...func(*User)) (*User, error) {
	user := User{
		ID:     id,
		Name:   name,
		Age:    0,
		Email:  "",
		Phone:  "",
		Gender: "felix",
	}
	for _, option := range options {
		option(&user)
	}
	//...
	return &user, nil
}

func main() {

	user, err := NewUser("1", "Golang",
		WithAge(27),
		WithPhone("123456789"),
		WithEmail("2793800263@qq.com"),
		WithGender("Felix"))

	if err != nil {
		log.Fatalf("NewUser:err:%+v", err)
	}

	fmt.Println(*user)
}
