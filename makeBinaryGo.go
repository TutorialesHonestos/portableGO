package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/cheggaaa/pb.v1"
)

func main() {
	var binFile string
	var goFile string
	flag.StringVar(&binFile, "b", "", "Archivo Binario")
	flag.StringVar(&goFile, "g", "", "Archivo de codigo Go")
	flag.Parse()
	createGoBinary(binFile, goFile)
}

func readNextBytes(file *os.File, number int64) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func getTemplate(binaryFile string) string {
	templateGO, err := os.Open(".\\template\\go.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer templateGO.Close()
	log.Println("Envio a Reemplazar")
	gocodeBytes, err := ioutil.ReadAll(templateGO)
	gocode := string(gocodeBytes)
	binaryFile = filepath.Base(binaryFile)
	structName := strings.Replace(binaryFile, ".", " ", -1)
	structName = strings.Title(structName)
	structName = strings.Replace(structName, " ", "", -1)
	gocode = strings.Replace(gocode, "${BINARY_NAME}", binaryFile, -1)
	gocode = strings.Replace(gocode, "${STRUCT_NAME}", structName, -1)
	return gocode
}

func createGoBinary(binaryFile, goFile string) {
	templateGoFile := getTemplate(binaryFile)
	file, err := os.Open(binaryFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	//pincel := make(map[string][]int64) //map[int][]int
	//lastOne := "0"
	fileStat, err := file.Stat()
	size := fileStat.Size()
	data := readNextBytes(file, size)
	bar := pb.StartNew(100)
	block := size / 100
	var count int64 = 0
	for _, b := range data {
		if count >= block {
			bar.Increment()
			block += (size / 100)
		}
		count++
		sb := strconv.Itoa(int(b))
		/*
			if (lastOne!=sb){
				lastOne =  sb
				pincel[sb] = append(pincel[sb], count)
			}
		*/
		templateGoFile = strings.Replace(templateGoFile, "${BIT}", sb+",${BIT}", -1)
	}
	templateGoFile = strings.Replace(templateGoFile, ",${BIT}", "", -1)
	//lienzo := util.StringifyJSON(pincel)
	//fmt.Println(lienzo)
	defer writeGoFile(goFile, templateGoFile)
}

func writeGoFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
