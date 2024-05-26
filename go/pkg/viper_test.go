package pkg

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
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
	fmt.Println(vp.AllSettings())

	type App struct {
		Env      string `yaml:"env"`
		HTTPPort int    `yaml:"httpPort"`
		Name     string `yaml:"name"`
		IP       string `yaml:"ip"`
	}
	var app App
	err = vp.UnmarshalKey("app", &app)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%+v\n", app)

	quit := make(chan os.Signal, 1)
	ch := make(chan struct{})
	go func() {
		vp.WatchConfig()
		vp.OnConfigChange(func(in fsnotify.Event) {
			err = vp.UnmarshalKey("app", &app)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			ch <- struct{}{}
		})
	}()

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-quit:
			goto End
		case <-ch:
			fmt.Printf("%+v\n", app)
		}
	}

End:
	fmt.Printf("%+v\n", app)
}
