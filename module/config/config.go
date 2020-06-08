package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

func  Init() *Config{
	c := new(Config)
	c.Viper = viper.New()
	return  c
}

func (c *Config) setConfig(fileName string, paths ...string) error{

	if fileName == "" {
		return errors.New("empty fileName ")
	}
	for _, path := range paths {
		c.AddConfigPath(path)
	}
	c.SetConfigName(fileName)
	return nil
}

func (c *Config) LoadJsonFileConfig(fileName string, paths  ...string) error {

	err := c.setConfig(fileName, paths...)
	if err != nil{
		return err
	}
	c.SetConfigType("json")
	err = c.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) LoadYamlFileConfig(fileName string, paths ...string) error {
	err := c.setConfig(fileName, paths...)
	if err != nil{
		return err
	}
	c.SetConfigType("yaml")
	err = c.ReadInConfig()
	if err != nil{
		return err
	}
	return  nil
}

func (c *Config) LoadIniFileConfig(fileName string, paths ...string) error {

	err := c.setConfig(fileName, paths...)
	if err != nil{
		return err
	}
	c.SetConfigType("ini")
	err = c.ReadInConfig()
	if err != nil{
		return err
	}
	return  nil
}


func (c *Config) NewOptions(options ...viper.Option) *Config{

	c.Viper =  viper.NewWithOptions(options...)
	return c
}
