package wrapper

type CommandLine interface {
	Execute() error
	CommandBuild() ([]string, error)
}

