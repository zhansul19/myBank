package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomOwner(t *testing.T) {
	owner:=RandomOwner()

	require.Equal(t,6,len(owner))
}
func TestRandomInt(t *testing.T) {
	randomNum:=RandomInt(1,1000)

	require.GreaterOrEqual(t,randomNum,int64(1))
	require.LessOrEqual(t,randomNum,int64(1000))
}
func TestRandomMoney(t *testing.T) {
	randomNum:=RandomMoney()

	require.GreaterOrEqual(t,randomNum,int64(1))
	require.LessOrEqual(t,randomNum,int64(1000))
}
func TestRandomEmail(t *testing.T) {
	randEmail:=RandomEmail()

	require.Contains(t,randEmail,"@email.com")
	require.LessOrEqual(t,16,len(randEmail))
}

func TestRandomCurrency(t *testing.T) {
	randomCurrency:=RandomCurrency()

	if randomCurrency=="KZT"||randomCurrency=="USD"||randomCurrency=="RUB"||randomCurrency=="EUR" {
		require.NotEmpty(t,randomCurrency)
		
	}
}