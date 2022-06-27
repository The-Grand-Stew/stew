package javaspringboot

import (
	"io/ioutil"
	"os"
	"path/filepath"
	templates "stew/pkg/templates/javaspringboot"
	"strings"
)

func AddDB(template string, dbDriverProtocol string) error {
	// TODO: Proper error handling
	currentDir, _ := os.Getwd()
	// Edit pom.xml
	pomFilePath := filepath.Join(currentDir, "pom.xml")
	content, _ := ioutil.ReadFile(pomFilePath)
	pomFile, _ := os.Create(pomFilePath)
	defer pomFile.Close()
	pomFile.WriteString(strings.Replace(string(content), templates.DatabasePlaceholder, templates.JavaSpringBootJpaLibraryTemplate+template, 1))
	// Edit application.properties
	applicationPropertiesFilePath := filepath.Join(currentDir, "src/main/resources", "application.properties")
	content, _ = ioutil.ReadFile(applicationPropertiesFilePath)
	applicationPropertiesFile, _ := os.Create(applicationPropertiesFilePath)
	defer applicationPropertiesFile.Close()
	applicationPropertiesFile.WriteString(string(content) + strings.Replace(templates.JavaSpringBootApplicationPropertiesBDConnection, templates.DatabaseDriverProtocolPlaceholder, dbDriverProtocol, 1))
	return nil
}

func AddPostgres() error {
	return AddDB(templates.JavaSpringBootPostgresLibraryTemplate, "postgresql")
}

func AddMySql() error {
	return AddDB(templates.JavaSpringBootMySqlLibraryTemplate, "mysql")
}
