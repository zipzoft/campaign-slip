package recaptcha

import (
	"encoding/json"
	"net/http"
	"net/url"
)

var _ GoogleRecaptcha = (*Verifier)(nil)

type Verifier struct {
	secret string
}

// Verify implements GoogleRecaptcha
func (recaptcha *Verifier) Verify(token string) (bool, error) {
	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", url.Values{
		"secret":   {recaptcha.secret},
		"response": {token},
	})
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	var result SiteVerifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Success, nil
}

func NewGoogleRecaptcha(secret string) GoogleRecaptcha {
	return &Verifier{
		secret: secret,
	}
}
