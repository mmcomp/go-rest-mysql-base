package auth

import (
	"testing"
)

func TestGetMD5Hash(t *testing.T) {
	hash, err := HashPassword("123456")
	if err != nil {
		t.Errorf("expected no error, got %s", err)
	}
	if !VerifyLaravelHash("123456", hash) {
		t.Errorf("expected true, got false")
	}
}
