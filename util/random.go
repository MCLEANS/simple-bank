package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

/*
*
This will be called automatically when the package is first used
*/
func init() {

	/* Set the seed used for randomization */
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min int64, max int64) int64 {

	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {

	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {

		c := alphabet[rand.Intn(k)]

		sb.WriteByte(c)

	}

	return sb.String()
}

/*
generates a random name for an owner
*/
func RandomOwner() string {

	return RandomString(6)
}

/*
generates a random amount of money
*/
func RandomMoney() int64 {

	return RandomInt(0, 10000)
}

func RandomID() int64 {

	return RandomInt(0, 10000)
}

/*
generate an random currency
*/
func RandomCurrency() string {

	currencies := []string{"KSH", "USD"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
