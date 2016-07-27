package cmd_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/raphael-trzpit/liard/cmd"
)

var out io.Writer = os.Stdout

// Test no questions.
func TestNoQuestion(t *testing.T) {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	cmd.AskQuestion()
	if len(out.(*bytes.Buffer).String()) > 0 {
		t.Error("Something was outputed during the question.")
	}
}

var questionTests = []struct{
	qs []*cmd.Question,
	ins []string,
	expected
}

// Test questions.
func TestQuestions(t *testing.T) {
	bak := out
	out = new(bytes.Buffer)
	defer func() { out = bak }()

	cmd.AskQuestion()
	_, err = io.WriteString(in, "4 5\n"+"1 2 3 4\n")
  if err != nil {
      t.Fatal(err)
  }
	if len(out.(*bytes.Buffer).String()) > 0 {
		t.Error("Something was outputed during the question.")
	}
}
