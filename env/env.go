package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type varsT struct {
	PORT         string
	PORT_INT     int
	TOKEN_SECRET string
}

var vars = varsT{}
var initialized = false

func envLog(key string, msg string) {
	log.Printf("[ENV] (%s) %s\n", key, msg)
}

func Vars() varsT {
	if initialized {
		return vars
	}

	// ---- Load file
	godotenv.Load(".env")

	// ---- PORT
	port := os.Getenv("PORT")
	portInt, err := strconv.Atoi(port)

	if err != nil {
		envLog("PORT", "invalid/missing PORT. defaulting to 4321.")
		vars.PORT = "4321"
		vars.PORT_INT = 4321
	} else {
		vars.PORT = port
		vars.PORT_INT = portInt
	}

	// ---- TOKEN_SECRET
	tokenSecret := os.Getenv("TOKEN_SECRET")
	if len(tokenSecret) < 12 {
		envLog("TOKEN_SECRET", "invalid/missing TOKEN_SECRET")
		log.Fatal("bad .env file. shutting down.")
	}

	// ---- return
	initialized = true
	log.Println("[ENV] vars loaded successfully")
	return vars
}
