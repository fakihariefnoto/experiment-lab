package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type (
	Content struct {
		Length       int64
		AcceptRanges bool
		FileName     string
		URL          string
		FileRanges   []string
	}
)

func panicAtTheDisco(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	c, e := check("https://raw.githubusercontent.com/fakihariefnoto/experiment-lab/master/database/testdata.json")
	panicAtTheDisco(e)
	download(c)
	combine(c)

}

func check(link string) (c Content, e error) {
	fmt.Println("Checking url " + link)
	u, err := url.Parse(link)
	if err != nil {
		return c, err
	}

	c.URL = link

	resp, err := http.Head(link)
	if err != nil {
		return c, err
	}

	for _, rangeType := range resp.Header["Accept-Ranges"] {
		if strings.Contains(rangeType, "bytes") {
			c.AcceptRanges = true
			break
		}
	}

	if len(resp.Header["Content-Length"]) > 0 {
		c.Length, e = strconv.ParseInt(resp.Header["Content-Length"][0], 10, 64)
	}

	c.FileName = strings.Replace(u.Path[1:], "/", "-", -1)

	startRange := int64(0)
	a := int64(c.Length / 10)
	fmt.Println("total length : ", c.Length)
	for x := 0; x < 10; x++ {
		endRange := startRange + a
		if endRange > c.Length {
			endRange = c.Length
		}
		theRange := fmt.Sprintf("bytes=%v-%v", startRange, endRange)
		fmt.Println("Range ", theRange)
		c.FileRanges = append(c.FileRanges, theRange)
		startRange = endRange + 1
	}

	fmt.Printf("File name : %v\n", c.FileName)
	fmt.Printf("Size : %v KB\n", c.Length/1000)

	return

}

func download(c Content) {
	defer fmt.Println("All download completed.")
	fmt.Printf("Downloading %v ...\n", c.FileName)
	var wg sync.WaitGroup
	for x := 0; x < 10; x++ {
		wg.Add(1)
		fmt.Printf("Download Part %v Started..\n", x)
		downloadPart(c, x, &wg)
	}
	wg.Wait()
}

func downloadPart(c Content, index int, wg *sync.WaitGroup) (err error) {
	startTime := time.Now()
	defer fmt.Printf("Download Part %v Completed : %v Second\n", index, time.Since(startTime))
	defer wg.Done()

	client := new(http.Client)
	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return
	}

	req.Header.Set("Range", c.FileRanges[index])
	res, err := client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	tempFilename := fmt.Sprintf("temp_%v", index)
	// err = os.WriteFile(tempFilename, data, 0644)
	// if err != nil {
	// 	return
	// }

	f, err := os.Create(tempFilename)
	if err != nil {
		return
	}

	_, err = f.Write(data)
	if err != nil {
		return
	}
	f.Close()
	return
}

func combine(c Content) error {
	fmt.Println("Merging all files...")
	defer fmt.Printf("File %v is ready\n", c.FileName)
	mergedFile, err := os.OpenFile(c.FileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.New(fmt.Sprintf("failed to open %v file:", c.FileName, err))
	}
	defer mergedFile.Close()

	for x := 0; x < 10; x++ {
		tempFilename := fmt.Sprintf("temp_%v", x)

		fileTemp, err := os.Open(tempFilename)
		if err != nil {
			return errors.New("Failed to open file " + tempFilename)
		}

		_, err = io.Copy(mergedFile, fileTemp)
		if err != nil {
			return errors.New("Failed to append file " + tempFilename)
		}
		fileTemp.Close()
	}
	return nil
}
