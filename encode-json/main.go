package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := &Message{
		Name: "Andygbox",
		Body: "Anybody here?",
		Time: time.Now().Unix(),
	}

	a, _ := json.Marshal(m)

	fmt.Println(string(a))
}
