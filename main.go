package main
import (
  "fmt"
	"os"
	"os/exec"
  c "colorcommandgo"
)

var debug = false // Set to true to debug the question 2 output

// ========== START: Console Splash ========== ========== ========== ==========
var appinfo = `
  ` + c.Blue + `======================================================================` + c.Bold + c.Cyan + `
  ,---.     |              ,---.                             |,---.,---.
  |    ,---.|    ,---.,---.|    ,---.,-.-.,-.-.,---.,---.,---||  _.|   |
  |    |   ||    |   ||    |    |   || | || | |,---||   ||   ||   ||   |
  `+"`"+`---'`+"`"+`---'`+"`"+`---'`+"`"+`---'`+"`"+`    `+"`"+`---'`+"`"+`---'`+"`"+` ' '`+"`"+` ' '`+"`"+`---^`+"`"+`   '`+"`"+`---'`+"`"+`---'`+"`"+`---'` + c.Clr + `
  ` + c.Blue + `======================================================================` + c.Clr + `
  ` + c.Cyan + `https://github.com/` + c.Yellow + `pathaugen` + c.Cyan + `/ColorCommandGO` + c.Clr + `
`
// ========== END: Console Splash ========== ========== ========== ==========

func main() {
  // Clear Screen
  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()

  // Output Simplification
  breakspace := ("\n")
	breakline := ( breakspace + c.Blue + "  ======================================================================" + c.Clr + breakspace )

  fmt.Print( appinfo )
  fmt.Print( breakspace )

  fmt.Print( "  ColorCommandGO Usage and Features:" )
  fmt.Print( c.ColorTest( "string", "red" ) )

  fmt.Print( breakspace )
  fmt.Print( breakline )
}
