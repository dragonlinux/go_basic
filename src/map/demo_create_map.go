package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func createMapInt() {
	var m = make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	m["3"] = 3

	var n = make(map[string]int)

	for k := range m {
		if strings.EqualFold("2", k) {
			n["4"] = 4
		}
	}

	for k, v := range n {
		m[k] = v
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func createMapString() {
	var tempMap = make(map[string]int)
	tempMap["datatype"] = 2
	tempMap["msgid"] = 17139

	for k, v := range tempMap {
		fmt.Println(k, v)
	}
}

func mapToJson1() {
	s := []map[string]interface{}{}

	m1 := map[string]interface{}{"datatype": 2, "msgid": 55555}
	m2 := map[string]interface{}{"msgid": 17139}

	s = append(s, m1, m2)
	s = append(s, m2)

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	fmt.Println("b:", string(b))
}

func mapToJson2() {
	//s := map[string]interface{}{}

	m1 := map[string]interface{}{"datatype": 2, "msgid": 55555}
	//m2 := map[string]interface{}{"msgid": 17139}

	//s = append(s, m1, m2)
	//s = append(s, m2)

	b, err := json.Marshal(m1)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	fmt.Println(string(b))
}

func demo() {
	foods := map[string]interface{}{
		"datas": struct {
			source string
			price  string
		}{"chicken", "1.75"},
		"datatype": 2,
		"msgid":    "55555",
	}

	fmt.Println(foods)

	b, err := json.Marshal(foods)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	fmt.Println(string(b))
}

func demo2() {
	result := []map[string]interface{}{}

	mp1 := map[string]interface{}{
		"datatype": 2,
		"msgid":    55555,
	}

	mp2 := map[string]interface{}{
		"datas": mp1,
		"four":  4,
	}

	mp3 := make(map[string]interface{})
	for k, v := range mp1 {
		if _, ok := mp1[k]; ok {
			mp3[k] = v
		}
	}

	for k, v := range mp2 {
		if _, ok := mp2[k]; ok {
			mp3[k] = v
		}
	}

	result = append(result, mp1, mp2)
	fmt.Println(result)

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	fmt.Println(string(b))
}

func demoType1() {
	mpXY := map[string]interface{}{
		"mensuo123": 1,
		"yanwu123":  2,
	}

	dragon := "dragon"
	dragonNmae := 555.23

	mpXY[dragon] = dragonNmae

	result := map[string]interface{}{
		"datas":    mpXY,
		"datatype": 1,
		"msgid":    55555,
	}

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	fmt.Println(string(b))
}

func demoType2() {
	mpXY := map[string]interface{}{
		"x": 1,
		"y": 2,
	}

	mpTag := map[string]interface{}{
		"0000-00-00 00:00:00": mpXY,
	}

	mpDatas := map[string]interface{}{
		"canDevice": mpTag,
	}

	mp2 := map[string]interface{}{
		"datas":    mpDatas,
		"datatype": 2,
		"msgid":    55555,
	}

	b, err := json.Marshal(mp2)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}

	fmt.Println(string(b))
}

func main() {
	//createMapInt()
	fmt.Println("=======")
	createMapString()
	fmt.Println("=======")
	//mapToJson1()
	demo()
	fmt.Println("=======")
	//demo2()
	demoType1()
	fmt.Println("=======")
	demoType2()
}
