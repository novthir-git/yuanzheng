package adb

type Point struct {
	X, Y int
}

// 日常坐标
var RiChangHome = Point{
	X: 750,
	Y: 25,
}

const (
	// 日常按钮修正值常数
	RiChangE = 95
)

var LockStart = Point{
	X: 380,
	Y: 270,
}

var LockEnd = Point{
	X: 650,
	Y: 270,
}

// 寻龙诀执行按钮
var XunLongJueExecBtn = Point{
	X: 835,
	Y: 450,
}

var XunLongJueBackBtn = Point{
	X: 880,
	Y: 40,
}

//日常翻页
var RiChangPageEnd = Point{
	X: 200,
	Y: 100,
}

//日常翻页
var RiChangPageStart = Point{
	X: 200,
	Y: 200,
}
