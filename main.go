package main

import (
	"log"
	"yuanzheng/adb"

	//"github.com/astaxie/beego"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

//func test() {
//	log.Println("connect ...")
//	cmd := exec.Command("adb", "version")
//
//	output, err := cmd.Output()
//	if err != nil {
//		fmt.Println("Execute Command failed:" + err.Error())
//		return
//	}
//
//	fmt.Printf("Execute Shell:finished with output:\n%s", string(output))
//}

func main() {
	//adb.Devices()
	//fmt.Println("-------")

	//adb.Connect("127.0.0.1", 62001)
	//adb.Connect("192.168.1.102", 5554)

	//adb.AdbShellInputKeyEvent2("3")
	//adb.Disconnect("127.0.0.1", 7555)
	//adb.ShellScreenCapPullRm()

	//point := adb.GetPoint("20171120161302025.png","20171120161329769.jpeg")
	point := adb.GetPoint("images/screen/screen.png", "images/template/222222.png")
	if point != nil {
		println(point.Y)
		println(point.X)
	}
	adb.AdbShellInputTap(point.X,point.Y)
	println("ttttt")

}
