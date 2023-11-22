package filereader

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type Parser func(filePath string, target interface{}) error

var parserMap = map[fileType]Parser{}

func registerParser(ft fileType, parser Parser) {
	parserMap[ft] = parser
}

func init() {
	registerParser(fileTypeJSON, jsonParser)
}

func jsonParser(filePath string, target interface{}) error {
	if err := checkTarget(target); err != nil {
		return err
	}

	jsonFile, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read config file error, filePath=%s, err=%s", filePath, err)
	}
	err = json.Unmarshal(jsonFile, target)
	if err != nil {
		return fmt.Errorf("unmarshal json file error, filePath=%s, err=%s", filePath, err)
	}
	return nil
}

func yamlParser(filePath string, target interface{}) error {
	if err := checkTarget(target); err != nil {
		return err
	}

	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("read config file error, filePath=%s, err=%s", filePath, err)
	}
	err = yaml.Unmarshal(yamlFile, target)
	if err != nil {
		return fmt.Errorf("unmarshal yaml file error, filePath=%s, err=%s", filePath, err)
	}
	return nil
}

func tomlParser(filePath string, target interface{}) error {
	if err := checkTarget(target); err != nil {
		return err
	}

	_, err := toml.DecodeFile(filePath, target)
	if err != nil {
		return fmt.Errorf("decode toml file error, filePath=%s, err=%s", filePath, err)
	}
	return nil
}

func checkTarget(input interface{}) error {
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("target is not a pointer")
	}
	if v.IsNil() {
		return fmt.Errorf("target is nil")
	}
	return nil
}
