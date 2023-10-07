package main

import (
	"context"
	"github.com/code-couch-potatoes/lingo-flow-back/internal/yadict"
	"log"
	"os"
)

func main() {
	ydClient := yadict.NewClient(os.Getenv("YADICT_API_KEY"))

	words, err := ydClient.LookUp(context.Background(), yadict.LangEnRu, "language")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(words)
}
