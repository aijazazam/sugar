package main

import (
	"fmt"
	"github.com/schollz/closestmatch"
)

// closestmatch is more accurate than Levenshtein for long strings (like in the test corpus).
func main() {

	// Take a slice of keys, say band names that are similar
	// http://www.tonedeaf.com.au/412720/38-bands-annoyingly-similar-names.htm
	wordsToTest := []string{"King Gizzard", "The Lizard Wizard", "Lizzard Wizzard"}

	// Choose a set of bag sizes, more is more accurate but slower
	bagSizes := []int{2}

	// Create a closestmatch object
	cm := closestmatch.New(wordsToTest, bagSizes)

	cm.Add([]string{"aijaz"})

	//fmt.Println(cm.Closest("kind gizard"))
	fmt.Println(cm.Closest("wizz"))
	// returns 'King Gizzard'

	fmt.Println(cm.ClosestN("wizz",2))
	// returns [King Gizzard Lizzard Wizzard The Lizard Wizard]

	// Calculate accuracy
	fmt.Println(cm.AccuracyMutatingWords())
	// ~ 66 % (still way better than Levenshtein which hits 0% with this particular set)

	// Improve accuracy by adding more bags
	bagSizes = []int{2, 3, 4}
	cm = closestmatch.New(wordsToTest, bagSizes)
	fmt.Println(cm.AccuracyMutatingWords())
	// accuracy improves to ~ 76 %
}

