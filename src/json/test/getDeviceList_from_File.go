package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"reflect"
)

func getCommands1(body []uint8) {

	{ //所以这么做，是因为原来的字符串最外层是array，去掉就是object了
		body = body[1 : len(body)-1]
	}
	result := gjson.Get(string(body), "commands")

	fmt.Println(reflect.TypeOf(result))
	fmt.Println(result.IsArray())

	count := 0

	if result.IsArray() {
		for _, name := range result.Array() {
			println(name.String())
			count++
		}
	}

	fmt.Println("array count is :", count)
}

func parseJsonArray(jsonStr []uint8) {
	var val []map[string]interface{} // <---- This must be an array to match input
	if err := json.Unmarshal([]byte(jsonStr), &val); err != nil {
		panic(err)
	}
	//fmt.Println(val)
	//fmt.Println(reflect.TypeOf(val))
	for i, content := range val {
		if i == 0 {
			fmt.Println(i, content)
			fmt.Println(reflect.TypeOf(content))
			fmt.Println(content["id"])
			fmt.Println(content["name"])
			fmt.Println(content["commands"])
			fmt.Println(reflect.TypeOf(content["commands"]))
			fmt.Println("==============>")
			v := content["commands"]
			s := reflect.ValueOf(v)

			//fmt.Println(reflect.TypeOf(s))
			for i := 0; i < s.Len(); i++ {
				//fmt.Println(i, s.Index(i), reflect.TypeOf(s.Index(i)).Kind())
				fmt.Println(i, s.Index(i), reflect.TypeOf(s.Index(i)), reflect.TypeOf(s.Index(i)).Kind())
				//fmt.Println(reflect.TypeOf(s.Index(i)))
				//fmt.Println(reflect.ValueOf(s.Index(i)))
				//fmt.Println(reflect.Indirect(s.Index(i)).FieldByName("name"))
				//v = s.Index(i)
				//fmt.Println(v)

				//fmt.Println(s.Index(i).Elem())
				//fmt.Println(s.Index(i).Elem().Len())
				//for i, x := range s.Index(i) {
				//	fmt.Println(i, x)
				//}
				//v = reflect.TypeOf(s.Index(i)).Kind().
				//for _, key := range v.Mapkeys() {
				//	strct := v.MapIndex(key)
				//	fmt.Println(key.Interface(), strct.Interface())
				//}

				switch reflect.TypeOf(s.Index(i)).Kind() {

				case reflect.Struct:
					name := reflect.ValueOf(s.Index(i))
					fmt.Println("+++++++", name.Interface())
					//fmt.Println("+++++++", name.Interface())

					v := name.Interface()
					fmt.Println("xxxxxxxx", v)
					fmt.Println("xxxxxxxx", reflect.TypeOf(v))

					//for s, a := range v {
					//	// a has type *Author
					//	fmt.Printf("%s: author=%c\n", s, a)
					//}
					//for i, key := range name.Interface() {
					//	fmt.Println(i, key)
					//}

				default:
					//error here, unexpected
				}
			}

		}
	}
}

func parseJsonArray1(jsonStr []uint8) (url string) {
	var val []map[string]interface{} // <---- This must be an array to match input
	if err := json.Unmarshal([]byte(jsonStr), &val); err != nil {
		panic(err)
	}

	for _, content := range val {
		DeviceNames := "Modbus_TCP_test_device"
		//DeviceNames := "Random-Integer-Generator01"

		if content["name"] == DeviceNames {
			//fmt.Println(i, content)
			fmt.Println(reflect.TypeOf(content))
			fmt.Println(content["id"])
			fmt.Println(content["name"])
			//fmt.Println(content["commands"])
			fmt.Println(reflect.TypeOf(content["commands"]))
			fmt.Println("==============>")

			johnJSON, err := json.Marshal(content)
			if err != nil {
				fmt.Println("error:", err)
			}
			//fmt.Println("再转换成json string+++++++++>>>", string(johnJSON), err)
			{
				result := gjson.Get(string(johnJSON), "commands")

				fmt.Println(reflect.TypeOf(result))
				fmt.Println(result.IsArray())

				count := 0

				if result.IsArray() {
					for _, name := range result.Array() {
						//println(i, name.String())
						{
							result1 := gjson.Get(string(name.String()), "name")
							fmt.Println("name ..........>", result1)

							result1 = gjson.Get(string(name.String()), "put")
							//fmt.Println("put  ..........>", result1)

							//fmt.Println(reflect.TypeOf(result1))

							result2 := result1
							result1 = gjson.Get(result1.String(), "url")
							fmt.Println("url  ..........>", result1)
							//url := result1.String()
							_ = result1.String()

							result1 = gjson.Get(result2.String(), "parameterNames")

							fmt.Println("parameterNames length.>", len(result1.Array()))
							if len(result1.Array()) == 1 {
								fmt.Println("parameterNames.>", result1)
								//return url
							}

							fmt.Println("")
						}
						count++
					}
				}
			}
		}
	}
	return ""
}

func DoFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is:", getValue)

	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法
	// 1. 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

func OperatingPlatform1() {

	//path := "./device.json"
	path := "./device_multi_array.json"

	uint8Result, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	//log.Println("uint8Result:", uint8Result)
	//getCommands1(uint8Result)

	url := parseJsonArray1(uint8Result)
	fmt.Println("final get", url)
}

func main() {
	OperatingPlatform1()
}
