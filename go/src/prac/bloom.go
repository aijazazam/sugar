package main

import (
	"fmt"
	bloom "github.com/tylertreat/BoomFilters"
)

func main() {

	sbf := bloom.NewDefaultScalableBloomFilter( 0.01 )

	sbf.Add([]byte(`a`))

	if sbf.Test([]byte(`a`)) {
		fmt.Println("contains a")
	} else {
		fmt.Println("not contains a")
	}

	if sbf.Test([]byte(`aijaz`)) {
		fmt.Println("contains aijaz")
	} else {
		fmt.Println("not contains aijaz")
	}

	// Restore to initial state.
	sbf.Reset()
}

