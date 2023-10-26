package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

/*
=== Утилита wget ===
Реализовать утилиту wget с возможностью скачивать сайты целиком
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

//	example link 		https://www.africau.edu/images/default/sample.pdf
//	example site		https://habr.com/ru/companies/ruvds/articles/346640/

type Cmd struct {
	Site bool
}

func (cmd *Cmd) Usage() {
	fmt.Printf("Usage of %s:\ngo run task.go [-m] [url] \n", filepath.Base(os.Args[0]))
	flag.PrintDefaults()
}

func (cmd *Cmd) Parse() {
	flag.Usage = cmd.Usage
	flag.BoolVar(&cmd.Site, "m", false, "скачать сайт целиком")
	flag.Parse()
}

func main() {
	var cmd = Cmd{}
	cmd.Parse()

	if cmd.Site {
		var mapPath = make(map[string]struct{})

		downloadSite := flag.Args()[0]
		u, err := url.Parse(downloadSite)
		if err != nil {
			panic(err)
		}
		folderName := strings.Replace(u.Host, ".", "", -1)
		filepathToDownload := "dev09/files/" + folderName

		RecursiveDownload(mapPath, filepathToDownload, downloadSite)

	} else {
		var err error
		downloadURL := flag.Args()[0]

		u, err := url.Parse(downloadURL)
		if err != nil {
			panic(err)
		}

		err = downloadFile("dev09/files"+u.Path, downloadURL)
		if err != nil {
			log.Println(err)
		}
	}
}

func downloadFile(filepath string, url string) (err error) {

	var out *os.File

	re := regexp.MustCompile(`\.\w{3}$`)
	match := re.FindStringSubmatch(filepath)
	if len(match) == 0 {
		err = os.MkdirAll(filepath, os.ModePerm)
		out, err = os.Create(filepath + "/index.html")
	} else {
		pat := path.Dir(filepath)
		err = os.MkdirAll(pat, os.ModePerm)
		out, err = os.Create(filepath)
	}

	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func RecursiveDownload(pathMap map[string]struct{}, folderPath, URL string) {
	u, _ := url.Parse(URL)

	err := downloadFile(folderPath+u.Path, URL)
	if err != nil {
		log.Println(err)
	}
	pathMap[u.Path] = struct{}{}

	dataFromFile, err := os.ReadFile(folderPath + u.Path + "/index.html")
	if err != nil {
		return
	}

	// files downloading
	r := regexp.MustCompile(`"(/css[a-z, /, .]+|` + u.Scheme + "://" + u.Host + `.+?)"`)
	sliceOfLinksFromCurrentFile := r.FindAllStringSubmatch(string(dataFromFile), -1)
	fmt.Println("slice-", sliceOfLinksFromCurrentFile)
	for _, link := range sliceOfLinksFromCurrentFile {
		uI, _ := url.Parse(link[1])
		if uI.Scheme == "" {
			link[1] = u.Scheme + "://" + u.Host + link[1]
		}
		if _, ok := pathMap[uI.Path]; !ok {
			RecursiveDownload(pathMap, folderPath, link[1])
		}
	}
	return
}
