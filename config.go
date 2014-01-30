package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

type config map[string]string

func (c config) file() (string, error) {
	home := os.Getenv("HOME")
	dir := filepath.Join(home, ".config")
	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
		dir = os.Getenv("APPDATA")
		if dir == "" {
			dir = filepath.Join(home, "Application Data")
		}
	}
	_, err := os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, 0700)
		if err != nil {
			return "", err
		}
	}
	dir = filepath.Join(dir, "milk")
	_, err = os.Stat(dir)
	if err != nil {
		err = os.Mkdir(dir, 0700)
		if err != nil {
			return "", err
		}
	}
	return filepath.Join(dir, "settings.json"), nil
}

func (c config) save() error {
	file, err := c.file()
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(file, b, 0700)
	if err != nil {
		return err
	}
	return nil
}

func (c config) load() error {
	file, err := c.file()
	if err != nil {
		return err
	}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		c["api_key"] = api_key
		c["shared_secret"] = shared_secret
	} else {
		err = json.Unmarshal(b, &c)
		if err != nil {
			return err
		}
	}
	return nil
}
