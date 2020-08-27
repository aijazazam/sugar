package main

import (
	"fmt"
	"encoding/json"

	com "common"
)

func main() {

	s := make( []com.Movie, 0 )

	s = append( s, com.Movie{ Name: "Mission Impossible 1", Rating: 9.5 } )
	s = append( s, com.Movie{ Name: "Mission Impossible 2", Rating: 9.6 } )
	s = append( s, com.Movie{ Name: "Mission Impossible 3", Rating: 9.7 } )

	blob, err := json.Marshal( s )
	fmt.Println(string(blob))
	fmt.Println( err )
}

