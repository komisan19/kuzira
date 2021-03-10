package main

import (
	"os"

	"github.com/komisan19/kuzira/memu"
)

func main(){
  app := menu.AppConfig()
  app.Run(os.Args)
}
