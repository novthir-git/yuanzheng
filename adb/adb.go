/**
Adb 基本操作集合
*/
package adb

import (
	"fmt"
	"log"
	"os/exec"
)

const (
	appPkgName = ""
	className  = ""
)

// adb链接模拟器
func Connect(ip string, port int) {
	host := fmt.Sprintf("%s:%d", ip, port)

	log.Printf("connect  host [%s]", host)

	output, _ := exec.Command("adb", "connect", host).Output()

	log.Printf("connect result %s", output)

}
func Disconnect(ip string, port int) {
	host := fmt.Sprintf("%s:%d", ip, port)

	log.Printf("disconnect  host [%s]", host)

	output, err := exec.Command("adb", "disconnect", host).Output()
	if err != nil {
		log.Printf("disconnect error %s", err)

	}
	log.Printf("disconnect result %s", output)
}

//截屏并保存到当前目录下。
//由于需在手机和电脑上复制文件，必要时可增加延时或用下面的PathExists()判断文件是否存在，如：
//time.Sleep(time.Duration(2) * time.Second)
func ShellScreenCapPullRm() {
	exec.Command("adb", "shell", "screencap", "-p", "/sdcard/screen.png").Run()
	exec.Command("adb", "pull", "/sdcard/screen.png", "./images/screen").Run()
	exec.Command("adb", "shell", "rm", "/sdcard/screen.png").Run()
}
//func DeviceNames() {
//	_output, _ := exec.Command("adb", "devices").Output()
//	log.Printf("%s", _output)
//
//	var listStr = string(_output)
//	var array = strings.Split(listStr, "\n")
//	for i := 0; i < len(array); i++ {
//		var content = array[i]
//		if content != "" {
//			if i != 0 {
//
//			}
//		}
//	}
//}

func SendCommand(deviceName string, cmd string) {
	output, err := exec.Command("adb", "-s", deviceName, "shell", cmd).Output()
	if err != nil {
		log.Printf("SendCommand error %s", err)
	}
	log.Printf("SendCommand result %s", output)

}

//func AdbShellInputKeyEvent2(s string) {
//	exec.Command("adb", "shell", "input", "keyevent", s).Run()
//}
