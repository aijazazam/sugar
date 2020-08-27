package main

import (
	"fmt"
	"strings"
)

func main() {

	A := []uint32{1, 2, 3, 4}

	in := strings.Trim(strings.Join(strings.Split(fmt.Sprint(A), " "), ","), "[]")

	o := fmt.Sprintf( `SELECT MovieId, Name, Genere, Description, Director, Actors, Year, Duration, Rating,
                     Votes, Revenue, CommentsCount FROM Movie WHERE MovieId in (%s)`, in )

	fmt.Println( o )
}

