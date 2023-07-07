package main

import (
	"fmt"
	"time"
)

func main() {
	//timeStr := "20230606_121314569"
	//timeStr := "20230606_121314"

	//location, _ := time.Parse("20060102_150405", timeStr)
	location := time.Now()
	fmt.Println(location.Format("20060102"))
	fmt.Println(location.Year())
	fmt.Println(location.Month())
	fmt.Println(location.Day())
	fmt.Println(location.Hour())
	fmt.Println(location.Minute())
	fmt.Println(location.Second())
	//fmt.Println(location.Mi())
	fmt.Println("========================")

	newLocation := location.Add(time.Hour * 24 * -7)
	fmt.Println(newLocation.Format("20060102"))
	fmt.Println("========================")

	location = location.AddDate(0, 0, -7)
	year, month, day := location.Date()
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)

	timeStr := "20230606"
	location, _ = time.Parse("20060102", timeStr)
	fmt.Println(location.Year())
	fmt.Println(location.Month())
	fmt.Println(location.Day())
	//hour, min, sec := location.Clock()
	//fmt.Println(hour)
	//fmt.Println(min)
	//fmt.Println(sec)
}
