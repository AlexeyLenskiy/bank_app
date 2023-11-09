package util

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < rand.Intn(n); i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() string {
	return cases.Title(language.English, cases.Compact).String(RandomString(15))
}

func RandomEmail() string {
	domen := RandomString(3)
	return RandomString(30) + "@" + domen + "mail.com"
}

func RandomPhone() string {
	phone := RandomInt(10000000, 9999999999)
	return "+" + strconv.Itoa(int(phone))
}

func RandomPin() string {
	pin := RandomInt(100000000, 9999999999)
	return strconv.Itoa(int(pin))
}

func RandomBalance() int64 {
	balance := RandomInt(0, 999999)
	return balance
}

func RandomCurrency() string {
	currencies := []string{EUR, USD, JPY}
	i := rand.Intn(len(currencies))
	return currencies[i]
}

func RandomAmount() int64 {
	amount := RandomInt(0, 1000)
	return amount
}
