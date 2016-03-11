package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/AlexPikalov/goro"
	"github.com/AlexPikalov/goro/utils"
)

const DEFAULT_GOROFILE = "./gorofile.go"
const LAUNCHER = "./launcher/launcher.go"
const DEFAULT_CONTAINER = "./containers"

type Composer struct {
	Config      ComposerConfig
	TaskManager goro.TaskManager
	Logger      *log.Logger
}

func NewComposer() *Composer {
	c := &Composer{}
	c.Config.Id = fmt.Sprintf("%d", time.Now().UnixNano())
	c.Config.Container = DEFAULT_CONTAINER
	l, err := newLogger(c.Config.LogFile, "")
	if err != nil {
		panic("cannot create composer logger because of " + err.Error())
		return nil
	}
	c.Logger = l
	return c
}

func (c *Composer) newContainer() error {
	return os.MkdirAll(c.GetContainerPath(), os.ModePerm)
}

func (c *Composer) CleanContainer() error {
	return os.RemoveAll(c.GetContainerPath())
}

func (c *Composer) GetContainerPath() string {
	return fmt.Sprintf("%s/%s", c.Config.Container, c.Config.Id)
}

func (c *Composer) ComposeBin() (string, error) {
	fmt.Println("Composing target bin")
	err := c.newContainer()
	if err != nil {
		return "", err
	}

	// copy gofofile.go to the container
	goroFile, err := filepath.Abs(c.Config.Gorofile)
	if err != nil {
		return "", err
	}

	destGoroFile, err := filepath.Abs(c.GetContainerPath() + "/gorofile.go")
	if err != nil {
		return "", err
	}

	err = utils.FS.CopyFile(goroFile, destGoroFile)
	if err != nil {
		return "", err
	}

	// copy launcher to the container
	c.GetLauncher()
	launcherFile, err := filepath.Abs(LAUNCHER)
	if err != nil {
		return "", err
	}

	destLauncher, err := filepath.Abs(c.GetContainerPath() + "/main.go")
	if err != nil {
		return "", err
	}

	err = utils.FS.CopyFile(launcherFile, destLauncher)
	if err != nil {
		return "", err
	}

	err = c.compileBin()
	if err != nil {
		return "", err
	}

	return c.GetContainerPath() + "/goro", nil
}

func (c *Composer) compileBin() error {
	fmt.Println("building...")
	cmd := exec.Command("go", "build", "-o", "goro")
	cmd.Dir = c.GetContainerPath()
	_, err := cmd.Output()
	return err
}

func (c *Composer) RunBin(path string) ([]byte, error) {
	cmd := exec.Command(path)
	return cmd.Output()
}

func (c *Composer) GetLauncher() (string, error) {
	importContext := build.Default
	pack, err := importContext.ImportDir("github.com/AlexPikalov/goro/launcher", build.FindOnly)
	if err != nil {
		return "", err
	}
	fmt.Println(pack.GoFiles)
	return "", nil
}

func newLogger(logfile, prefix string) (*log.Logger, error) {
	var lf *os.File
	if logfile == "" {
		lf = os.Stdout
	} else {
		path, err := filepath.Abs(logfile)
		if err != nil {
			return nil, err
		}
		lf, err = os.Create(path)
		if err != nil {
			return nil, err
		}
	}
	return log.New(lf, prefix, log.Lshortfile), nil
}

type ComposerConfig struct {
	Id            string
	Gorofile      string
	LogFile       string
	Container     string
	KeepContainer bool
}
