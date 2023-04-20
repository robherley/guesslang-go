package guesser_test

import (
	"fmt"
	"testing"

	"github.com/robherley/guesslang-go/pkg/guesser"
)

const (
	CCode = `
#include <stdio.h>

int main(int argc, char* argv[])
{
	printf("Hello world");
}
`
	PythonCode = `
from __future__ import print_function

if __name__ == "__main__":
		print("Hello world")
`
	RustCode = `
fn main() {
	println!("Hello world");
}
`
	Markdown = `
# Hello World
Check out this [amazing website](https://reb.gg)!
`
)

func TestGuesserNew(t *testing.T) {
	gsr, err := guesser.New()

	if err != nil {
		t.Error(err)
	}

	if gsr == nil {
		t.Error("gsr is nil")
	}
}

func TestGuesserGuess(t *testing.T) {
	cases := []struct {
		lang string
		code string
	}{
		{"c", CCode},
		{"py", PythonCode},
		{"rs", RustCode},
		{"md", Markdown},
	}

	gsr, err := guesser.New()
	if err != nil {
		t.Error(err)
		return
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%s code", tc.lang), func(t *testing.T) {
			answer, err := gsr.Guess(tc.code)
			if err != nil {
				t.Error(err)
				return
			}

			if answer.Predictions[0].Language != tc.lang {
				t.Errorf("want %s, got %s", tc.lang, answer.Predictions[0].Language)
			}

			if !answer.Reliable {
				t.Error("want reliable, got unreliable")
			}
		})
	}
}
