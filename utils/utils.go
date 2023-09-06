package utils

import (
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func GetSpecialRand() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s)
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}
