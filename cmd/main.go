package main

import (
	"context"
	"github.com/code-couch-potatoes/lingo-flow-back/internal/skyeng"
	"github.com/code-couch-potatoes/lingo-flow-back/internal/yadict"
	"log"
	"os"
)

func main() {
	testSkyeng()

	testYaDict()
}

func testSkyeng() {
	skyClient := skyeng.NewClient()

	words, err := skyClient.WordsSearch(context.Background(), "language")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(words)
}

func testYaDict() {
	ydClient := yadict.NewClient(os.Getenv("YADICT_API_KEY"))

	words, err := ydClient.LookUp(context.Background(), yadict.LangEnRu, "language")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(words)
}
