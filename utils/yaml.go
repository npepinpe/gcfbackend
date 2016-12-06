package utils

import (
	"os"

	yaml "github.com/go-yaml/yaml"
)

func ReadYAML(path string, config interface{}) (err error) {
	rawConfig, err := readFile(path)

	if err != nil {
		return
	}

	err = yaml.Unmarshal(rawConfig, config)
	return
}

func ReadEnvironmentYAML(path string, config interface{}, env string) (err error) {
	multiEnv := make(map[string]*rawMessage)
	err = ReadYAML(path, multiEnv)

	if err != nil {
		return
	}

	err = multiEnv[env].Unmarshal(config)
	return
}

func readFile(path string) (buffer []byte, err error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return
	}

	info, err := file.Stat()
	if err != nil {
		return
	}

	buffer = make([]byte, info.Size())
	_, err = file.Read(buffer)
	return
}

// Support for multi environment files by delaying the marshalling of inner YAML structures
type rawMessage struct {
	unmarshal func(interface{}) error
}

func (msg *rawMessage) UnmarshalYAML(unmarshal func(interface{}) error) error {
	msg.unmarshal = unmarshal
	return nil
}

func (msg *rawMessage) Unmarshal(value interface{}) error {
	return msg.unmarshal(value)
}
