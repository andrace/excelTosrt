package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

const srtDirPath = `./srtFiles`

func main() {
	CreateDir(srtDirPath)
	readDir("./")
}

func excelToSrt(path string, name string) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	srt := ""
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Println(colCell)
			srt += colCell + "\r\n"
		}
		srt += "\r\n"
		fmt.Println()
	}
	create, err := os.Create(srtDirPath + "/" + strings.Trim(name, ".xlsx") + ".srt")
	if err != nil {
		fmt.Println(err.Error())

	}
	write, err := create.Write([]byte(srt))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(srtDirPath+"/"+name, write)

}

func readDir(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return
	}
	for _, v := range dir {
		vpath := path + "/" + v.Name()
		if v.IsDir() {
			readDir(vpath)
		} else {
			if strings.Contains(v.Name(), ".xlsx") {
				excelToSrt(vpath, v.Name())
			}
		}
	}
}

// HasDir 判断文件夹是否存在
func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

// CreateDir 创建文件夹
func CreateDir(path string) {
	_exist, _err := HasDir(path)
	if _err != nil {
		fmt.Printf("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		fmt.Println("文件夹已存在！")
	} else {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
}
