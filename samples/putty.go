package main

import (
	"io/ioutil"
)

//PuttyExe , representacion en byte de Binario .exe
type PuttyExe struct {
	bytes []byte
	name  string
}

//Make , the binary file
func (app *PuttyExe) Make() {
	app.name = "putty.exe"
	app.createFile()
}

//createFile , final step!
func (app *PuttyExe) createFile() {
	e := ioutil.WriteFile(app.name, app.bytes, 0644)
	if e != nil {
		panic(e)
	}
}

func main() {
	content := PuttyExe{}
	content.Make()

}