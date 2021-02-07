// Copyright 2021 Juan Ismael Vasquez <ismael.juan@gmail.com>.
// All rights reserved. Use of this source code is governed by
// a MIT style license that can be found in the LICENSE file.

package cexio

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

// nonce return the nonce
func nonce() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// toHmac256 encrypt the signature
func toHmac256(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

func hasText(s string) bool {
	return s != ""
}
