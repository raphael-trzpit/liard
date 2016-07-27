package cmd

import (
	"fmt"
	"io"
	"log"
)

// A question asked by command line.
type Question struct {
	Text      string            // The text of the question that will be displayed.
	Label     string            // Label to identify the question.
	Check     func(string) bool // Check function to see if the answer is valid.
	ErrorText string            // Text displayed if check function fails.
}

// Ask one or several questions.
func AskQuestion(in io.Reader, out io.Writer, qs ...*Question) map[string]string {
	m := make(map[string]string)

	for _, q := range qs {
		var response string
		var isValid bool = false

		fmt.Fprintln(out, q.Text)

		for isValid == false {
			_, err := fmt.Fscanln(in, &response)
			if err != nil {
				log.Fatal(err)
			}
			isValid = q.Check(response)
			if isValid == false {
				fmt.Fprintln(out, q.ErrorText)
			}
		}
		m[q.Label] = response
	}

	return m
}
