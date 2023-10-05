package Text

import (
	"fmt"
)

func RGBText(redText, greenText, blueText string) {
	fmt.Println("\033[31m" + redText + "\033[0m")
	fmt.Println("\033[32m" + greenText + "\033[0m")
	fmt.Println("\033[34m" + blueText + "\033[0m")
}

func ColourText(colourCode int, text string) {
	fmt.Println("\033[" + fmt.Sprintf("%d", colourCode) + "m" + text + "\033[0m")
}
