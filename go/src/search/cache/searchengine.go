package cache

import (
	"fmt"
	"strings"
	"github.com/agnivade/levenshtein"

	sql "search/db"
)

var	_EMPTY_	struct{}

type SearchMovie struct {
	Query								string
	Edit_distance_0	[]uint32
	Edit_distance_1	[]uint32
	Edit_distance_2	[]uint32
	Response						chan SearchMovie
}

type Cache struct {
  m map[string] []uint32
}

var (
	AddMovie			chan sql.MovieMin
	SearchMovies	chan SearchMovie
)

func init() {

	AddMovie = make( chan sql.MovieMin, 100 )
	SearchMovies = make( chan SearchMovie, 100 )
}

func GoSearchEngine() {

	response := make( chan sql.MovieMin, 100 )
	cache		 := Cache{ make( map[string] []uint32 ) }
	db			 := sql.GetDB()

	go db.GoScanMovieName( response )

	for tMovieMin := range response {
		cache.Add( tMovieMin.Name, tMovieMin.MovieId )
	}

	dist_0, dist_1, dist_2 := cache.SearchQuery( "Lost" )
	fmt.Println( dist_0, dist_1, dist_2 )

	fmt.Println("done scanning db")

	for {

		select {
			case tAdd := <-AddMovie: {
				cache.Add( tAdd.Name, tAdd.MovieId )
			}

			case tSearch := <-SearchMovies: {

				fmt.Println("Searching ", tSearch.Query )

				dist_0, dist_1, dist_2 := cache.SearchQuery( tSearch.Query )

				fmt.Println("Searched ", tSearch.Query, dist_0, dist_1, dist_2 )

				tSearch.Edit_distance_0 = dist_0
				tSearch.Edit_distance_1 = dist_1
				tSearch.Edit_distance_2 = dist_2
				tSearch.Response <- tSearch
			}
		}
	}

	fmt.Println("dead GoSearchEngine ..")
}

func ( cache *Cache ) Add( tName string, tMovieId uint32 ) {

  words := strings.Fields(tName)

  for _, w := range words {
    val, ok := cache.m[strings.ToLower(w)]

    if ok {
      val = append( val, tMovieId )
    } else {
      val = make( []uint32, 1 )
      val[0] = tMovieId
    }

    cache.m[strings.ToLower(w)] = val
  }
}

func ( cache *Cache ) SearchQuery( user_string string ) (distanceZero, distanceOne, distanceTwo []uint32 ) {

  editDistance := make( []map[uint32]struct{}, 3 )

  editDistance[0] = make( map[uint32]struct{} )
  editDistance[1] = make( map[uint32]struct{} )
  editDistance[2] = make( map[uint32]struct{} )

  user_keyword_list := strings.Fields( user_string )

  for _, user_keyword := range user_keyword_list {

    for keyword_in_movie, movieId_list := range cache.m {

      if dis := levenshtein.ComputeDistance( keyword_in_movie, strings.ToLower(user_keyword) ); dis < 3 {

        for _, movieId := range movieId_list {

          editDistance[dis][movieId] = _EMPTY_
        }
      }
    }
  }

  for movieId, _ := range editDistance[0] {
		distanceZero = append( distanceZero, movieId )
  }

  for movieId, _ := range editDistance[1] {
    _, alreadyAdded := editDistance[0][movieId]
    if !alreadyAdded {
      editDistance[0][movieId] = _EMPTY_
			distanceOne = append( distanceOne, movieId )
    }
  }

  for movieId, _ := range editDistance[2] {
    _, alreadyAdded := editDistance[0][movieId]
    if !alreadyAdded {
      editDistance[0][movieId] = _EMPTY_
			distanceTwo = append( distanceTwo, movieId )
    }
  }

  return
}

