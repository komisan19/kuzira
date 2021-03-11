package action

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func createDockerfile() {
  df, err := os.Create("Dockerfile")
  if err != nil {
    log.Fatal(err)
  }
  defer df.Close()
}

func Init(c *cli.Context) {
  createDockerfile()
  fmt.Println("Dockerfile is already")
}

func Create(c *cli.Context) {
  rd, err := os.ReadFile("../cookbook/go-dockerfile")
  if err != nil {
    log.Fatal(err)
  }
  err = os.WriteFile("Dockerfile", rd, 0755)
  if err != nil {
    log.Fatal(err)
  }
}
