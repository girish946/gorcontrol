package cmdhandler

import (
	"fmt"
	"github.com/girish946/robotgo"
	"os/exec"
	"time"
)

func ShowDesktop() {

	robotgo.KeyToggle("command", "down")
	robotgo.KeyToggle("control", "down")
	robotgo.KeyToggle("d", "down")

	wait(500)
	robotgo.KeyToggle("command", "up")
	robotgo.KeyToggle("control", "up")
	robotgo.KeyToggle("d", "up")

}

func AltTab() {
	robotgo.KeyToggle("alt", "down")
	robotgo.KeyToggle("tab", "down")
	wait(500)
	robotgo.KeyToggle("tab", "up")
	robotgo.KeyToggle("alt", "up")
}

func Up() {
	robotgo.KeyTap("up")
}

func Down() {
	robotgo.KeyTap("down")
}

func Left() {
	robotgo.KeyTap("left")
}

func Right() {
	robotgo.KeyTap("right")
}

func ShowDashBoard() {

	robotgo.KeyTap("command")
}

func ShowFilemanager() {
	_, err := exec.Command("xdg-open", "/media/").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

}

func AltF4() {
	robotgo.KeyToggle("alt", "down")
	robotgo.KeyToggle("f4", "down")
	robotgo.KeyToggle("f4", "up")
	robotgo.KeyToggle("alt", "up")
}

func HoldWindows() {
	robotgo.KeyToggle("alt", "down")
	robotgo.KeyTap("tab")
}

func Enter() {
	robotgo.KeyTap("enter")
}

func Escape() {
	robotgo.KeyTap("escape")
}

func Home() {
	robotgo.KeyTap("home")
}

func End() {
	robotgo.KeyTap("end")
}

func Backspace() {
	robotgo.KeyTap("backspace")
}

func Tab() {
	robotgo.KeyTap("tab")
}

func F5() {
	robotgo.KeyTap("f5")
}

func PageUp() {
	robotgo.KeyTap("pageup")
}

func PageDown() {
	robotgo.KeyTap("pagedown")
}

func Power() {
	_, err := exec.Command("gnome-session-quit --power-off", "/media/").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
}

func wait(ms uint) {
	<-time.After(time.Millisecond * time.Duration(ms))
}
