//go:build exclude

package main

import (
	"fmt"
	"os"

	"github.com/robherley/guesslang-go/pkg/guesser"
)

/*

This is a simple example of how to use the guesser package.

Expected output:

Best => rs
Reliable => true
Top Five:
- rs (21.66%)
- java (4.87%)
- ts (4.66%)
- js (4.62%)
- html (4.07%)

*/

const snippet = `
fn bubble_sort(arr: &mut [i32]) {
	let len = arr.len();
	for i in 0..len {
			for j in 0..len - i - 1 {
					if arr[j] > arr[j + 1] {
							arr.swap(j, j + 1);
					}
			}
	}
}
`

func main() {
	gsr, err := guesser.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	answer, err := gsr.Guess(snippet)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Best =>", answer.Predictions[0].Language)
	fmt.Println("Reliable =>", answer.Reliable)
	fmt.Println("Top Five:")
	for _, lang := range answer.Predictions[:5] {
		fmt.Printf("- %s (%.2f%%)\n", lang.Language, lang.Confidence*100)
	}
}
