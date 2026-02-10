package main

type Book struct {
	Title string
	Desc  string
	Type  string
	ISBN  string
}

var bookDatabase = map[string]Book{
	"9781386123613": {
		Title: "The Two Week Curse",
		Desc:  "Two soldiers are transported to a world of magic, cultivation and stat screens and have to rely on old skills and newly learned ones to survive.",
		Type:  "E-Book",
		ISBN:  "9781386123613",
	},
	"9781638493495": {
		Title: "The Two Week Curse",
		Desc:  "Two soldiers are transported to a world of magic, cultivation and stat screens and have to rely on old skills and newly learned ones to survive.",
		Type:  "Hardcover",
		ISBN:  "9781638493495",
	},
}
