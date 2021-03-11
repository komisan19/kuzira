package menu

import (
	"github.com/komisan19/kuzira/action"
	"github.com/urfave/cli"
)

func AppConfig() *cli.App {
  app := cli.NewApp()
  config(app)
  app.Commands = commands()
  return app
}

func config(app *cli.App){
  app.Name = "kuzira"
  app.Usage = "kuzira can automatic creation Dockerfile"
  app.Version = "0.0.1"
}

func commands() []cli.Command{
  return []cli.Command{
    {
      Name: "init",
      Usage: "create on empty Dockerfile",
      Action: action.Init,
    },
    {
      Name: "create",
      Usage: "create Dockerfile for each language",
      Action: action.Create,
    },
  }
}
