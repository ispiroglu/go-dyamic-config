package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"strings"
)

func main() {
	initViper()
	m := &MethodExecutor{}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			fmt.Println("Exiting application.")
			break
		}

		methodName := viper.GetString("admin")
		invokeMethodByName(m, methodName)
	}
}

func invokeMethodByName(target interface{}, methodName string) {
	value := reflect.ValueOf(target)
	method := value.MethodByName(methodName)
	if !method.IsValid() {
		fmt.Printf("Method %s not found\n", methodName)
		return
	}

	method.Call(nil)
}
