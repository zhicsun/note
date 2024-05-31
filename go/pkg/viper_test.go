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

type App struct {
	Env      string `yaml:"env"`
	HTTPPort []int  `yaml:"httpPort"`
	Name     string `yaml:"name"`
	IP       string `yaml:"ip"`
}

func TestViperSave(t *testing.T) {
	vp, err := GetViper()
	if err != nil {
		t.Log(err)
	}

	vp.Set("app.env", "prod")
	err = vp.WriteConfigAs("./dev_bak.yaml")
	if err != nil {
		t.Log(err)
		return
	}
}

func TestViperGetSet(t *testing.T) {
	vp, err := GetViper()
	if err != nil {
		t.Log(err)
	}

	vp.SetDefault("app.ip", "192.168.0.1")
	if vp.IsSet("app.httpPort") {
		t.Log(vp.GetIntSlice("app.httpPort"))
	}

	if vp.IsSet("app") {
		t.Log(vp.GetStringMap("app"))
	}

	if vp.IsSet("app.ip") {
		t.Log(vp.GetString("app.ip"))
	}

	t.Log(vp.AllSettings())
}

func TestViperUnmarshal(t *testing.T) {
	vp, err := GetViper()
	if err != nil {
		t.Log(err)
	}

	var app App
	err = vp.UnmarshalKey("app", &app)
	if err != nil {
		t.Log(err)
		return
	}
	t.Logf("%+v", app)
}

func TestViperOnchange(t *testing.T) {
	vp, err := GetViper()
	if err != nil {
		t.Log(err)
	}

	var app App
	quit := make(chan os.Signal, 1)
	ch := make(chan struct{})
	go func() {
		vp.WatchConfig()
		vp.OnConfigChange(func(in fsnotify.Event) {
			err = vp.UnmarshalKey("app", &app)
			if err != nil {
				t.Log(err)
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
			t.Logf("%+v\n", app)
		}
	}

End:
	fmt.Printf("%+v\n", app)
}

func GetViper() (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName("dev")
	vp.AddConfigPath("./")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return vp, nil
}
