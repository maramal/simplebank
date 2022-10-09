package util

var currencies = []string{"USD", "EUR", "UYU"}

func IsSupportedCurrency(currency string) bool {
	for _, c := range currencies {
		if c == currency {
			return true
		}
	}

	return false
}
