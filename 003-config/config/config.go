package config

import (
	"io/ioutil"
	"log"
	"reflect"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Name      string
	Timestamp string

	Apstra struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port" default:"443"`
		User     string `yaml:"username" default:"admin"`
		Password string `yaml:"password" default:"admin"`
	} `yaml:"apstra"`
}

func New(name string) *Config {
	return &Config{
		Name: name,
	}
}

// dump configuration data in yaml format
func (c *Config) DumpYaml(filename string) error {
	// set timestamp with now()
	now := time.Now().Format(time.ANSIC)
	c.Timestamp = now
	yamlString, err := yaml.Marshal(&c)
	// log.Printf("yamlString = %s", yamlString)

	err = ioutil.WriteFile(filename, yamlString, 0)
	return err
}

// iterate struct elements and fill the default when data was not given
func (c *Config) FillDefaults(data any) {
	element := reflect.ValueOf(data).Elem()
	for i := 0; i < element.NumField(); i++ {
		field := element.Type().Field(i)
		fieldName := field.Name
		fieldDefault := field.Tag.Get("default")
		fieldValue := element.FieldByName(fieldName).Interface()
		// log.Printf("\n  field = %+v, fieldName = %+v, default = %+v, value = %+v", field, fieldName, fieldDefault, fieldValue)
		if fieldValue == "" && fieldDefault != "" {
			element.FieldByName(fieldName).SetString(fieldDefault)
		}
	}
}

// load yaml file and update the configuration data
func (c *Config) LoadYaml(configFile string) {
	dataString, err := ioutil.ReadFile(configFile)
	// log.Printf("dataString = \n%s", dataString)
	if err != nil {
		log.Fatal(err)
	}
	err2 := yaml.Unmarshal(dataString, &c)
	if err2 != nil {
		log.Fatal(err2)
	}

	c.FillDefaults(&c.Apstra)

	// log.Printf("c = %+v", c)
}
