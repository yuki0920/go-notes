package util_test

import (
	"testing"
	"yuki0920/go-notes/util"
)

func TestGenerateJwtToken(t *testing.T) {
	issuer := "dummy"

	_, err := util.GenerateJwtToken(issuer)
	if err != nil {
		t.Errorf("GenerateJwtToken error = %v", err)
		return
	}
}

func TestParseJwtWithValidToken(t *testing.T) {
	cookie := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyIn0.l5OzH8D-jhBGpWOaTICi65_Njdgq78TV6t_z-5JymtQ"

	err := util.ParseJwt(cookie)
	if err != nil {
		t.Errorf("ParseJwt error = %v", err)
		return
	}

}

func TestParseJwtWithInvalidToken(t *testing.T) {
	cookie := "invalid"

	err := util.ParseJwt(cookie)
	if err == nil {
		t.Errorf("ParseJwt error invalid token is parsed")
		return
	}
}
