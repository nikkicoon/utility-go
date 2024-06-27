package pkg

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateRandomString(length int, lower bool) string {
	var charset string
	if lower {
		charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	} else {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

// GenerateEmail generates a random email address-like string
func GenerateEmail(lower bool) string {
	var sb strings.Builder
	sb.WriteString(GenerateRandomString(10, lower))
	sb.WriteString("@")
	sb.WriteString(GenerateRandomString(5, lower))
	sb.WriteString(".")
	sb.WriteString(GenerateRandomString(3, lower))
	return sb.String()
}

// Generate DSV line generates a string with n fields, leading with a random email address
func GenerateDSVLine(length int) string {
	var sb strings.Builder
	sb.WriteString(GenerateEmail(true))
	sb.WriteString(" ")
	var k int
	var l int
	for i := 1; i < length; i++ {
		if i <= 10 {
			l = i + 3
			if l < 11 {
				k = l
			} else {
				k = i
			}
		}
		sb.WriteString(GenerateRandomString(k, true))
		sb.WriteString(" ")
	}
	return strings.TrimSpace(sb.String())
}

// Generate DSV file
func GenerateDSVFile(length int) string {
	var sb strings.Builder
	for i := 1; i < length; i++ {
		sb.WriteString(GenerateDSVLine(4) + "\n")
	}
	return sb.String()
}
