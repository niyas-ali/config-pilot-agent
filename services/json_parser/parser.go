package json_parser

import (
	"config-pilot-agent/utils/logger"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func JsonToModel(file string, model interface{}) error {
	dir, _ := os.Getwd()
	logger.Println("loading file:", dir, file)
	jsonFile, err := os.Open(file)
	if err != nil {
		return errors.New((err.Error()))
	}
	defer jsonFile.Close()
	value, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(value, &model); err != nil {
		return err
	}
	return nil
}

func JsonToFile(data string, file string) {
	dir, _ := os.Getwd()
	logger.PrintSuccessln("saving to file:", dir, file)
	ioutil.WriteFile(fmt.Sprintf("%s\\%s", dir, file), []byte(data), os.FileMode(0777))
}
func LoadFile(file string) (string, error) {
	dir, _ := os.Getwd()
	logger.Println("loading file:", dir, file)
	jsonFile, err := os.Open(fmt.Sprintf("%s\\%s", dir, file))
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()
	value, _ := ioutil.ReadAll(jsonFile)
	return string(value), nil
}
