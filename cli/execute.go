package cli

import (
	"bufio"
	"fmt"
	"os"

	"../configs"
	"../helpers"
)

// Execute brings up the command line
func Execute() {
	reader := bufio.NewScanner(os.Stdin)
	var text string
	fmt.Println(configs.StartMessage())
	fmt.Println(configs.PromptMessage())

	for !helpers.Include(configs.QuitCommands(), text) {
		reader.Scan()
		text = reader.Text()
	}
}
