package main

import (
	"github.com/koushilab/MRUAppTest/internal/routers"
	"github.com/labstack/gommon/log"
)

func main() {

	log.Info("...Server Started..")

	routers.StartServer()
}
