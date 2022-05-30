package utils

import "time"

func IsExpire(expireAt int64) (isExpire bool) {
	isExpire = time.Now().Unix() > expireAt
	return
}
