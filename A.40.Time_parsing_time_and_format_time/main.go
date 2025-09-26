package main

import (
	"fmt"
	"time"
)

func main() {
	// time.Time
	time1 := time.Now()
	fmt.Printf("time1 %v\n", time1)

	// time.Date(tahun, bulan, tanggal, jam, menit, detik, nanodetik, timezone)
	time2 := time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)

	now := time.Now()
	fmt.Printf("year: %v, month: %v\n", now.Year(), now.Month())

	// parse time to string
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Printf("%v menjadi %v\n", value, date)

	layoutFormat = "02/01/2006 MST"
	value = "02/09/2015 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Printf("%v menjadi %v\n", value, date)

	// pre-defined layout format
	// time.RubyDate -> Mon Jan 02 15:04:05 -0700 2006
	// time.RFC822 -> 02 Jan 06 15:04 MST
	// dan masih banyak lagi
	date2, _ := time.Parse(time.RFC822, "22 sep 25 17:56 WITA")
	fmt.Println(date2)

	date3 := date.Format("Monday 02, January 2006 15:04 MST")
	fmt.Println(date3)

	date4 := date.Format(time.RubyDate)
	fmt.Println(date4)
}
