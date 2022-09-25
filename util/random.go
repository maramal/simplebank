package util

import (
	"math/rand"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Genera un número aleatorio entre el valor min y el max
func RandomInt(min, max int64) int64 {
	return rand.Int63n(max - min + 1)
}

// Genera una cadena de texto de largo n
func RandomString(n int) string {
	var builder strings.Builder
	alphabetLength := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(alphabetLength)]
		builder.WriteByte(c)
	}

	return builder.String()
}

// Devuelve una cadena de texto como título
func TitleCase(str string) string {
	caser := cases.Title(language.LatinAmericanSpanish)
	return caser.String(str)
}

// Devuelve un Owner aleatorio
func RandomOwner() string {
	return TitleCase(RandomString(6))
}

// Devuelve una cantidad de dinero aleatoria
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Devuelve un tipo de cambio aleatorio
func RandomCurrency() string {
	currencies := []string{"UYU", "USD", "EUR"}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}
