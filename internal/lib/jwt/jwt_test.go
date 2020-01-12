package jwt

import "testing"

func TestJwtClaims_GenerateToken(t *testing.T) {
	claims := Claims{
		UserId:            123,
		BusinessProfileId: 4567,
	}
	tokenString, err := Resolver{
		SecretKey:     "1qaz2wsx",
		ValidDuration: 1000,
	}.GenerateToken(claims)
	if err != nil {
		t.Error(err)
	}
	t.Log(tokenString)
}

func TestResolver_ValidateToken(t *testing.T) {
	claims := Claims{
		UserId:            123,
		BusinessProfileId: 4567,
	}
	tokenString, err1 := Resolver{
		SecretKey:     "1qaz2wsx",
		ValidDuration: 1000,
	}.GenerateToken(claims)
	if err1 != nil {
		t.Error(err1)
	}
	t.Log(tokenString)
	t.Log(claims)

	isValid, claimsOut, err2 := Resolver{
		SecretKey:     "1qaz2wsx",
		ValidDuration: 1000,
	}.ValidateToken(tokenString)
	if err2 != nil {
		t.Error(err2)
	}
	t.Log(isValid)
	t.Log(claimsOut)
}
