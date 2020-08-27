package handler

import (
	"io"
	"fmt"
	"net/http"
	//"io/ioutil"
	"encoding/json"

	com "common"
	en "search/cache"
	sql "search/db"
)

func Search(w http.ResponseWriter, r *http.Request) {

	fmt.Println( r.URL.Path )

	toSearch, ok := r.URL.Query()["q"]
	if !ok {
		fmt.Printf("/search=< search string >")
		return
	}

	rsp := make( chan en.SearchMovie, 1 )

	en.SearchMovies <- en.SearchMovie{ Query: toSearch[0], Response: rsp }

	result := <-rsp

	db := sql.GetDB()

	movie_list := make( []com.Movie, 0 )

	if len(result.Edit_distance_0) > 0 {
		tmp_movie_list := db.GetSearch( result.Edit_distance_0 )
		movie_list = append( movie_list, tmp_movie_list... )
	}

	if len(result.Edit_distance_1) > 0 {
		tmp_movie_list := db.GetSearch( result.Edit_distance_1 )
		movie_list = append( movie_list, tmp_movie_list... )
	}

	if len(result.Edit_distance_0) == 0 && len(result.Edit_distance_1) == 0 && len(result.Edit_distance_2) > 0 {
		tmp_movie_list := db.GetSearch( result.Edit_distance_2 )
		movie_list = append( movie_list, tmp_movie_list... )
	}

	blob, err := json.Marshal( movie_list )
	if err != nil {
		fmt.Printf("Failed to Marchal json with err=", err)
		return
	}

	io.WriteString( w, string(blob) )

	return
}

