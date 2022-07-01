package govalidator

import (
	"fmt"
	"io"

	"github.com/spf13/viper"
)

var msgConfig *viper.Viper

func ReadInFile(filename string, fileType ...string) error {
	v := viper.New()
	// v.SetConfigType(fileType)
	if len(fileType) > 0 {
		v.SetConfigType(fileType[0])
	}
	v.SetConfigFile(filename)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	// v.AutomaticEnv()
	// fmt.Println(v.AllKeys())
	// for _, m := range v.AllKeys() {
	// 	fmt.Println(v.GetString(m))
	// }

	msgConfig = v
	return nil
}

func ReadConfig(in io.Reader, fileType string) error {
	v := viper.New()
	v.SetConfigType(fileType)
	if err := v.ReadConfig(in); err != nil {
		return err
	}
	msgConfig = v
	return nil
}

func GetMsgError(s InvalidValidateStruct) string {
	v := msgConfig
	if v == nil {
		return ""
	}

	//* GetBy: StructName.Field.Tag
	cfgName := fmt.Sprintf("%s.%s", s.Namespace, s.Tag)
	if v.InConfig(cfgName) {
		return v.GetString(cfgName)
	}

	//* GetBy: Field.Tag
	//? in Global
	cfgName = fmt.Sprintf("global.%s.%s", s.Field, s.Tag)
	if v.InConfig(cfgName) {
		return v.GetString(cfgName)
	}

	//* GetBy: Tag
	//? in Global
	cfgName = fmt.Sprintf("global.%s", s.Tag)
	if v.InConfig(cfgName) {
		return v.GetString(cfgName)
	}

	return ""
}
