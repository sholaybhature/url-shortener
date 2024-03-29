package utils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

/*
Mapping between URLs should be unique with no collisions.
- Hash (URL+randomInt) and then base64 encode it
- getUniqueID (DB id or random int) and then base64 encode it
- Store a list of ready-to-use shortened links
*/
func EncodeURL(url string) string {
	rand.Seed(time.Now().UnixNano())
	appendRandInt := rand.Uint64()
	toHash := fmt.Sprint(url, appendRandInt)
	hashedURL := md5.Sum([]byte(toHash))
	// use just 36 bits as we need a link of length 6 only?
	encoded := base64.StdEncoding.EncodeToString(hashedURL[:5])
	// base64 uses 6 bits and we provided 40bits (8*5), it'll add extra 2 bits
	// to get the 7 char which we don't want, so return just till 6 chars
	// also replace /+ with ~. for url safe
	replacer := strings.NewReplacer("/", "~", "+", ".")
	encoded = replacer.Replace(encoded[:6])
	return encoded
}

// no decoding requried as links are unqiue ids? if used db ids to encode
// urls then would have used decoding?
