package model

import (
	"log"
	"os"
)

type Chat struct {
	ChatId         int    `json:"chatId"`
	ChatName       string `json:"chatName"`
	AvatarUrl      string `json:"avatarUrl"`
	UnreadMessages int    `json:"unreadMessages"`
	LastMessage    string `json:"lastMessage"`
	LastUpdate     int    `json:"lastUpdate"`
}

var chats []Chat

func ChatInit() {
	GetChats()
}

func GetChats() {
	chats = []Chat{}
	var path = "/face2/"
	entries, err := os.ReadDir(root + path)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		chats = append(chats, Chat{ChatId: 1, AvatarUrl: path + e.Name()})
	}
}
