package box_service

import (
	"CLogServer/pkg/setting"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Box struct {
	From     string   `json:"from"`
	Filename string   `json:"filename"`
	Lines    []string `json:"lines"`
}

func (box *Box) ReceiveBox() {
	nameSection := strings.Split(box.Filename, "/")
	size := len(nameSection)
	filename := setting.AppSetting.LogSavePath + box.From + "-" + nameSection[size-1]

	file := openFile(filename)
	defer closeFile(file)

	writer := bufio.NewWriter(file)
	lines := box.Lines
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("writer写入异常", err)
			return
		}
	}
	err := writer.Flush()
	if err != nil {
		fmt.Println("writer刷新异常", err)
		return
	}
}

func openFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	return file
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		fmt.Println("文件关闭失败", err)
	}
}
