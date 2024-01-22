package filereader

import (
	"encoding/json"
	"fmt"
	"os"

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
	registerParser(fileTypeYAML, yamlParser)
	registerParser(fileTypeTOML, tomlParser)
}

func jsonParser(filePath string, target interface{}) error {
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
	_, err := toml.DecodeFile(filePath, target)
	if err != nil {
		return fmt.Errorf("decode toml file error, filePath=%s, err=%s", filePath, err)
	}
	return nil
}
