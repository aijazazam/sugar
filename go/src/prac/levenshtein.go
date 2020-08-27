package main

import (
	"fmt"
	"strings"
	"github.com/agnivade/levenshtein"
)

type Cache struct {
	m map[string] []uint32
}

var _EMPTY_ struct{}

func main() {

	cache := Cache{ make( map[string] []uint32 ) }

	//words := strings.Fields( "one    two   three four " )

	cache.Add( "one    two   three four ", 1234 )
	cache.Add( "  two   four ", 24 )

	rsp := cache.Query( "zero" )

	fmt.Println( cache.m )
	fmt.Println( rsp )
}

func ( cache *Cache ) Add( tName string, tMovieId uint32 ) {

	words := strings.Fields(tName)

	for _, w := range words {
		val, ok := cache.m[w]

		if ok {
			val = append( val, tMovieId )
		} else {
			val = make( []uint32, 1 )
			val[0] = tMovieId
		}

		cache.m[w] = val
	}
}

func ( cache *Cache ) Query( user_string string ) []uint32 {

	editDistance := make( []map[uint32]struct{}, 3 )

	editDistance[0] = make( map[uint32]struct{} )
	editDistance[1] = make( map[uint32]struct{} )
	editDistance[2] = make( map[uint32]struct{} )

	user_keyword_list := strings.Fields( user_string )

	for _, user_keyword := range user_keyword_list {

		for keyword_in_movie, movieId_list := range cache.m {

			if dis := levenshtein.ComputeDistance( keyword_in_movie, user_keyword ); dis < 3 {

				for _, movieId := range movieId_list {

					editDistance[dis][movieId] = _EMPTY_
				}
			}
		}
	}

	searchResult := make( []uint32, 0 )

	for movieId, _ := range editDistance[0] {
		searchResult = append( searchResult, movieId )
	}

	for movieId, _ := range editDistance[1] {
		_, alreadyAdded := editDistance[0][movieId]
		if !alreadyAdded {
			editDistance[0][movieId] = _EMPTY_
			searchResult = append( searchResult, movieId )
		}
	}

	for movieId, _ := range editDistance[2] {
		_, alreadyAdded := editDistance[0][movieId]
		if !alreadyAdded {
			editDistance[0][movieId] = _EMPTY_
			searchResult = append( searchResult, movieId )
		}
	}

	return searchResult
}


