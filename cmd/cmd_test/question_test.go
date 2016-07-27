package cmd_test

import (
	"bytes"
	"testing"

	"github.com/raphael-trzpit/liard/cmd"
)

// Test no questions.
func TestNoQuestion(t *testing.T) {
	var in, out bytes.Buffer

	cmd.AskQuestion(&in, &out)

	if len(out.String()) > 0 {
		t.Error("Something was outputed during the question.")
	}
}

var questionTests = []struct {
	qs       []*cmd.Question
	ins      []string
	expected []string
}{
	{
		[]*cmd.Question{
			{
				Text:      "question",
				Label:     "a",
				ErrorText: "fail",
				Check:     func(s string) bool { return false },
			},
		},
		[]string{"test"},
		[]string{"question", "fail"},
	},
}

// Test questions.
func TestQuestions(t *testing.T) {

	/*
		_, err = io.WriteString(in, "4 5\n"+"1 2 3 4\n")
		if err != nil {
			t.Fatal(err)
		}
	*/
}
