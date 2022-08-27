package util

//supported currency
const (
	USD = "USD"
	KZT = "KZT"
	RUB = "RUB"
)

func IsSupported(currency string)bool{
	switch currency{
	case USD,KZT,RUB:
		return true
	}
	return false
}
