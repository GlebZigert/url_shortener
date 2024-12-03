package services

import "testing"

func TestGenerateRandomString(t *testing.T) {
	if short := generateRandomString(8); len(short) != 8 {
		t.Fatalf("fail with length")
	}
}
