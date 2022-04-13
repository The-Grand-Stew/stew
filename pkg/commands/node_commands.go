package commands

import "os"

// 1. clone the repo
// 2. run npm install in the root directory

// 3. to scaffold run npm run scaffold <appname> <path to app>

// 4. To add models use run npm run addModel <modelName> <appname> <path to app>

// 5. To add utils like postgres, mongo  run npm run addUtil <utilName> <appname> <path to app>

func NpmInstall(directoryPath string) error {
	currentDir, _ := os.Getwd()
	os.Chdir(directoryPath)
	options := []string{"install"}
	err := ExecCommand("npm", options, true)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	return nil

}

func npmRun(directoryPath string, options []string) error {
	currentDir, _ := os.Getwd()
	os.Chdir(directoryPath)
	options = append(options, directoryPath)
	runoptions := []string{"run"}
	runoptions = append(runoptions, options...)
	err := ExecCommand("npm", options, true)
	if err != nil {
		return err
	}
	os.Chdir(currentDir)
	return nil
}

func NpmRunModel(directoryPath, modelName string) error {
	options := []string{"addModel", modelName}
	err := npmRun(directoryPath, options)
	return err

}

func NpmRunUtils(directoryPath, utilName string) error {
	options := []string{"addUtil", utilName}
	err := npmRun(directoryPath, options)
	return err
}
