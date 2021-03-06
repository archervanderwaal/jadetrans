// Copyright (c) 2019 Archer VanderWaal.
//
// A very geeky command line translation tool.
// Source code and other details for the project are available at GitHub:
//
// https://github.com/archervanderwaal/jadetrans
package main

import (
	"flag"
	"fmt"
	"os"

	"strings"

	"github.com/archervanderwaal/jadetrans/config"
	"github.com/archervanderwaal/jadetrans/engine"
	"github.com/archervanderwaal/jadetrans/utils"
	"github.com/aybabtme/rgbterm"
)

const (
	// Version the version of jadetrans.
	Version = "1.0"
	// Usage usage of jadetrans.
	Usage = "Usage of jadetrans: jadetrans <Sentences to be translated> <command>"
	// Logo the logo of jadetrans.
	Logo = `
       __          __   ______                     
      / /___ _____/ /__/_  __/________ _____  _____
 __  / / __ / __  / _ \/ / / ___/ __ / __ \/ ___/
/ /_/ / /_/ / /_/ /  __/ / / /  / /_/ / / / (__  )
\____/\__,_/\__,_/\___/_/ /_/   \__,_/_/ /_/____/
  ::一北::                                 <1.0>
`
)

var (
	version bool
	eng     string
	help    bool
	voice   string
)

func init() {
	flag.BoolVar(&help, "h", false, "Show usage and exit.")
	flag.BoolVar(&version, "v", false, "Show version and exit.")
	flag.StringVar(&eng, "e", "youdao",
		"Set translate engine(Is access to Google translation engine).")
	flag.StringVar(&voice, "voice", "",
		"Set which voice to read aloud. 0 is female voice and 1 is male voice"+
			"(It can only be used on Linux or MacOsx os).")
	flag.Usage = usage
	flag.Parse()
}

func main() {
	words, _ := utils.ParseArgs(os.Args)
	if version {
		showVersion()
	} else if help || len(words) < 1 {
		flag.Usage()
	}

	// trans.
	// TODO google translate engine.
	conf := config.LoadConfig()
	e := engine.NewYoudaoEngine(strings.Join(words, " "), "auto", "auto", voice, conf)
	res, err := e.Query()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(res)
}

func usage() {
	// #00FF00
	logo := rgbterm.FgString(Logo, 0, 255, 0)
	// #FF42E1
	usage := rgbterm.FgString(Usage, 255, 66, 225)
	fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n\n%s\n", logo, usage))
	flag.PrintDefaults()
	os.Exit(0)
}

func showVersion() {
	version := rgbterm.FgString(Version, 0, 255, 0)
	fmt.Println(version)
	os.Exit(0)
}
