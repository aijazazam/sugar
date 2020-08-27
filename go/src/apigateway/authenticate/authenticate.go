package db

import (
	"fmt"
	"strconv"
	"github.com/dchest/siphash"

	com "common"
	sql "apigateway/db"
	bloom "apigateway/bloom"
)

const (
  HASH_KEY_1 = 0xdeafdeaddeafdeaf
  HASH_KEY_2 = 0xdeafdeafdeafdead
)

func Authentication( tUserName, tPassword string ) ( IsAdmin bool, IsAuthenticated bool ) {

	IsAdmin = false
	IsAuthenticated = false

	if ( len(tUserName) == 0 || len(tPassword) == 0 ) {
		return
	}

	cResponse := make( chan bloom.Bloom, 1 )
	bloom.CheckBloom <- bloom.Bloom{ com.Credential{ UserName: tUserName}, false, cResponse }

	rsp := <-cResponse

	if !rsp.Present {
		fmt.Printf( "UserName = %v, not present in Bloom Filter" , tUserName )
		return
	}

	db := sql.GetDB()

	tCredential := db.GetCredentials( tUserName )

	if(tCredential.IsDisabled) {
		fmt.Println("tUser.IsDisabled")
		return
	}

	dbPasswordHash, err := strconv.ParseUint(tCredential.PasswordHash, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	tDigest := siphash.Hash( HASH_KEY_1, HASH_KEY_2, []byte(tPassword + tCredential.PasswordSalt) )

	if ( tDigest != dbPasswordHash ) {
		fmt.Printf("%v != %v", tDigest, dbPasswordHash)
		return
	}

	return tCredential.IsAdmin, true
}

/*func GeneratePasswordHashSalt( tPassword string ) (tPasswordHash uint64, tPasswordSalt string) {

	// Generate Random String of 10 charecters, this is our Salt specific to given User
	tPasswordSalt = randString(10)
	tPasswordHash = siphash.Hash( HASH_KEY_1, HASH_KEY_2, []byte(tPassword + tPasswordSalt) )

	return
}*/

