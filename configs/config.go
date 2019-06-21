package configs

type appConfig struct {
	StartMessage  string
	PromptMessage string
	EndMessage    string
	QuitCommands  []string
}

const startMessage = "\nWelcome to Event Reporter!"
const promptMessage = "\nEnter load <filename>, help <command>, queue <command> or find <attribute><criteria>."

var quitCommands = []string{"q", "quit", "exit", "e"}
var config appConfig

// Simply playing around with the idea of loadiing configs rather than
// returning raw values per exported function
func init() {
	config.StartMessage = startMessage
	config.PromptMessage = promptMessage
	config.QuitCommands = quitCommands
}

// StartMessage will return the first prompt
func StartMessage() string {
	return config.StartMessage
}

// PromptMessage will return the last message of the program
func PromptMessage() string {
	return config.PromptMessage
}

// QuitCommands is the list of viable commands to end the program
func QuitCommands() []string {
	return config.QuitCommands
}
