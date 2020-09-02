package handler

import (
	"io"
	"fmt"
	"strconv"
	"net/http"
	//"io/ioutil"
	"encoding/json"

	"apigateway/bloom"

	com "common"
	sql "apigateway/db"
	auth "apigateway/authenticate"
)

type CreateUser struct {
	UserName	string
	IsAdmin		bool
	Password	string
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println( r )
			io.WriteString(w, `{"Status": Failed}`)
		}
	}()

	c := CreateUser{}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(c.UserName) == 0 || len(c.Password) == 0 {
		w.WriteHeader( http.StatusUnprocessableEntity )
		io.WriteString(w, `{"error":"empty_credentials, please enter both username and password"}`)
		return
	}

	tPasswordHash, tPasswordSalt := auth.GeneratePasswordHashSalt( c.Password )

	db := sql.GetDB()
	credential := com.Credential{ c.UserName, strconv.FormatUint( tPasswordHash, 10), tPasswordSalt, c.IsAdmin, false }
	db.InsertCredentials( credential )

	// In case there are multiple apigateways, send it through kafka to all apigateway's
	// so that their scalable bloom filter is updated with just newly added username
	bloom.AddBloom <- credential

	io.WriteString(w, `{"Status": Successfull}`)

	return
}

