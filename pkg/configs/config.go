package configs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type StewConfig struct {
	AppName   string   `json:"appName"`
	Language  string   `json:"language"`
	Database  string   `json:"database"`
	Framework string   `json:"framework"`
	Domains   []string `json:"domains"`
	AppPath   string   `json:"appPath"`
}

func (s *StewConfig) WriteToConfigFile() error {
	currentDir, _ := os.Getwd()
	appPath := filepath.Join(currentDir, s.AppName, ".stew")
	f, err := os.Create(appPath)
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
func LoadConfig() (*StewConfig, error) {
	config := StewConfig{}
	//read from current dir
	currentDir, _ := os.Getwd()
	configPath := filepath.Join(currentDir, ".stew")
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("stew config file doesn't exist. Might be a good idea to add manually")
	}
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	byteValue, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (s *StewConfig) UpdateConfig() error {
	currentDir, _ := os.Getwd()
	configPath := filepath.Join(currentDir, ".stew")
	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		return errors.New("stew config file doesn't exist. Might be a good idea to add manually")
	}
	f, err := os.Create(configPath)
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
