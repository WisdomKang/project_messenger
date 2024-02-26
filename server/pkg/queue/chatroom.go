package queue

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AmqpBroker struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		UserName string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"amqp_broker"`
}

func handleError(err error) {
	if err != nil {
		log.Fatalf("system error : %s\n", err)
	}
}

func init() {
	// Read configuration properties on yaml file
	filePath, err := os.Getwd()
	handleError(err)

	configData, err := os.ReadFile(strings.Join([]string{filePath, "\\server_config.yaml"}, ""))

	handleError(err)

	var config Config

	err = yaml.Unmarshal(configData, &config)

	handleError(err)

	brokerUrl := strings.Join(
		[]string{
			"amqp://",
			config.AmqpBroker.UserName, ":",
			config.AmqpBroker.Password, "@",
			config.AmqpBroker.Host, ":",
			strconv.Itoa((config.AmqpBroker.Port))}, "")

	fmt.Println(brokerUrl)
}
