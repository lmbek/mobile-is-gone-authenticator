package main

import (
	"crypto/hmac"
	"crypto/sha1"
	_ "embed"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

//go:embed secrets/google_secret.txt
var googleSecret string // Base32 encoded TOTP secret

func main() {
	// Ticker that triggers every 3 seconds
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	var lastCode string

	for {
		// Generate the current TOTP code
		totp, err := generateTOTP(googleSecret)
		if err != nil {
			log.Fatalf("Failed to generate TOTP: %v", err)
		}

		// Only print if the TOTP code has changed
		if totp != lastCode {
			log.Printf("Generated TOTP: %s", totp)
			lastCode = totp
		}

		// Wait for the next tick (3 seconds)
		<-ticker.C
	}
}

func generateTOTP(secret string) (string, error) {
	secretBytes, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", fmt.Errorf("failed to decode secret: %w", err)
	}

	// Calculate the time step (current Unix time divided by 30)
	timestep := time.Now().Unix() / 30

	// Prepare the buffer for HMAC SHA1
	buffer := make([]byte, 8)
	binary.BigEndian.PutUint64(buffer, uint64(timestep))

	// Create HMAC hash
	h := hmac.New(sha1.New, secretBytes)
	h.Write(buffer)
	hash := h.Sum(nil)

	// Extract the dynamic offset and calculate the OTP
	offset := hash[len(hash)-1] & 0x0F
	code := (binary.BigEndian.Uint32(hash[offset:offset+4]) & 0x7FFFFFFF) % 1000000

	// Return the code as a 6-digit string
	return fmt.Sprintf("%06d", code), nil
}
