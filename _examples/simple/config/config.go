package config

import (
	"io"
	"strings"
)

const ValidConfigType = "ini"

var validConfigMsg = `
[global]
required=กรุณาใส่ข้อมูลใน %field%
email=กรุณาใส่ข้อมูล Email ให้ถูกต้อง
iscolor=กรุณาใส่ค่าสีให้ถูกต้อง (%detail%)

[user]
age.gte=ข้อมูลต้องมากว่า %param%
age.lte=ข้อมูลต้องน้อยกว่า %param%

`

func GetValidConfigMsg() io.Reader {
	return strings.NewReader(validConfigMsg)
}
