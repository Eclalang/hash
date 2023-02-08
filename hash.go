package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// Hashmd5 returns the md5 hash of a string
func Hashmd5(str string) string {
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// Hashsha1 returns the sha1 hash of a string
func Hashsha1(str string) string {
	hash := sha1.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

// Hashsha224 returns the sha224 hash of a string
func Hashsha224(str string) string {
	hash := sha256.Sum224([]byte(str))
	return hex.EncodeToString(hash[:])
}

// Hashsha256 returns the sha256 hash of a string
func Hashsha256(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}

// Hashsha384 returns the sha384 hash of a string
func Hashsha384(str string) string {
	hash := sha512.Sum384([]byte(str))
	return hex.EncodeToString(hash[:])
}

// Hashsha512 returns the sha512 hash of a string
func Hashsha512(str string) string {
	hash := sha512.Sum512([]byte(str))
	return hex.EncodeToString(hash[:])
}

// Hashsha512_224 returns the sha512_224 hash of a string
func Hashsha512_224(str string) string {
	hash := sha512.Sum512_224([]byte(str))
	return hex.EncodeToString(hash[:])
}

// Hashsha512_256 returns the sha512_256 hash of a string
func Hashsha512_256(str string) string {
	hash := sha512.Sum512_256([]byte(str))
	return hex.EncodeToString(hash[:])
}
