package token

import "time"

// Maker is an interface for managing tokens
type Maker interface {
	// CreateToken creates a new token for a speicfic username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerfiyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
