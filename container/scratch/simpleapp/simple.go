package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

/*

Simpleapp will be used for creaeting simple container app
file dir = wd + files

*/

var (
	fileDirectory = "files"
	content       = `SAMPAH!!`
	path          string
	ext           = ".txt"

	// max file name 35
	maxChar = 35
)

func main() {
	fmt.Println("Contoh aplikasi yang dijalanin dalam kontainer..")
	fmt.Println("yang akan listen ke port 9000 untuk melayani http request..")
	fmt.Println("[GET][ /read/:namafile ] untuk baca file dari dir.")
	fmt.Println("[GET][ /list ] untuk check list file apa aja")
	fmt.Println("[POST] /write/:namafile ] untuk menulis file pake default content")

	var err error
	path, err = os.Getwd()
	if err != nil {
		panic("Failed to get working directory" + err.Error())
	}

	if _, err := os.Stat(fileDirectory); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(fileDirectory, os.ModePerm)
		if err != nil {
			panic("Failed to create files directory" + err.Error())
		}
	}

	router := httprouter.New()

	router.GET("/read/:filename", read)
	router.POST("/write/:filename", write)
	router.GET("/list", listFile)

	fmt.Println("Listen to :9000")
	err = http.ListenAndServe(":9000", router)
	if err != nil {
		fmt.Println(err)
	}
}

func write(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	filename := p.ByName("filename")

	if len(filename) > maxChar-4 {
		filename = filename[0 : maxChar-4]
	}

	fullPath := fmt.Sprintf("%v/%v/%v%v", path, fileDirectory, filename, ext)

	fullContent := ""
	random := 33
	randomBig, _ := rand.Int(rand.Reader, big.NewInt(500))
	if randomBig != nil {
		random = int(randomBig.Int64())
	}

	fmt.Printf("Writing files %v with random content %v \n", filename+ext, random)
	for i := 0; i < random; i++ {
		fullContent += content
	}

	err := ioutil.WriteFile(fullPath, []byte(fullContent), 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Sorry, There's trouble when writing your file for a moment, try again later, %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Writing file %v success!\n", filename+ext)
}

func read(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	filename := p.ByName("filename")

	fullPath := fmt.Sprintf("%v/%v/%v%v", path, fileDirectory, filename, ext)

	data, err := ioutil.ReadFile(fullPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Sorry, There's trouble when reading your file for a moment, try again later, %v\n", err)
		return
	}

	w.Header().Add("content-type", "text")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(data))
}

func listFile(w http.ResponseWriter, req *http.Request, p httprouter.Params) {

	files, err := ioutil.ReadDir(path + "/" + fileDirectory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Sorry, Can't listing files for a moment, try again later, %v\n", err)
		return
	}

	listFiles := "\nYou can check your file here :\n\n"
	listFiles += "|=============== NAME ================|================= DATE ==================|\n"

	for _, file := range files {
		listFiles += "| "
		for i := 0; i < maxChar; i++ {
			if i == 0 {
				listFiles += file.Name()
				i += len(file.Name())
			}
			listFiles += "."
			if i == maxChar-1 {
				listFiles += fmt.Sprintf(" | %v |\n", file.ModTime())
			}
		}
	}
	listFiles += "|=====================================|=========================================|\n"

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, listFiles)
}
