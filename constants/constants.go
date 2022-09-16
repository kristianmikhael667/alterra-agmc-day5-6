package constants

import "os"

type Keys struct {
	Key string
}

func SecretKey() string {
	keys := os.Getenv("SECRET_KEY")
	user := Keys{
		Key: keys,
	}

	return user.Key
}
