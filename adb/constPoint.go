package task

type Point struct {
	X, Y int
}
// 日常坐标
var RiChangHome  = Point{
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
	X:835,
	Y:450,
}