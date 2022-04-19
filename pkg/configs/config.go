package configs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type AppConfig struct {
	ConfigType string   `json:"configType"`
	AppName    string   `json:"appName"`
	Language   string   `json:"language"`
	Database   string   `json:"database"`
	Framework  string   `json:"framework"`
	Domains    []string `json:"domains"`
}

type StewConfig struct {
	ConfigType  string   `json:"configType"`
	ProjectName string   `json:"projectName"`
	Apps        []string `json:"apps"`
}

//CreateConfig : Create the config file for the first time setup. also for the first service
func (s StewConfig) CreateConfig() error {
	s.ConfigType = "project"
	currentDir, _ := os.Getwd()
	projectConfigPath := filepath.Join(currentDir, ".stew")
	f, err := os.Create(projectConfigPath)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

//LoadConfig : Loads the config file for the project
func (s *StewConfig) LoadConfig() error {
	currentDir, _ := os.Getwd()
	projectConfigPath := filepath.Join(currentDir, ".stew")
	// check for
	if _, err := os.Stat(projectConfigPath); errors.Is(err, os.ErrNotExist) {
		return errors.New("stew config file doesn't exist. Are you sure this is an app created by stew??")
	}
	configFile, err := os.Open(projectConfigPath)
	if err != nil {
		return err
	}
	byteValue, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(byteValue, s)
	if err != nil {
		return err
	}
	return nil
}

//CreateAppConfig : Update the config file for app
func (a *AppConfig) CreateAppConfig() error {
	currentDir, _ := os.Getwd()
	a.ConfigType = "app"
	configPath := filepath.Join(currentDir, a.AppName, ".stew")
	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

//CreateAppConfig : Update the config file for app
func (a *AppConfig) UpdateAppConfig() error {
	currentDir, _ := os.Getwd()
	configPath := filepath.Join(currentDir, ".stew")
	f, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := json.Marshal(a)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}

//LoadConfig : Loads the config file for the app
func (a *AppConfig) LoadAppConfig() error {
	currentDir, _ := os.Getwd()

	//read from current dir
	configPath := filepath.Join(currentDir, ".stew")
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		return errors.New("stew config file doesn't exist. Are you sure this is an app created by stew??")
	}
	configFile, err := os.Open(configPath)
	if err != nil {
		return err
	}
	byteValue, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(byteValue, &a)
	return err
}

func DetectConfigType() (int, []string) {
	var projectConfig StewConfig
	err := projectConfig.LoadConfig()
	if err != nil {
		return -1, []string{}
	}
	if projectConfig.ConfigType == "project" {
		return 1, projectConfig.Apps
	} else if projectConfig.ConfigType == "app" {
		return 0, []string{}
	}
	return -1, []string{}
}
