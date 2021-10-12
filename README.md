# go-discord-utils

### How to use

```bash
go get -u github.com/bytixo/go-discord-utils
```
```go
package main

import (
   dcutils "github.com/bytixo/go-discord-utils"
  "github.com/bytixo/go-discord-utils/logger"
)

func main() {
	// Cloudflare Cookie
	v := dcutils.GetCfCookie()
	logger.Info(v)

	// Check a token
	z := dcutils.CheckSingleToken("Token")
	if z {
		logger.Info("Valid")
	} else {
		logger.Error("Invalid")
	}
	// Fingerprint
	f := dcutils.GetFingerprint()
	logger.Info(f)
}
