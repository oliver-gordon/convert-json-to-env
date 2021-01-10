package cmd

import (
	"convert-json-to-env/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var (
	openedFile map[string]interface{}
	envVars    []string
)

func ConvertJSON(fileName string, printOnly bool) {
	file, err := ioutil.ReadFile(fileName)
	json.Unmarshal(file, &openedFile)
	if err != nil {
		fmt.Printf("Error reading file: %s", fileName)
		os.Exit(0)
	}
	parseMap(openedFile, "")

	if printOnly {
		printMap(envVars)
	} else {
		WriteToFile(envVars)
	}
}

func parseMap(jsonData map[string]interface{}, parentKey string) {
	for key, val := range jsonData {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			parseMap(val.(map[string]interface{}), key)
		default:
			var newEnv string
			if parentKey == "" {
				newEnv = fmt.Sprintf("%s=%v \n", strings.ToUpper(key), concreteVal)
			} else {
				newEnv = fmt.Sprintf("%s.%s=%v \n", strings.ToUpper(parentKey), strings.ToUpper(key), concreteVal)
			}
			envVars = append(envVars, newEnv)
		}

	}
}

func printMap(envVars []string) {
	for i := 0; i < len(envVars); i++ {
		fmt.Printf("%s", envVars[i])
	}
}

func WriteToFile(dataToWrite []string) {
	envFileName := fmt.Sprintf("%s.env", time.Now().Format("15-04-05"))
	file, err := os.Create(envFileName)
	if err != nil {
		fmt.Println("unable to create new file")
		os.Exit(0)
	}

	defer util.CloseFile(file)

	for _, item := range dataToWrite {
		_, err := file.WriteString(item)
		if err != nil {
			fmt.Println("unable to write to file")
			os.Exit(0)
		}
	}
}
