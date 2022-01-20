package govalidator

import (
	"reflect"

	"github.com/attapon-th/go-valid-struct/struct_tag"
)

func GetJsonTag(fld reflect.StructField) string {
	name, _ := struct_tag.ParseTagDefault(fld, "json")
	if name == "-" {
		return ""
	}
	// fmt.Println(name, opt)
	return name
}

func GetByTagName(fld reflect.StructField, tagName string) string {
	name, _ := struct_tag.ParseTagDefault(fld, tagName)
	if name == "-" {
		return ""
	}
	return name
}

func GetGormTag(fld reflect.StructField) string {
	name, _ := struct_tag.ParseTagGORM(fld)
	if name == "-" {
		return ""
	}
	return name
}
