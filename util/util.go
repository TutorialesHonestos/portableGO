package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

)
func IndexOf(element string, data []string) (int) {
   for k, v := range data {
       if element == v {
           return k
       }
   }
   return -1    //not found.
}

//LoadFile :Carga archivos!
func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

//StringifyJSON , convierte cualquier objeto en una representacion STRING en formato JSON
func StringifyJSON(object interface{}) string {
	b, err := json.Marshal(object)
	if err != nil {
		return "Error"
	}
	return string(b)
}

//IsEmpty es un string vacio?
func IsEmpty(data string) bool {
	if len(data) <= 0 {
		return true
	} else {
		return false
	}
}

//ExistsPath : Retorna si existe archivo o directorio
func ExistsPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateFile(path string, content string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)		
	}
	l, err := f.WriteString(content)
	if err != nil {
		log.Println(err)
		f.Close()		
	}
	log.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		log.Println(err)		
	}
}
