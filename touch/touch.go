package touch

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
)

func GetDir() string {
	dir, _ := os.Getwd()
	for i := len(dir) - 1; i >= 0; i-- {
		if dir[i] == '/' {
			return dir[i+1:]
		}
	}
	return dir
}

func Touch(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	_, suffix := Split(filename)
	_ = file.Close()
	if suffix == "cpp" {
		Fill(filename, "cpp")
	}
	color.Green("Generate " + filename + " done.")
	Open_vsc(filename)
}

func Dir_exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

var Ch = "0abcdefghijklmnopqrstuvwxyz"

func Generate_dir(cnt int) {
	for i := 1; i <= cnt; i++ {
		path := Ch[i : i+1]
		if !Dir_exists(path) {
			os.Mkdir(path, 0666)
		}
		Touch(path + "/" + path + ".cpp")
	}
}

func Fill(filename, temp_name string) {
	if temp_name == "" {
		temp_name = "cpp"
	}
	temp_name = "/mnt/d/Codes/template/" + temp_name + ".cpp"
	temp, err := os.OpenFile(temp_name, os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("%v", err)
	}
	defer temp.Close()
	data, err := ioutil.ReadAll(temp)
	if err != nil {
		fmt.Printf("%v", err)
	}
	ioutil.WriteFile(filename, data, 0666)
	Open_vsc(filename)
}
