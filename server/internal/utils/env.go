package utils

import "os"

func GetPort() string {
	return os.Getenv("port")
}
