package funcs

import (
	"fmt"

	"github.com/spf13/viper"
)

func RunConfig() {

	// The viper config is available for the entire module.
	v := viper.GetViper()
	fmt.Println(v.GetString("name"))
	fmt.Println(v.GetString("age"))
	fmt.Println(v.GetString("copyright"))

}
