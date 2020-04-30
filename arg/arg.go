package arg

import (
	"ffacs/ff/touch"
	"github.com/fatih/color"
	"os"
	"strconv"
)

func Call() {
	if len(os.Args) < 2 {
		color.Blue("r\trun")
		color.Blue("c\tcompile")
		color.Blue("o\topen in vsc")
		color.Blue("g\tmake dirs")
		color.Blue("t\tmake files")
		color.Blue("f\tformat file")
		return
	}
	ope := os.Args[1]
	filename := make([]string, 0, 3)
	for i := 2; i < len(os.Args); i++ {
		filename = append(filename, os.Args[i])
	}
	if ope == "r" {
		if len(filename) == 0 {
			touch.Run(touch.GetDir() + ".cpp")
			return
		}
		for _, i := range filename {
			touch.Run(i)
		}
	} else if ope == "c" { //
		for _, i := range filename {
			touch.Compile(i)
		}
	} else if ope == "o" { //Open
		touch.Open_vsc(filename...)
	} else if ope == "g" { //Gendir
		cnt, _ := strconv.Atoi(filename[0])
		touch.Generate_dir(cnt)
	} else if ope == "t" { //Touch
		if len(filename) == 0 {
			touch.Touch(touch.GetDir() + ".cpp")
		} else {
			for _, i := range filename {
				touch.Touch(i)
			}
		}
	} else if ope == "f" { //Format
		if len(filename) < 1 {
			color.Red("No enough parameters")
			return
		}
		if len(filename) == 1 {
			touch.Fill(filename[0], "")
			return
		}
		touch.Fill(filename[0], filename[1])
	} else {
		color.Red("No operation " + ope)
	}
}
