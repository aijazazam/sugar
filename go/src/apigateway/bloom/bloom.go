package bloom

import (
	"fmt"
	com "common"
	DB "apigateway/db"
	bloom "github.com/tylertreat/BoomFilters"
)

type Bloom struct {
	Cred			com.Credential
	Present		bool
	Response	chan Bloom
}

var (
	AddBloom   chan Bloom
  CheckBloom chan Bloom
)

func init() {

	AddBloom     = make( chan Bloom, 100 )
  CheckBloom   = make( chan Bloom, 100 )

	go GoBloom()
}

func GoBloom() {

	cCredential := make( chan com.Credential, 100 )

	// sbf -> Scalable Bloom Filters
	sbf := bloom.NewDefaultScalableBloomFilter( 0.01 )

	db := DB.GetDB()

	go db.GoScanCredentials( cCredential )

	for tCredential := range cCredential {
		sbf.Add( []byte( tCredential.UserName ) )
		fmt.Println( "Bloom added ", tCredential.UserName )
	}

	//db.GetCredentials()

	for {
		select {
			case tAdd := <-AddBloom: {
				sbf.Add( []byte( tAdd.Cred.UserName) )
				fmt.Printf( "%v added to Bloom", tAdd.Cred.UserName )
			}
			case tCheck := <-CheckBloom: {
				tCheck.Present = sbf.Test( []byte(tCheck.Cred.UserName) )
				fmt.Printf( "%v is [%v] in Bloom\n", tCheck.Cred.UserName, tCheck.Present )
				tCheck.Response <- tCheck
			}
		}
	}
}

