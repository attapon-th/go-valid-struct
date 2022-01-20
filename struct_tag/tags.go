package struct_tag

import (
	"reflect"
	"strings"
)

type StructTags struct {
	Tag  string
	Tags []TagDetail
}

type TagDetail struct {
	Index     int
	FielName  string
	TagName   string
	TagOption string
}

func New(structSource interface{}, tagName string) *StructTags {
	s := structSource
	rv := reflect.ValueOf(s)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}
	st := rv.Type()
	tags := []TagDetail{}
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		tag := TagDetail{Index: i, FielName: field.Name}
		switch tagName {
		case "gorm":
			tag.TagName, tag.TagOption = ParseTagGORM(field)
		default:
			tag.TagName, tag.TagOption = ParseTagDefault(field, tagName)
		}
		tags = append(tags, tag)
	}
	return &StructTags{Tag: tagName, Tags: tags}
}

func ParseTagDefault(field reflect.StructField, tag_name string) (string, string) {
	if tag, ok := field.Tag.Lookup(tag_name); ok {
		if tag == "-" {
			return "-", ""
		}
		// fmt.Println(strings.SplitN(tag, ",", 2))
		if idx := strings.SplitN(tag, ",", 2); len(idx) > 1 {

			return idx[0], idx[1]
		} else {
			return idx[0], ""
		}
	}
	return field.Name, ""
}

func ParseTagGORM(field reflect.StructField) (string, string) {
	sep := ";"
	if tag, ok := field.Tag.Lookup("gorm"); ok {
		if tag == "-" {
			return "-", ""
		}
		names := strings.Split(tag, sep)
		for i := 0; i < len(names); i++ {
			j := i
			if len(names[j]) > 0 {
				for {
					if names[j][len(names[j])-1] == '\\' {
						i++
						names[j] = names[j][0:len(names[j])-1] + sep + names[i]
						names[i] = ""
					} else {
						break
					}
				}
			}
			values := strings.Split(names[j], ":")
			k := strings.TrimSpace(strings.ToUpper(values[0]))
			if k == "COLUMN" {
				if len(values) >= 2 {
					return values[1], tag
				}
			}
		}
	}
	return field.Name, ""
}
