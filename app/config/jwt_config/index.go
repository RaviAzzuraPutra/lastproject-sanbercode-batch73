package jwt_config

import "os"

var JWT string

func JWT_Config() {
	JWT = os.Getenv("JWT_SECRET")
}
