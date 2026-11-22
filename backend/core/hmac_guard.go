// PROJECT NAME: ThitNueaHub-Core-Master // FOLDER PATH: /backend/core/ // FILE NAME: hmac_guard.go
package core

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
)

// ErrMissingSecret indicates TNH_MASTER_SECRET is not configured in the environment.
var ErrMissingSecret = errors.New("security: TNH_MASTER_SECRET environment variable is required")

func loadMasterSecret() ([]byte, error) {
	secret := os.Getenv("TNH_MASTER_SECRET")
	if len(secret) < 16 {
		return nil, ErrMissingSecret
	}
	return []byte(secret), nil
}

// SignPayload generates an L8 Glock HMAC-SHA256 hex signature for an envelope payload.
func SignPayload(payload []byte) (string, error) {
	secret, err := loadMasterSecret()
	if err != nil {
		return "", err
	}
	mac := hmac.New(sha256.New, secret)
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil)), nil
}

// VerifyPayload validates an incoming L8 Glock HMAC-SHA256 hex signature in constant time.
func VerifyPayload(payload []byte, signatureHex string) (bool, error) {
	expectedHex, err := SignPayload(payload)
	if err != nil {
		return false, err
	}
	expectedBytes, err := hex.DecodeString(expectedHex)
	if err != nil {
		return false, err
	}
	actualBytes, err := hex.DecodeString(signatureHex)
	if err != nil {
		return false, nil
	}
	return hmac.Equal(expectedBytes, actualBytes), nil
}
