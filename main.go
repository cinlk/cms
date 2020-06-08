package main

import (
	"fmt"
	"framework/module/config"
)

type demo struct {

	Address string
	Port int
}
func main() {


	//config := viper.New()
	//
	//config.AddConfigPath("./")
	//config.SetConfigName("config")
	//config.SetConfigType("json")
	//if err := config.ReadInConfig(); err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("%v \n", config)
	//
	//var  d demo
	//fmt.Println(config.GetStringMap("host"))
	//
	//if err := mapstructure.Decode(config.GetStringMap("host"), &d); err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Print(d)

	cf := config.Init()

	cf.AddConfigPath("./")
	cf.SetConfigName("config")

	cf.SetConfigType("json")
	_ = cf.ReadInConfig()

	fmt.Printf("%v \n", cf.Viper)

	var  d demo
	fmt.Println(cf.GetString("appId"))
	if err := cf.UnmarshSection("host", &d); err != nil{
		panic(err)
	}
	fmt.Println(d)


}
