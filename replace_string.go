package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

const (
	expectSuffix = ".md"    // 需替换文件后缀。若为""，替换全部
	dir          = "./test" // 根目录
	sourceStr    = "0"      // 待替换文本
	targetStr    = "1"      // 替换后文本
)

//main 按需更改常量值
func main() {
	ReplaceTextForFilesIn(dir)
}

//ReplaceTextForFilesIn 替换目录下所有文件的文本
func ReplaceTextForFilesIn(dir string) {
	files, err := getAllFiles(dir)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}

	for i, file := range files {
		fmt.Printf("开始处理文件%d: [%s]\n", i, file)
		ReplaceText(file)
	}
}

//ReplaceText 替换某个文件中的字符串
func ReplaceText(filePath string) {
	output, err := readAndReplace(filePath)
	if err != nil {
		if err == io.EOF {
			writeFile(filePath, output)
		}
	} else if len(output) > 0 {
		writeFile(filePath, output)
	}
}

//readAndReplace 读文件并替换
func readAndReplace(filePath string) ([]byte, error) {
	readfile, err := os.OpenFile(filePath, os.O_RDONLY, 0400)
	if err != nil {
		return nil, err
	}
	defer readfile.Close()

	reader := bufio.NewReader(readfile)
	output := make([]byte, 0)

	for {
		line, err := reader.ReadString('\n')
		output = append(output, strings.ReplaceAll(line, sourceStr, targetStr)...)
		if err != nil {
			if err == io.EOF {
				return output, err
			}
		}
	}
}

//writeFile 抹掉文件原来的文本并写入替换后的
func writeFile(filePath string, output []byte) error {
	// 0600为读写文件权限，r:4 w:2 x:1
	outFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	_, err = writer.Write(output)
	if err != nil {
		return err
	}

	writer.Flush()
	return nil
}

//getAllFiles 获取指定目录下的所有文件，包含子目录下的文件（BFS）
func getAllFiles(dirPath string) (files []string, err error) {
	var dirs []string
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	PathSeparator := string(os.PathSeparator)

	// 遍历当前目录
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			// 目录, 递归遍历
			dirs = append(dirs, dirPath+PathSeparator+fileInfo.Name())
		} else {
			// 文件，过滤指定格式
			if ok := len(expectSuffix) == 0 || strings.HasSuffix(fileInfo.Name(), expectSuffix); ok {
				files = append(files, dirPath+PathSeparator+fileInfo.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, subDir := range dirs {
		subDirFiles, _ := getAllFiles(subDir)
		for _, subDirFile := range subDirFiles {
			files = append(files, subDirFile)
		}
	}
	return files, nil
}
