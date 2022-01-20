# Go-Validate-Struct

> ใส่สำหรับ Validate Struct เท่านั้น  

## Go Module
```go
import (
    "github.com/attapon-th/go-valid-struct/govalidator"
	"github.com/go-playground/validator/v10"
)
```

## Go Get
```shell
go get github.com/attapon-th/go-valid-struct/govalidator@latest
go get github.com/go-playground/validator/v10

```

## How to Use:
- [simple/main.go](./_examples/simple/main.go)


## Example Config

### GetConfigFileBy: `json` tag
- [exConfigError.yaml](./exConfigError.yaml)


## Easy Use
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/attapon-th/go-valid-struct/govalidator"
	"github.com/go-playground/validator/v10"
)

type User struct {
	FirstName      string     `validate:"required" json:"first_name"`
	LastName       string     `validate:"required" json:"last_name"`
	Age            uint8      `validate:"gte=0,lte=130" json:"age"`
	Email          string     `validate:"required,email" json:"email"`
	FavouriteColor string     `validate:"iscolor" json:"favourite_color"`
}

var  user = &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smithgmail.com",
		FavouriteColor: "#000-",
	}


func main(){
    // Read File Config by filename 
	if err := govalidator.ReadInFile("./exConfigError.yaml"); err != nil {
		panic(err)
	}


    // Validator Struct
    err := govalidator.Struct(user)
    // or
    // err := govalidator.Struct(*user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
		jErr, _ := json.MarshalIndent(err.(govalidator.ValidateStructErrors), "", "   ")
		fmt.Println(string(jErr))

	}

}
```
### `#OUTPUT`
```json
[
   {
      "field": "age",
      "valid": "lte",
      "detail": "lte",
      "param": "130",
      "error": "ข้อมูลต้องน้อยกว่า 130"
   },
   {
      "field": "email",
      "valid": "email",
      "detail": "email",
      "param": "",
      "error": "กรุณาใส่ข้อมูล Email ให้ถูกต้อง"
   },
   {
      "field": "favourite_color",
      "valid": "iscolor",
      "detail": "hexcolor|rgb|rgba|hsl|hsla",
      "param": "",
      "error": "กรุณาใส่ค่าสีให้ถูกต้อง (hexcolor|rgb|rgba|hsl|hsla)"
   },
   {
      "field": "city",
      "valid": "required",
      "detail": "required",
      "param": "",
      "error": "กรุณาใส่ข้อมูลใน city"
   }
]
```