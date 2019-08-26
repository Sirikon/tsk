package application

// Command .
type Command struct {
	Name        string
	Path        string
	Runnable    bool
	SubCommands []*Command
}
