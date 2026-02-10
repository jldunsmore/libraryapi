package main

type User struct {
	Name  string
	Token string
}

var userDatabase = map[string]User{
	"user1": {
		Name:  "John Doe",
		Token: "1234",
	},
}
