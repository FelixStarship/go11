package main

import (
	_ "net/http/pprof"
	. "time"
)

const π = 3.1416
const Pi = π // 等价于：const Pi = 3.1416
const x = float64(3.14)

const (
	No         = !Yes
	Yes        = true
	MaxDegrees = 360
	Unit       = "弧度"
)

func main() {

	var (
		lang, website string = "2008", "example.org"
		releaseAt     int    = 20
	)
	{

		println(lang, website, releaseAt)

		println(Pi)
		//短变量声明
		lang, year := "go language", 2007
		year, createBy := 2009, "Google Research"
		println(lang, "由", year)
		println(year, createBy)
	}

	println("time:",Now().Second())
}
