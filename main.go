package main

import (
	"log"
	"time"
	"yuanzheng/task"

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
	//img := adb.ShellScreenCapPullRm("emulator-5558")
	//point,err :=adb.GetPoint(img, "images/template/richang/tiaoZhan.png")
	//if err != nil {
	//	println(point.Y)
	//	println(point.X)
	//}
	//adb.TestMatchTemplate("images/template/33333.png", "images/template/2222.png")
	////adb.AdbShellInputTap(point.X,point.Y)
	//adb.ShellScreenCapPullRm("emulator-5558")
	//go task.RiChangMain("emulator-5556")
	go task.RiChangMain("emulator-5558")
	//task.RiChangMain("emulator-5560")
	//go task.RiChangMain("emulator-5562")
	//go task.RiChangMain("emulator-5564")
	//go task.RiChangMain("emulator-5566")
	//go task.RiChangMain("emulator-5568")

	//go task.RiChangMain("emulator-5556")

	//task.Unlock("emulator-5558")
	//task.Back("emulator-5558")
	//task.SureClick("emulator-5558")
	time.Sleep(10000 * time.Second)
}
