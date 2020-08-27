package handler

import (
	"io"
	"fmt"
	"net/http"
	//"io/ioutil"
	"encoding/json"

	sql "recommendation/db"
)

func Home(w http.ResponseWriter, r *http.Request) {

	//io.WriteString(w, `{"jwt":"`+token+`"}`)
	db := sql.GetDB()

	recommended := db.GetRecommend( 8.0, 100 )

	blob, err := json.Marshal( recommended )
	if err != nil {
		fmt.Printf("Failed to Marchal json with err=", err)
		return
	}

	io.WriteString( w, string(blob) )

	return
}

