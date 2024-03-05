package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// https://cli.urfave.org/v2/examples/greet/
func main() {
    app := &cli.App{
        Name:  "greet",
        Usage: "fight the loneliness!",
        Action: func(*cli.Context) error {
            fmt.Println("Hello friend!")
            return nil
        },
    }

    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}
