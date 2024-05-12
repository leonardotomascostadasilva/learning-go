package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Features map[string]bool `yaml:"features"`
}

func main() {

	CONNECTION_STRING := "CONNECTION_STRING"

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar arquivo .env", err)
	}

	connectionString := os.Getenv(CONNECTION_STRING)

	if connectionString == "" {
		fmt.Println("A variavel de ambiente connection string não está definida")
	} else {
		fmt.Println("O valor da variavel connection string é :", connectionString)
	}

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo YAML: %v", err)
	}

	// Decodificar o conteúdo YAML no struct Config
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Erro ao decodificar o YAML: %v", err)
	}

	// Acessar as feature flags
	if featureX, ok := config.Features["feature_x"]; ok {
		fmt.Println("Feature X está ativada:", featureX)
	} else {
		fmt.Println("Feature X não está definida no arquivo de configuração.")
	}

	if featureY, ok := config.Features["feature_y"]; ok {
		fmt.Println("Feature Y está ativada:", featureY)
	} else {
		fmt.Println("Feature Y não está definida no arquivo de configuração.")
	}
}
