package colorcommandgo
import ()

// Testing this package:
// > go env
//
// Ubuntu:
// Update Golang:
// > sudo apt-get update
// > sudo apt-get dist-upgrade
// Remove stock golang:
// > sudo apt-get purge golang*
// > sudo apt autoremove
// > sudo apt-get update
// Install golang:
// > wget -P $HOME "https://dl.google.com/go/go1.12.linux-amd64.tar.gz"
// > sudo tar -C /usr/local -xzf $HOME/go1.12.linux-amd64.tar.gz
// > export PATH=$PATH:/usr/local/go/bin:$HOME/go-work/bin
// > export GOPATH="$HOME/go-work:$HOME/go-personal"
// > export GOBIN="$HOME/go-work/bin"
// > source $HOME/.profile
// Bash in Windows (Windows Subsystem for Linux: Ubuntu/bash):
// > cd /mnt/c/git
// > cp -R /mnt/c/git/ColorCommandGo/colorcommandgo $HOME/go-personal/src/colorcommandgo
// > cd /mnt/c/git/ColorCommandGo
// > go install
// > ColorCommandGO
//
// Windows:
// > setx GOPATH "c:\go-work;c:\go-personal"
// > setx GOBIN "c:\go-work\bin"
//
// With GOBIN properly set, you can "go install" from any path:
// > go install
// Binary then available straight from command line, stored in GOBIN location.
//
// Package Installation:
// > xcopy /E /I C:\git\ColorCommandGo\colorcommandgo C:\go-personal\src\colorcommandgo
// > cd C:\go-personal\src\colorcommand
// > go install
// > rm -r c:\go-personal\src\colorcommandgo

// Color Variables
// Clear all color settings in a string
var Clr         = "\u001b[0m"
// Enhance color values
var Bold				= "\u001b[1m"
// String colors
var Black       = "\u001b[30m"
var Red         = "\u001b[31m"
var Green       = "\u001b[32m"
var Yellow			= "\u001b[33m"
var Blue				= "\u001b[34m"
var Magenta     = "\u001b[35m"
var Cyan				= "\u001b[36m"
var White       = "\u001b[37m"
// Background colors
var BlackBG     = "\u001b[40m"
var RedBG       = "\u001b[41m"
var GreenBG     = "\u001b[42m"
var YellowBG		= "\u001b[43m"
var BlueBG			= "\u001b[44m"
var MagentaBG   = "\u001b[45m"
var CyanBG			= "\u001b[46m"
var WhiteBG     = "\u001b[47m"

// Detect platform and set color variables
// xxx

func ColorTest() string {
  return "color test"
}
