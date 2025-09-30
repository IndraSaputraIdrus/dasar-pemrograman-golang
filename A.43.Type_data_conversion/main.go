package main

import (
	"fmt"
	"strconv"
)

func printResult(value any) {
	fmt.Printf("Nilai: %v, Type: %T\n", value, value)
}

func stringToInt() {
	var str = "123"
	var num, err = strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error when convert string to int:", err.Error())
		return
	}

	printResult(num)
}

func intToString() {
	var num = 123
	var str = strconv.Itoa(num)

	printResult(str)
}

func parseInt() {
	var str = "123"
	var num, err = strconv.ParseInt(str, 10, 24)

	if err != nil {
		fmt.Println("Error when parseInt:", err.Error())
		return
	}

	printResult(num)
}

func formatInt() {
	var num = int64(24)
	var str = strconv.FormatInt(num, 8)

	printResult(str)
}

func parseFloat() {
	var str = "24.12"
	var num, err = strconv.ParseFloat(str, 32)

	if err != nil {
		fmt.Println("Error in parseFloat:", err.Error())
		return
	}

	printResult(num)
}

func formatFloat() {
	// Format Eksponen 	Penjelasan
	// b 	-ddddp±ddd, a, eksponen biner (basis 2)
	// e 	-d.dddde±dd, a, eksponen desimal (basis 10)
	// E 	-d.ddddE±dd, a, eksponen desimal (basis 10)
	// f 	-ddd.dddd, tanpa eksponen
	// g 	Akan menggunakan format eksponen e untuk eksponen besar dan f untuk selainnya
	// G 	Akan menggunakan format eksponen E untuk eksponen besar dan f untuk selainnya

	var num = float64(24.12)
	var str = strconv.FormatFloat(num, 'f', 6, 64)

	printResult(str)
}

func parseBool() {
	var str = "true"
	var bul, err = strconv.ParseBool(str)

	if err != nil {
		fmt.Println("Error in format bool:", err.Error())
		return
	}

	printResult(bul)
}

func formatBool() {
	var bul = true
	var str = strconv.FormatBool(bul)
	printResult(str)
}

func typeAssertionInAny() {
	var data = map[string]any{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	fmt.Println(data["nama"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Printf("%s bertipe string\n", val)
		case int:
			fmt.Printf("%d bertipe int\n", val)
		case float64:
			fmt.Printf("%f bertipe float64\n", val)
		case bool:
			fmt.Printf("%t bertipe boolean\n", val)
		case []string:
			fmt.Printf("%v bertipe []string\n", val)
		default:
			fmt.Println("type variable tidak di ketahui")
		}
	}
}

func main() {
	stringToInt()
	intToString()

	parseInt()
	formatInt()

	parseFloat()
	formatFloat()

	parseBool()
	formatBool()

	// casting
	var a float64 = float64(24)
	printResult(a)

	b := int32(24.00)
	printResult(b)

	var text1 = "halo"
	var c = []byte(text1)
	fmt.Printf("%d %d %d %d \n", c[0], c[1], c[2], c[3])

	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)
	fmt.Printf("%s \n", s)

	var d int64 = int64('h')
	fmt.Println(d) // 104

	typeAssertionInAny()
}
