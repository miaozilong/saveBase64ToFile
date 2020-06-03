package main

import (
	"encoding/base64"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
)

func main() {
	f, err := excelize.OpenFile("base64.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	rows := f.GetRows("Sheet1")
	for i, v := range rows {
		if i == 0 {
			continue
		}
		base64code := v[2]
		filename := v[0] + "." + v[1]
		dist, _ := base64.StdEncoding.DecodeString(base64code)
		//写入新文件
		f, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
		f.Write(dist)
		f.Close()
	}
}
