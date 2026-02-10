package main

type ClientProfile struct {
	Email string
	Id    string
	Name  string
	Token string
}

var database = map[string]ClientProfile{
	"user1": {
		Email: "george_bush@email.com",
		Id:    "user1",
		Name:  "George Buish",
		Token: "123",
	},
	"user2": {
		Email: "jonDoe@email.com",
		Id:    "user2",
		Name:  "jon doe",
		Token: "456",
	},
}
