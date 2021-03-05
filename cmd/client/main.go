package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	//"time"

	log "github.com/sirupsen/logrus"
)

const (
	//PARSER is the tool which used to parse the expression
	PARSER       = "/usr/local/bin/parser"
	COMFIGMAPKEY = "variables"
)

type Variables struct {
	Key      string      `json:"key"`
	Value    string      `json:"value"`
	RawValue interface{} `json:"rawvalue"`
}

type VariablesCM struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func init() {
}

func parseExpression(expression string) (*[]Variables, error) {
	expressionArg := expression

	fmt.Println(expressionArg)
	output, err := exec.Command(PARSER, expressionArg).Output()
	fmt.Println(string(output))
	if err != nil {
		log.Errorf("Parse variables hit error: %+v", err)
		return nil, err
	}

	var variables []Variables
	fmt.Println(strings.ReplaceAll(string(output), "'", "\""))
	err = json.Unmarshal([]byte(strings.ReplaceAll(string(output), "'", "\"")), &variables)
	if err != nil {
		log.Errorf("Unmarshal variables hit error: %+v", err)
		return nil, err
	}

	for i, v := range variables {
		fmt.Println(v.Key)
		fmt.Println(v.RawValue)
		variables[i].Value = fmt.Sprintf("%v", v.RawValue)
		fmt.Println(v.Value)
	}

	return &variables, nil
}

func main() {
	var expressions string

	flag.StringVar(&expressions, "expressions", "", "expression of variables")
	flag.Parse()

	var variables *[]Variables
	var err error
	if expressions != "" {
		variables, err = parseExpression(expressions)
		if err != nil {
			log.Errorf("parse variables hit error: %+v", err)
			os.Exit(1)
		}
	}

	if variables == nil {
		log.Errorf("no variables need to set")
		os.Exit(0)
	}

}
