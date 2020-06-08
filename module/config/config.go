package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type Config struct {

	*viper.Viper
}




func (c *Config) DataBase(){

}


func  Init() Config{


	c := new(Config)
	c.Viper = viper.New()
	return  *c
}


func (c *Config) LoadJsonConfig(json string) {

	c.SetConfigType("json")

}


func (c *Config) UnmarshSection(section string, mode interface{}) error{

	//fmt.Println(c.Get(section))
	return mapstructure.Decode(c.GetStringMap(section), mode)


}

func (c *Config) NewOptions(options ...viper.Option){

	c.Viper =  viper.NewWithOptions(options...)
}



