package pkg

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
)

func TestViper(t *testing.T) {
	vp := viper.New()
	vp.SetConfigName("dev")
	vp.AddConfigPath("./")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	type App struct {
		Env      string `yaml:"env"`
		HTTPPort int    `yaml:"httpPort"`
		Name     string `yaml:"name"`
		IP       string `yaml:"ip"`
	}

	fmt.Println(vp.AllSettings())
	var app App
	err = vp.UnmarshalKey("app", &app)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v\n", app)
}
