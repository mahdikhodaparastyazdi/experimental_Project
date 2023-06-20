package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sort"
)

func generateString(length int) (string, error) {
	temp := make([]byte, length)
	_, err := rand.Read(temp)
	if err != nil {
		return "", err
	}
	randomString := base64.URLEncoding.EncodeToString(temp)
	return randomString, nil
}

func hash(s string) string {
	hashed := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hashed)
}

var maxNumber = 50

func main() {
	var hashedStrings []string
	var i = 0
	for i < maxNumber {
		randomString, err := generateString(10)
		if err == nil {
			hashedString := hash(randomString)
			if hashedString[len(hashedString)-3:] == "000" {
				hashedStrings = append(hashedStrings, hashedString)
				i++
			}
			sort.Slice(hashedStrings, func(i, j int) bool {
				sumOFI := 0
				sumPfJ := 0

				for _, char := range hashedStrings[i] {
					sumOFI += int(char)
				}

				for _, char := range hashedStrings[j] {
					sumPfJ += int(char)
				}
				if sumOFI < sumPfJ {
					return true
				} else {
					return false
				}
			})
		}
	}
	for i, hashedString := range hashedStrings {
		fmt.Printf("Hashed:%d %s \n", i+1, hashedString)
	}
}
