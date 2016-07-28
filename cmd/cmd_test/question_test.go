package cmd_test

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/raphael-trzpit/liard/cmd"
)

// Test questions.
func TestAskQuestions(t *testing.T) {
	var questionTests = []struct {
		qs   []*cmd.Question
		ins  []string
		outs []string
	}{
		{
			[]*cmd.Question{},
			[]string{},
			[]string{},
		},
		{
			[]*cmd.Question{},
			[]string{"test", "toto"},
			[]string{},
		},
		{
			[]*cmd.Question{
				{
					Text:      "text",
					Label:     "a",
					ErrorText: "fail",
					Check:     func(s string) bool { return true },
				},
			},
			[]string{"test"},
			[]string{"text"},
		},
		{
			[]*cmd.Question{
				{
					Text:      "question",
					Label:     "a",
					ErrorText: "fail",
					Check:     func(s string) bool { return s == "success" },
				},
			},
			[]string{"test", "success"},
			[]string{"question", "fail"},
		},
		{
			[]*cmd.Question{
				{
					Text:      "question1",
					Label:     "a",
					ErrorText: "fail1",
					Check:     func(s string) bool { return s == "success1" },
				},
				{
					Text:      "question2",
					Label:     "a",
					ErrorText: "fail2",
					Check:     func(s string) bool { return s == "success2" },
				},
			},
			[]string{"test", "success1", "test", "success2"},
			[]string{"question1", "fail1", "question2", "fail2"},
		},
	}

	for _, qt := range questionTests {
		var out bytes.Buffer
		// Simulate user input
		in, err := ioutil.TempFile("", "")
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			err = in.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()

		for _, s := range qt.ins {
			_, err = in.WriteString(s + "\n")
			if err != nil {
				t.Fatal(err)
			}
		}

		_, err = in.Seek(0, os.SEEK_SET)
		if err != nil {
			t.Fatal(err)
		}

		_, err = cmd.AskQuestion(in, &out, qt.qs...)
		if err != nil {
			t.Fatal(err)
		}
		output := strings.Split(out.String(), "\n")

		for i, e := range qt.outs {
			if len(output) < i+2 {
				t.Errorf("No output was found, expected : %s", e)
				continue
			}

			if output[i] != e {
				t.Errorf("The output is wrong : expected %s, actual %s", e, output[i])
			}
		}
	}

}
