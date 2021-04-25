package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	messages = [][]string{
		{"Привет, Мир!", "Russian"},
		{"Hello World!", "English"},
		{"Hola Mundo!", "Spanish"},
		{"Olá Mundo!", "Portugal"},
		{"Hej världen!", "Swedish"},
		{"Здравей свят!", "Bolgarian"},
		{"Hei maailma!", "Finnish"},
	}

	seed = time.Now().UnixNano()
	rnd  = rand.New(rand.NewSource(seed))
)

func message() []string {
	return messages[rnd.Intn(len(messages))]
}

func main() {
	m := message()
	fmt.Printf("%s (%s)\n", m[0], m[1])
}
