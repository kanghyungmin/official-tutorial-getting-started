package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetFileList(path string) ([]string, error) {
	return filepath.Glob(path)
}

func PrintAllFiles(files []string) {
	for _, path := range files {
		fileList, err := GetFileList(path)
		if err != nil {
			fmt.Println("파일을 찾을 수 없습니다. ", path)
			continue
		}

		for _, file := range fileList {
			fmt.Println(file)
		}
	}
}

// func main() {
// 	if len(os.Args) < 3 {
// 		fmt.Println("2개 이상의 실행 인수가 필요합니다 ex) ex25.1 word filePath")
// 		return
// 	}

// 	word := os.Args[1]
// 	files := os.Args[2:]

// 	fmt.Println("찾으려는 단어: ", word)
// 	PrintAllFiles(files)
// }

// func main() {
// 	PrintFile("hamlet.txt")
// }

// func PrintFile(filename string) {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		fmt.Println("파일을 찾을 수 없습니다.", file)
// 		return
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// }

type LineInfo struct {
	lineNo int
	Line   string
}

type FindInfo struct {
	filename string
	lines    []LineInfo
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("2개 이상의 실행 인수가 필요합니다 ex) ex25.1 word filePath")
		return
	}

	word := os.Args[1]
	files := os.Args[2:]

	findInfos := []FindInfo{}
	for _, path := range files {
		findInfos = append(findInfos, FindWordInAllFiles(word, path)...)
	}

	for _, findInfo := range findInfos {
		fmt.Println(findInfo.filename)
		fmt.Println("----------------------------------")
		for _, lineInfo := range findInfo.lines {
			fmt.Println(lineInfo.lineNo, lineInfo.Line)
		}

		fmt.Println("----------------------------------")
		fmt.Println()
	}
}

func FindWordInAllFiles(word, path string) []FindInfo {
	findInfos := []FindInfo{}
	fileList, err := GetFileList(path)

	if err != nil {
		fmt.Println("파일 경로가 잘못되었습니다. err:", err, "path:", path)
		return findInfos
	}

	for _, file := range fileList {
		findInfos = append(findInfos, FindWordInFile(word, file))
	}
	return findInfos
}

func FindWordInFile(word, filename string) FindInfo {
	findInfo := FindInfo{filename: filename, lines: []LineInfo{}}
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다.", file)
		return findInfo
	}

	defer file.Close()

	lineNo := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, word) {
			findInfo.lines = append(findInfo.lines, LineInfo{lineNo, line})
		}
		lineNo++
	}
	return findInfo
}
