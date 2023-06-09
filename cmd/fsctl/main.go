package main

import (
	"flag"
	"os"

	"github.com/no-src/fsctl/command"
	"github.com/no-src/fsctl/internal/about"
	"github.com/no-src/fsctl/internal/version"
	"github.com/no-src/log"
)

func main() {
	if c := run(); c != 0 {
		os.Exit(c)
	}
}

const errCode = 1

func run() (code int) {
	defer log.Close()

	var (
		printVersion bool
		printAbout   bool
		conf         string
	)
	flag.BoolVar(&printVersion, "v", false, "print the version info")
	flag.BoolVar(&printAbout, "about", false, "print the about info")
	flag.StringVar(&conf, "conf", "", "the path of config file")
	flag.Parse()

	if printVersion {
		version.PrintVersion("fsctl")
		return
	}

	if printAbout {
		about.PrintAbout()
		return
	}

	if len(conf) == 0 {
		log.Info("please specify the config file by -conf flag")
		code = errCode
		return
	}

	if err := command.Exec(conf); err != nil {
		code = errCode
		log.Error(err, "execute commands failed")
	} else {
		log.Info("execute commands successfully")
	}
	return
}
