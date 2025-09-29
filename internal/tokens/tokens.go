package tokens

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"
)

const (
	ScopeAuth = "authentication"
)

type Token struct {
	Plaintext string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserID    int       `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     string    `json:"-"`
}

func GenerateToken(userID int, ttl time.Duration, scope string) (*Token, error) {
	// * Generate a random byte array
	// * Encode the byte array to a string (base32 encoding is good for tokens)
	// * Hash the token using SHA-256
	// * Return a Token struct containing the plaintext token, its hash, userID, expiry time, and scope
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
		Scope:  scope,
	}

	emptyBytes := make([]byte, 32)
	_, err := rand.Read(emptyBytes)
	if err != nil {
		return nil, err
	}

	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(emptyBytes)
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]
	return token, nil
}
