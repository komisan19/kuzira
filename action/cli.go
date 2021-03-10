package action

import (
	"log"
	"os"
  "fmt"

	"github.com/urfave/cli"
)

func Init(c *cli.Context) {
  df, err := os.Create("Dockerfile")
  if err != nil {
    log.Fatal("Dockerfile is already")
  }
  defer df.Close()
  fmt.Println("Dockerfile is creaed")
}

