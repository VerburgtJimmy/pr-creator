package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "creat"
	app.Usage = "Creates front and back-end project with one command"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "extra",
			Value: "",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "pr",
			Usage: "Creates a project with a given name and platform",
			Subcommands: []*cli.Command{
				{
					Name:  "ang",
					Usage: "Create Angular project with Laravel backend",
					Flags: myFlags,
					Action: func(c *cli.Context) error {
						prName := c.String("name")

						if prName == "" {
							log.Fatal("Needs a name as first argument!")
						}

						prExtra := c.String("extra")

						cmd := exec.Command("sh", "-c", "ng new "+prName+" "+prExtra)
						var stdoutBuf, stderrBuf bytes.Buffer
						cmd.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
						cmd.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)
						err := cmd.Run()
						if err != nil {
							log.Fatal(err)
						}

						outStr, errStr := stdoutBuf.String(), stderrBuf.String()
						fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

						fmt.Println("Your Angular project: " + prName + " is created succesfully!")

						addLaravel(prName + "-API")
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

//Function to create the laravel project based on user input
func addLaravel(name string) {
	cmdL := exec.Command("sh", "-c", "composer create-project laravel/laravel "+name)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmdL.Stdout = io.MultiWriter(os.Stdout, &stdoutBuf)
	cmdL.Stderr = io.MultiWriter(os.Stderr, &stderrBuf)

	errL := cmdL.Run()
	if errL != nil {
		log.Fatal(errL)
	}

	outStrL, errStrL := stdoutBuf.String(), stderrBuf.String()
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStrL, errStrL)

	fmt.Println("Your Laravel project: " + name + " is created succesfully!")
}
