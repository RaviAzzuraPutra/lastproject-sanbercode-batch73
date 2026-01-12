package gemini_config

import "os"

var GEMINI_API_KEY string

func Gemini_Config() {
	GEMINI_API_KEY = os.Getenv("GEMINI_API_KEY")
}
