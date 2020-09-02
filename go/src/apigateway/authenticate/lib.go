package authenticate

import (
	"time"
	"math/rand"
	_ "github.com/go-sql-driver/mysql"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// For good randomeness in production we use crypto/rand
// for this prototype math/rand is ok
func randString(n int) string {

	alpha := make([]rune, n)

	rand.Seed( time.Now().UnixNano() )

	for i := 0; i < n; i++ {
		alpha[i] = letters[rand.Intn(len(letters))]
	}

	return string(alpha)
}

