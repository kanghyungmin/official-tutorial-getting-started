package main

import "fmt"

const M = 3

func main() {
	m := make(map[string]string)
	m["이화랑"] = "서울시 광진구"
	m["강형민"] = "경기도 성남시"
	m["백두산"] = "서울시 광진구"
	m["이보원"] = "경기도 부천시"

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "강형민")
	delete(m, "강형민1")

	fmt.Println(m["강형민"])
	fmt.Println(m["강형민1"])
}
