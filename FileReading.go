package gorcontrol

import (
	"io/ioutil"
)

func check(e error) {

	if e != nil {
		panic(e)
	}
}

func Read(filename string) (string, error) {

	data, err := ioutil.ReadFile(filename)
	check(err)
	return string(data), err
}
