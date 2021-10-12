package main

import "github.com/bytixo/go-discord-utils/logger"

func main() {
	// Cloudflare Cookie
	v := GetCfCookie()
	logger.Info(v)

	// Check a token
	z := CheckSingleToken("Token")
	if z {
		logger.Info("Valid")
	} else {
		logger.Error("Invalid")
	}
	// Fingerprint
	f := GetFingerprint()
	logger.Info(f)
}
