package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var stringB strings.Builder
	alphabetLength := len(alphabet)

	for i := 0; i < n; i++ {
		a := alphabet[rand.Intn(alphabetLength)]
		stringB.WriteByte(byte(a))
	}
	return stringB.String()
}

func RandomOwner() string {
	return RandomString(6)
}
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomOwner())
}
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}
func RandomCurrency() string {
	list := []string{"EUR", "KZT", "RUB", "USD"}
	return list[rand.Intn(len(list))]
}
