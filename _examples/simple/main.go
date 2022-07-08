package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/attapon-th/go-valid-struct/govalidator"
	"github.com/go-playground/validator/v10"
)

// User contains user information
type User struct {
	FirstName      string     `json:"first_name,omitempty" validate:"required"`
	LastName       string     `validate:"required" json:"last_name,omitempty"`
	Age            uint8      `validate:"gte=0,lte=130" json:"age,omitempty"`
	Email          string     `validate:"required,email" json:"email,omitempty"`
	FavouriteColor string     `validate:"iscolor" json:"favourite_color,omitempty"`          // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required" json:"addresses,omitempty"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required" json:"street,omitempty"`
	City   string `validate:"required" json:"city,omitempty"`
	Planet string `validate:"required" json:"planet,omitempty"`
	Phone  string `validate:"required" json:"phone,omitempty"`
}

// TOML Configs massage
var ValidConfigMsgTOML = `
[default]
required = "กรุณาใส่ข้อมูลใน ${field}"
email = "กรุณาใส่ข้อมูล อีเมล์ ให้ถูกต้อง"
iscolor = "กรุณาใส่ค่าสีให้ถูกต้อง (${detail})"

[user.age]
gte = "ข้อมูลต้องมากว่า ${param}"
lte = "ข้อมูลต้องน้อยกว่า ${param}"
`

func main() {

	// * Read Config Error Massage with Golang String
	if err := govalidator.ReadConfig(strings.NewReader(ValidConfigMsgTOML), "toml"); err != nil {
		panic(err)
	}

	// * Read Config Error Massage with File config `toml`
	// if err := govalidator.ReadInFile("./config_error_massage.toml", "toml"); err != nil {
	// 	panic(err)
	// }

	validateStruct()
}

func validateStruct() {

	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smithgmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	// valid.RegisterTagNameFunc()("json")
	err := govalidator.Struct(*user)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
		jErr, _ := json.MarshalIndent(err.(govalidator.ValidateStructErrors), "", "   ")
		fmt.Println(string(jErr))
		_ = jErr

	}

}
