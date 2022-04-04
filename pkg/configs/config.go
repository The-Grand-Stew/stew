package configs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type AppDetails struct {
	AppName   string   `json:"appName"`
	Language  string   `json:"language"`
	Database  string   `json:"database"`
	Framework string   `json:"framework"`
	Domains   []string `json:"domains"`
	AppPath   string   `json:"appPath"`
}
type StewConfig struct {
	ProjectName string       `json:"projectName"`
	Apps        []AppDetails `json:"appName"`
}

//CreateConfig : Create the config file for the first time setup. ie for the first service
func (s *StewConfig) CreateConfig() error {
	for _, app := range s.Apps {
		appPath := filepath.Join(app.AppName, ".stew")
		f, err := os.Create(appPath)
		if err != nil {
			return err
		}
		data, err := json.Marshal(app)
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil
}

func (a *AppDetails) LoadAppConfig() (*AppDetails, error) {
	//read from current dir
	currentDir, _ := os.Getwd()
	configPath := filepath.Join(currentDir, ".stew")
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("stew config file doesn't exist. Are you sure this is an app created by stew??")
	}
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(byteValue, &a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AppDetails) UpdateAppConfig() error {
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
