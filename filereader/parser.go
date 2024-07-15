package filereader

import (
	"encoding/csv"
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
	registerParser(fileTypeYAML, yamlParser)
	registerParser(fileTypeTOML, tomlParser)
	registerParser(fileTypeCSV, csvParser)
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

// target can be []struct
func csvParser(filePath string, target interface{}) error {
	// check target is [][]string
	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Slice ||
		targetType.Elem().Kind() != reflect.Slice ||
		targetType.Elem().Elem().Kind() != reflect.String {
		return fmt.Errorf("target must be [][]string")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// put records to target
	targetValue := reflect.ValueOf(target).Elem()
	targetValue.Set(reflect.MakeSlice(targetType, 0, 0))
	for _, record := range records {
		recordValue := reflect.ValueOf(record)
		targetValue.Set(reflect.Append(targetValue, recordValue))
	}
	return nil
}
