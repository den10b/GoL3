package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strings"
)

var visited map[string]bool

func wGet(mainLink string, addLink string, mirror bool, path string) error {
	response, err := http.Get(mainLink + addLink)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Посетил: %s%s\n", mainLink, addLink)
	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		if visited[href] {
			return
		}
		if (href != "/") && (len(href) > 0) {
			if href[0] == '/' {
				visited[href] = true

				err := wGet(mainLink, href[1:], mirror, path)
				if err != nil {
					return
				}

			}

		}
	})
	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create an empty file

	//mainTrimmer := strings.Trim(strings.Replace(mainLink, "//", "/", -1), "/")
	//mainSplitter := strings.Split(mainTrimmer, "/")

	trimmer := strings.Trim(strings.Replace(addLink, "//", "/", -1), "/")
	splitter := strings.Split(trimmer, "/")
	var file *os.File
	if len(splitter) > 1 {
		joiner := strings.Join(splitter[:len(splitter)-1], "/")
		err = os.MkdirAll(path+joiner, 0777)
		if err != nil {
			panic(err)
		}
		file, err = os.Create(path + strings.Join(splitter[:len(splitter)-1], "/") + ".html")
		if err != nil {
			return err
		}
	} else {
		file, err = os.Create(path + splitter[0] + ".html")
		if err != nil {
			return err
		}
	}
	defer file.Close()
	html, err := doc.Html()
	if err != nil {
		return err
	}

	//Write the bytes to the field
	_, err = file.WriteString(strings.Replace(html, mainLink, path, -1))
	if err != nil {
		return err
	}
	//_, err = io.Copy(file, )
	//if err != nil {
	//	return err
	//}

	return nil
}

func main() {
	path := flag.String("P", "develop/dev09/PARSED/", "путь сохраненияя")
	mirror := flag.Bool("mirror", false, "загрузить содержимое целого веб-сайта")
	flag.Parse()
	visited = make(map[string]bool)
	link := os.Args[len(os.Args)-1]
	link = "https://express-center-msk.ru/"
	err := wGet(link, "", *mirror, *path)
	if err != nil {
		log.Fatal(err)
	}

}
