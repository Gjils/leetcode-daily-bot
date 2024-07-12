package entities

type Group struct {
	ChatId int `bson:"chatId"`
	IsThread bool `bson:"isThread"`
	ThreadId int `bson:"threadId"`
	Title string `bson:"title"`
}

type GroupInfo struct {
	ChatId int `bson:"chatId"`
	IsThread bool `bson:"isThread"`
	ThreadId int `bson:"threadId"`
	Title string `bson:"title"`
}

type GroupQuery struct {
	ChatId int `bson:"chatId"`
	ThreadId int `bson:"threadId"`
}