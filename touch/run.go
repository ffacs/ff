package touch

import (
	"github.com/fatih/color"
	"os"
	"os/exec"
)

func Compile(filename string) error {
	name, _ := Split(filename)
	err := Exe("c++", filename, "-o", name)
	if err != nil {
		return err
	}
	return nil
}

func Run(filename string) {
	name, suffix := Split(filename)
	if suffix == "cpp" {
		err := Compile(filename)
		if err != nil {
			color.Red("%v", err)
			return
		}
		color.Green("Start Running " + name)
		err = Exe("./" + name)
		if err != nil {
			color.Red("%v", err)
		}
	} else if suffix == "py" {
		err := Exe("python3", filename)
		if err != nil {
			color.Red("%v", err)
		}
	}
}

func Exe(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

func Open_vsc(filename ...string) {
	Exe("code", filename...)
}
