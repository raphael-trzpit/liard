package main

import (
	"log"
	"os"

	"github.com/raphael-trzpit/liard/cmd"
)

func main() {
	q := cmd.Question{
		Text:      "This is a test, press yes.",
		Label:     "test",
		ErrorText: "failed.",
		Check:     func(s string) bool { return s == "yes" },
	}
	qs := []*cmd.Question{&q, &q}

	_, err := cmd.AskQuestion(os.Stdin, os.Stdout, qs...)
	if err != nil {
		log.Fatal(err)
	}
}
