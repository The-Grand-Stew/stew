package commands

func JavaSpringBootInit(directoryPath string) error {
	options := []string{"clean", "package"}
	command := "mvn"
	return ExecCommandWrapper(command, options, directoryPath)
}
