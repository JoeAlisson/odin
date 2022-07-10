package env

import "os"

var (
	BindAddress = ":8080"
)

func Load() {
	if os.Getenv("BIND_ADDRESS") != "" {
		BindAddress = ":" + os.Getenv("BIND_ADDRESS")
	}

}
