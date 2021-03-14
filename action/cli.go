package action

import (
	"embed"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

//go:embed cookbook/*
var cookbook embed.FS

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
	switch os.Args[2] {
	case "go":
		rd, err := cookbook.ReadFile("cookbook/go-dockerfile")
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile("Dockerfile", rd, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("create Dockerfile for Golang")
	case "python":
		rd, err := cookbook.ReadFile("cookbook/go-dockerfile")
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile("Dockerfile", rd, 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("create Dockerfile for Golang")
	default:
		fmt.Println("support language is Golang and Python")
	}
}
