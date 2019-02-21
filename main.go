// Copyright 2019 Archer VanderWaal. All rights reserved.
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/aybabtme/rgbterm"
	"github.com/archervanderwaal/jadetrans/utils"
)

const (
	Version                = "1.0"
	Usage                  = "Usage of jadetrans: jadetrans <Sentences to be translated> <command>"
	IllegalParametersError = "Illegal parameters error."
	Logo                   = `
       __          __   ______                     
      / /___ _____/ /__/_  __/________ _____  _____
 __  / / __ / __  / _ \/ / / ___/ __ / __ \/ ___/
/ /_/ / /_/ / /_/ /  __/ / / /  / /_/ / / / (__  )
\____/\__,_/\__,_/\___/_/ /_/   \__,_/_/ /_/____/
  ::一北::                                 <1.0>
`
)

var (
	voice bool
	engine string
	help bool
	version bool
)

func init() {
	flag.BoolVar(&help, "help", false, "Show usage and exit")
	flag.BoolVar(&version, "version", false, "Show version and exit")
	flag.StringVar(&engine, "engine", "youdao", "Set translate engine [youdao, google]")
	flag.BoolVar(&voice, "voice", false, "Set up the sound to read aloud")
	flag.Usage = usage
	flag.Parse()
}

func main()  {
	words, _ := utils.ParseArgs(os.Args)
	if version {
		showVersion()
	} else if help {
		flag.Usage()
	} else if len(words) < 1 {
		showIllegalParametersError()
		flag.Usage()
	}
}

func usage() {
	// #00FF00
	logo := rgbterm.FgString(Logo, 0, 255, 0)
	// #FF42E1
	usage := rgbterm.FgString(Usage, 255, 66, 225)
	fmt.Fprintf(os.Stderr, fmt.Sprintf("%s\n\n%s\n", logo, usage))
	flag.PrintDefaults()
}

func showVersion() {
	version := rgbterm.FgString(Version, 0, 255, 0)
	fmt.Println(version)
}

func showIllegalParametersError() {
	// #FF0000
	inputIllegalParametersError := rgbterm.FgString(IllegalParametersError, 255, 0, 0)
	fmt.Fprintf(os.Stderr, inputIllegalParametersError)
}