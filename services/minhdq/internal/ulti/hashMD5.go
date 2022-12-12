package ulti

import (
	"crypto/md5"
	"fmt"
)

func MD5(payload string) (md5Hash string) {
	return fmt.Sprintf("%x", md5.Sum([]byte(payload)))
}
