package main

import (
	"encoding/base64"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup //定义一个同步等待的组

func main() {
	startTime := time.Now()
	f, err := excelize.OpenFile("base64.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	const PICPATH = "pics"
	os.Mkdir(PICPATH, os.ModePerm)
	rows := f.GetRows("Sheet1")
	for i, v := range rows {
		if i == 0 {
			continue
		}
		base64code := v[2]
		filename := v[0] + "." + v[1]
		dist, _ := base64.StdEncoding.DecodeString(base64code)
		wg.Add(1)
		go writePic(PICPATH, filename, dist)
	}
	wg.Wait()
	elapsed := time.Since(startTime)
	fmt.Println("耗时", elapsed)
	fmt.Println("按回车键退出")
	fmt.Scanln()
}

func writePic(PICPATH string, filename string, dist []byte) {
	//写入新文件
	f, _ := os.OpenFile(PICPATH+"/"+filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	f.Write(dist)
	f.Close()
	wg.Done()
}
