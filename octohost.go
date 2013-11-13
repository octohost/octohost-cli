package main

import (
	"fmt"
	goopt "github.com/droundy/goopt"
)

// The Flag function creates a boolean flag, possibly with a negating
// alternative.  Note that you can specify either long or short flags
// naturally in the same list.
var amVerbose = goopt.Flag([]string{"-v", "--verbose"}, []string{"--quiet"},
	"output verbosely", "be quiet, instead")

// This is just a logging function that uses the verbosity flags to
// decide whether or not to log anything.
func log(x ...interface{}) {
	if *amVerbose {
		fmt.Println(x...)
	}
}

var install = goopt.Alternatives([]string{"--install", "-i"},
	[]string{"default", "middleman", "sinatra", "harp", "php"},
	"Dockerfile to install")

func main() {
	goopt.Description = func() string {
		return "Install Dockerfile for use with octohost."
	}
	goopt.Version = "1.0"
	goopt.Summary = "octohost Dockerfile installer"
	goopt.Parse(nil)
  switch *install {
    case "default": fmt.Println("Need to pick something."); return
    case "middleman": fmt.Println("wget middleman")
    case "sinatra": fmt.Println("wget sinatra")
    case "harp": fmt.Println("wget harp")
    case "php": fmt.Println("wget php")
  }
	fmt.Printf("Installing the %s Dockerfile.\n", *install)
}
