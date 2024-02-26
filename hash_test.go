package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"testing"
)

func TestHashmd5(t *testing.T) {
	testStr := "test"
	expect := md5.Sum([]byte(testStr))
	expected := hex.EncodeToString(expect[:])
	actual := Hashmd5(testStr)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestHashsha1(t *testing.T) {
	testStr := "test"
	expect := sha1.Sum([]byte(testStr))
	expected := hex.EncodeToString(expect[:])
	actual := Hashsha1(testStr)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestHashsha224(t *testing.T) {
	testStr := "test"
	expect := sha256.Sum224([]byte(testStr))
	expected := hex.EncodeToString(expect[:])
	actual := Hashsha224(testStr)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestHashsha256(t *testing.T) {
	testStr := "test"
	expect := sha256.Sum256([]byte(testStr))
	expected := hex.EncodeToString(expect[:])
	actual := Hashsha256(testStr)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestHashsha384(t *testing.T) {
	testStr := "test"
	expect := sha512.Sum384([]byte(testStr))
	expected := hex.EncodeToString(expect[:])
	actual := Hashsha384(testStr)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestHashsha512(t *testing.T) {
	testStr := "test"
	expect := sha512.Sum512([]byte(testStr))
	expected := hex.EncodeToString(expect[:])
	actual := Hashsha512(testStr)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestHashsha512_224(t *testing.T) {
	testStr := "test"
	expect := sha512.Sum512_224([]byte(testStr))
	expected := hex.EncodeToString(expect[:])
	actual := Hashsha512_224(testStr)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func TestHashsha512_256(t *testing.T) {
	testStr := "test"
	expect := sha512.Sum512_256([]byte(testStr))
	expected := hex.EncodeToString(expect[:])
	actual := Hashsha512_256(testStr)
	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
