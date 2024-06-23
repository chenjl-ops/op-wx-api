package weichat

import (
	"testing"
)

// TestGetAccessToken
func TestGetAccessToken(t *testing.T) {
	_, err := getAccessToken()
	if err != nil {
		t.Fatal(err)
	}
}
