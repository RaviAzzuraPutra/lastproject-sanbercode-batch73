package port_config

import "os"

var PORT string

func Port_Config() {
	PORT = os.Getenv("APP_PORT")
}
