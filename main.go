package main

import (
	"errors"

	"github.com/ttmday/go-logger-color/src/logger"
)

func main() {
	logger.Info().Println("INICIALIZE APP")

	logger.Error().Printf("ERROR INICIALIZE APP %v\n", errors.New("HA OCURRIDO UN ERROR"))

	logger.Info().Println("INICIALIZE APP")

	logger.Success().Println("SERVER RESPONSE SUCCESS")
}
