package assets
import (
	"github.com/charmbracelet/lipgloss"
	"fmt"
)

var logo =
` _______  _____   ____   _____ _      _____
|__   __||  __ \ / __ \ / ____| |    |_   _|
   | |   | |  | | |  | | |    | |      | |
   | |   | |  | | |  | | |    | |      | |
   | |   | |__| | |__| | |____| |____ _| |_
   |_|   |_____/ \____/ \_____|______|_____|`

var style = lipgloss.NewStyle().
  Foreground(lipgloss.Color("202")).
  Bold(true)

func RenderLogo() {
	fmt.Println(style.Render(logo + "\n\nVersion: 1.0.0\nA CLI tool for managing tasks"))
}
