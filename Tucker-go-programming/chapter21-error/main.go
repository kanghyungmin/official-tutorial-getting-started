package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(line)
	return err
}

const filename string = "data.txt"

// func main() {
// 	line, err := ReadFile(filename)
// 	if err != nil {
// 		err = WriteFile(filename, "This is WriteFile function\n")
// 		if err != nil {
// 			fmt.Println("Failed to write file:", err)
// 			return
// 		}
// 		line, err = ReadFile(filename)
// 		if err != nil {
// 			fmt.Println("Failed to read file:", err)
// 			return
// 		}
// 	}
// 	fmt.Println("file contents:", line)
// }

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 합니다. f:%g", f)
	}
	return math.Sqrt(f), nil
}

// func main() {
// 	sqrt, err := Sqrt(-2)
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("Sqrt: %v\n", sqrt)
// }

type error interface {
	Error() string
}

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAcdcount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}
	return nil
}

// func main() {
// 	err := RegisterAcdcount("myID", "myPw")
// 	if err != nil {
// 		if errInfo, ok := err.(PasswordError); ok {
// 			fmt.Printf("%v Len:%v RequireLen:%v\n", errInfo, errInfo.Len, errInfo.RequireLen)
// 		}
// 	} else {
// 		fmt.Println("회원 가입됐습니다.")
// 	}
// }

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt(), pos : %d, err:%w", pos, err)
	}

	pos += n + 1
	b, _, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt(), pos : %d, err:%w", pos, err)
	}
	return a * b, nil
}

func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("failed to scan")
	}

	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert word to int, word:%s, err:%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		fmt.Printf("%v\n", numError)
		if errors.As(err, &numError) {
			fmt.Println("NumberError:", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
	readEq("123")
}

// func f() {
// 	fmt.Println("f() 함수 시작")
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("panic 복구 -", r)
// 		}
// 	}()
// 	g()
// 	fmt.Println("f() 함수 끝")
// }

// func g() {
// 	fmt.Printf("9 / 3 = %d\n", h(9, 3))
// 	fmt.Printf("9 / 0 = %d\n", h(9, 0))
// }

// func h(a, b int) int {
// 	if b == 0 {
// 		panic("제수는 0일 수 없습니다.")
// 	}
// 	return a / b
// }

// func main() {
// 	f()
// 	fmt.Println("프로그램이 계속 실행됨")
// }
