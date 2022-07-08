package govalidator

import (
	"fmt"
	"reflect"
	"strings"
)

type TNameMsgError string
type ValidateStructErrors []InvalidValidateStruct

type InvalidValidateStruct struct {
	Namespace   string       `json:"-"`
	StructField string       `json:"-"`
	Field       string       `json:"field"`
	Tag         string       `json:"valid"`
	ActualTag   string       `json:"detail"`
	Param       string       `json:"param"`
	Value       interface{}  `json:"-"`
	Kind        reflect.Kind `json:"-"`
	Type        reflect.Type `json:"-"`
	MsgError    string       `json:"massage"`
}

func (v *InvalidValidateStruct) Error() string {
	if v.MsgError == "" {
		v.SetMsgError()
	}
	return v.MsgError
}

func (v *InvalidValidateStruct) SetMsgError() {
	if msgErr := GetMsgError(*v); msgErr != "" {
		v.MsgError = msgErr

	} else if msgErr := GetMsgError(*v); msgErr != "" {
		v.MsgError = msgErr

	}

	if v.MsgError != "" {
		v.MsgError = strings.ReplaceAll(v.MsgError, `${param}`, v.Param)
		v.MsgError = strings.ReplaceAll(v.MsgError, `${field}`, v.Field)
		v.MsgError = strings.ReplaceAll(v.MsgError, `${valid}`, v.Tag)
		v.MsgError = strings.ReplaceAll(v.MsgError, `${detail}`, v.ActualTag)
	} else if v.MsgError == "" && v.Tag != v.ActualTag {
		if v.Param != "" {
			v.MsgError = fmt.Sprintf("Error: Validate field:`%s` %s is %s", v.Field, v.Tag, v.ActualTag)
		} else {
			v.MsgError = fmt.Sprintf("Error: Validate field:`%s` %s=%s", v.Field, v.Tag, v.Param)
		}
	} else {
		v.MsgError = fmt.Sprintf("Error: Validate field:`%s` %s", v.Field, v.Tag)
	}

}

func (v ValidateStructErrors) Error() string {
	e := ""
	for _, v := range v {
		e += fmt.Sprintf("%s\n", v.Error())
	}
	return e
}
