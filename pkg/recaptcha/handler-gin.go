package recaptcha

import (
	"campiagn-slip/config"
	"github.com/gin-gonic/gin"
)

func NewGinHandler() gin.HandlerFunc {
	conf := config.GetConfig()
	verifier := NewGoogleRecaptcha(conf.GoogleRecaptcha.Secret)

	// Validate the secret
	if conf.GoogleRecaptcha.Enabled && conf.GoogleRecaptcha.Secret == "" {
		panic("Please provide a secret key for google recaptcha")
	}

	return func(ctx *gin.Context) {
		if conf.GoogleRecaptcha.Enabled {
			// Ignore the request if called from local cluster
			// Header: X-Forwarded-For
			if ctx.Request.Header.Get("X-Forwarded-For") == "" {
				return
			}

			token := ctx.Request.Header.Get(RequestKeyName)
			if token == "" {
				ctx.AbortWithStatusJSON(400, gin.H{"error": "Authorization"})
				return
			}

			success, err := verifier.Verify(token)
			if err != nil || !success {
				ctx.AbortWithStatusJSON(400, gin.H{"error": "Invalid Authorization"})
				return
			}
		}
	}
}
