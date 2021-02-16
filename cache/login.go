package cache

import "time"


// Login 在redis里存放1h的token
func Login(username, token string) error {
	return Get().Set(GetCtx(), username, token, time.Hour * 1).Err()
}
