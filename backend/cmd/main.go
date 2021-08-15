package cmd

import (
	"fmt"

	"github.com/ITK13201/golang-nextjs-template/backend/config"
	"github.com/ITK13201/golang-nextjs-template/backend/infrastructure"
)

func Run() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	router := infrastructure.InitRouter(cfg)
	if err != nil {
		panic(err)
	}

	router.Run(fmt.Sprintf(":%d", (*cfg).Port))
}
