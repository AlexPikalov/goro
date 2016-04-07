package main

import (
	"flag"
	"fmt"
	"os"
)

var config *Config = &Config{}

func main() {
	composer := NewComposer()

	err := InitConfig(composer)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Println("[Goro Task runner]")
	fmt.Println("Goro file", composer.Config.Gorofile)

	composer.CleanContainer()
	bin, err := composer.ComposeBin()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	if composer.Config.KeepContainer {
		fmt.Println("compiled file: ", bin)
	}

	out, err := composer.RunBin(bin)
	if err != nil {
		fmt.Println(string(out))
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(out))

	defer func() {
		if !composer.Config.KeepContainer {
			composer.CleanContainer()
		}
	}()
}

func InitConfig(c *Composer) error {
	var id string
	flag.StringVar(&c.Config.Gorofile, "gf", DEFAULT_GOROFILE, "Path to goro file. Optional")
	flag.StringVar(&c.Config.Gorofile, "gorofile", DEFAULT_GOROFILE, "Path to goro file. Optional")
	flag.BoolVar(&c.Config.KeepContainer, "keep", false, "Keep container with bin")
	flag.StringVar(&id, "id", "", "Goro id")
	flag.StringVar(&c.Config.LogFile, "log", "gorocli.log", "Log file")
	flag.StringVar(&c.Config.Container, "dst", DEFAULT_CONTAINER, "Container Folder")
	flag.Parse()

	if id != "" {
		c.Config.Id = id
	}

	return nil
}

type Config struct {
	logFile string
}
