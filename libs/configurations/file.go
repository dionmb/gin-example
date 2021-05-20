package configurations

import (
	"fmt"
	"gin_example/app"
	"github.com/spf13/viper"
	"path"
)

func LoadConfig(name string, st interface{})  {
	v := viper.New()
	v.AddConfigPath(path.Join(app.Root, "config"))
	v.SetConfigName(name)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	key := app.Env

	if key == "" {
		panic("Env not set")
	}

	if !v.IsSet(key) {
		panic(fmt.Sprintf("Can not load %s config with %s env", name, key))
	}

	v.UnmarshalKey(key, st)
}
