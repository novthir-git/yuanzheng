package main

import (
	"log"
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
	//安装好 go 之后，直接运行 go get -u -d gocv.io/x/gocv 命令获取 gocv 库，进入该库根目录 cd $GOPATH/src/gocv.io/x/gocv, 运行 source ./env.sh，然后就可以使用 go run命令运行里面的示例了，刚开始学可以直接在示例里面修改代码运行。

	////point := adb.GetPoint("20171120161302025.png","20171120161329769.jpeg")
	//point := adb.GetPoint("images/screen/screen.png", "images/template/1580982624032.jpg")
	//if point != nil {
	//	println(point.Y)
	//	println(point.X)
	//}
	//adb.AdbShellInputTap(point.X,point.Y)\
	println("ttttt")

}
