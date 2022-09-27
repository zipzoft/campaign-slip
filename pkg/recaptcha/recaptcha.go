package recaptcha

import "time"

const (
	Version = "v3"

	RequestKeyName = "x-mvp"
)

type GoogleRecaptcha interface {
	// Verify verifies the token and returns true if the token is valid.
	Verify(token string) (bool, error)
}

type SiteVerifyResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}
