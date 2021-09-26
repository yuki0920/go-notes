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
	cookie := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzI3Mjk3NjEsImlzcyI6InVzZXIifQ.9NlQvzKYioB2um6NvB_ZpdKf9og5nRDb9oUzNAjkohk"

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
