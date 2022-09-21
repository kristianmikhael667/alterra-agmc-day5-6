package util

import "os"

func GetEnv(key, fallback string) string {
	var (
		val       string
		isisExist bool
	)
	val, isisExist = os.LookupEnv(key)
	if !isisExist {
		val = fallback
	}

	return val
}
