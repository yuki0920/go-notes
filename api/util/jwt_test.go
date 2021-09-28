package util

import (
	"testing"
)

func TestGenerateJwtToken(t *testing.T) {
	issuer := "dummy"

	_, err := GenerateJwtToken(issuer)
	if err != nil {
		t.Errorf("GenerateJwtToken error = %v", err)
		return
	}
}

func TestParseJwtWithValidToken(t *testing.T) {
	cookie := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ"

	err := ParseJwt(cookie)
	if err != nil {
		t.Errorf("ParseJwt error = %v", err)
		return
	}

}

func TestParseJwtWithInvalidToken(t *testing.T) {
	cookie := "invalid"

	err := ParseJwt(cookie)
	if err == nil {
		t.Errorf("ParseJwt error invalid token is parsed")
		return
	}
}
