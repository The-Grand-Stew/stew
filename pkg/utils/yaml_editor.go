package utils

import (
	"fmt"
	commands "stew/pkg/commands"
)

// type Map map[string]interface{}

var ExtensionMap = map[string]string{"nodejs": ".js", "go": ".go", "java": ".java", "packagejson": ".json"}

// func (m Map) M(s string) Map {
// 	return m[s].(map[string]interface{})
// }

// func (m Map) S(s string) string {
// 	return m[s].(string)
// }

// func readYml(filename string) (Map, error) {
// 	buf, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var body Map

// 	err = yaml.Unmarshal(buf, &body)
// 	if err != nil {
// 		return nil, fmt.Errorf("in file %q: %v", filename, err)
// 	}
// 	return body, nil
// }

// func updateYamlContent(yamlMap Map, prop string, value string) Map {
// 	var updated map[string]interface{}
// 	// marshalled, err := yaml.Marshal(value)
// 	fmt.Println(value)
// 	yaml.Unmarshal([]byte(value), &updated)
// 	fmt.Println(updated)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	yamlMap[prop] = updated
// 	return yamlMap
// }

// func saveUpdatedYaml(filename string, yamlData Map) {
// 	d, err := yaml.Marshal(&yamlData)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = ioutil.WriteFile(filename, d, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func UpdateYmlContents(filename string, property string, value string) {
// 	c, err := readYml(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	updatedYaml := updateYamlContent(c, property, value)
// 	saveUpdatedYaml(filename, updatedYaml)
// 	fmt.Println("saved updated yaml")

// }

func UpdateYmlContents(filename string, property string, value string) {
	setString := "." + property + " +=" + value + ""
	UpdateYmlFile(setString, filename)
	// fmt.Println("Updated:" + filename + " contents")
	// options := []string{"-i", setString, filename}
	// err := commands.ExecCommand("yq", options, true)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("saved updated yaml")
	// }
}

func UpdateYmlFromRoot(filename string, property string, value string, operator string) {
	setString := property + " " + operator + " " + value
	UpdateYmlFile(setString, filename)
	// var options []string

	// options = []string{"-i", setString, filename}

	// err := commands.ExecCommand("yq", options, true)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("saved updated yaml")
	// }
}

func UpdateYmlArray(filename string, property string, value string) {
	setString := "." + property + " +=" + value
	UpdateYmlFile(setString, filename)
	// options := []string{"-i", setString, filename}
	// err := commands.ExecCommand("yq", options, true)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("saved updated yaml")

	// }
}

func UpdateYmlFile(setString string, filename string) {
	options := []string{"-i", setString, filename}
	err := commands.ExecCommand("yq", options, true)
	if err != nil {
		fmt.Println(err)
	}
	// else {
	// 	fmt.Println("saved updated yaml")
	// }
}
