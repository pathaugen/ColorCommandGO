package colorcommandgo
import ()

// Testing this package:
// > go env
//
// Ubuntu:
//
// Update Golang:
// > sudo apt-get update
// > sudo apt-get dist-upgrade
// Remove stock golang:
// > sudo apt-get purge golang*
// > sudo apt autoremove
// > sudo apt-get update
//
// Install golang:
// > wget -P $HOME "https://dl.google.com/go/go1.12.linux-amd64.tar.gz"
// > sudo tar -C /usr/local -xzf $HOME/go1.12.linux-amd64.tar.gz
// > export PATH=$PATH:/usr/local/go/bin:$HOME/go-work/bin
// > export GOPATH="$HOME/go-work:$HOME/go-personal"
// > export GOBIN="$HOME/go-work/bin"
// > source $HOME/.profile
//
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
// >> xcopy /E /I C:\git\ColorCommandGo\colorcommandgo C:\go-personal\src\colorcommandgo && cd C:\go-personal\src\colorcommandgo && go install


// GOROOT=C:\Go
// GOPATH=c:\go-work;c:\go-personal
// GOBIN=c:\go-work\bin


// Options:
// color.red('hello')
// var error = color.red.bold;
// error('Error!')

// Color Variables
var (
  // Internal variables (not available to other packages)
  breakspace      = "\n"

  // Clear all color settings in a string
  Clr             = "\u001b[0m"
  // Brighten color values
  Bold            = "\u001b[1m"
  Bright          = Bold
  // String colors
  Black           = "\u001b[30m"
  Red             = "\u001b[31m"
  Green           = "\u001b[32m"
  Yellow          = "\u001b[33m"
  Blue            = "\u001b[34m"
  Magenta         = "\u001b[35m"
  Cyan            = "\u001b[36m"
  White           = "\u001b[37m"
  // Brightened color volues
  BrightBlack     = Bright + "\u001b[30m"
  BrightRed       = Bright + "\u001b[31m"
  BrightGreen     = Bright + "\u001b[32m"
  BrightYellow    = Bright + "\u001b[33m"
  BrightBlue      = Bright + "\u001b[34m"
  BrightMagenta   = Bright + "\u001b[35m"
  BrightCyan      = Bright + "\u001b[36m"
  BrightWhite     = Bright + "\u001b[37m"
  // String background colors
  BGBlack         = "\u001b[40m"
  BGRed           = "\u001b[41m"
  BGGreen         = "\u001b[42m"
  BGYellow        = "\u001b[43m"
  BGBlue          = "\u001b[44m"
  BGMagenta       = "\u001b[45m"
  BGCyan          = "\u001b[46m"
  BGWhite         = "\u001b[47m"
)

// Detect platform and set color variables
// xxx

func ColorTest( stringToColor string, colorChoice string ) string {
  // "Color \"" + colorName + "\" is not recognized, please choose one of the following colors!",
  stringOutput := breakspace +
    "  Each color shown next to bold version used for even numbered output entries:" + breakspace +
    "  * " + Green +   "green" + Clr + "   / bold " + Bold + Green + "green" + Clr + breakspace +
    "  * " + Yellow +  "yellow" + Clr + "  / bold " + Bold + Yellow + "yellow" + Clr + breakspace +
    "  * " + Red +     "red" + Clr + "     / bold " + Bold + Red + "red" + Clr + breakspace +
    "  * " + Blue +    "blue" + Clr + "    / bold " + Bold + Blue + "blue" + Clr + breakspace +
    "  * " + Magenta + "magenta" + Clr + " / bold " + Bold + Magenta + "magenta" + Clr + breakspace +
    "  * " + Cyan +    "cyan" + Clr + "    / bold " + Bold + Cyan + "cyan" + Clr + breakspace
  return stringOutput
}
