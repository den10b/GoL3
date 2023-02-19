package main

import (
	"errors"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

func wGet(link string, mirror bool, path string) error {
	response, err := http.Get(link)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create an empty file
	file, err := os.Create(path + "response.html")
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the field
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	path := flag.String("P", "develop/dev09/", "путь сохраненияя")
	mirror := flag.Bool("mirror", false, "загрузить содержимое целого веб-сайта")
	flag.Parse()

	link := os.Args[len(os.Args)-1]
	link = "https://streamdj.app/"
	err := wGet(link, *mirror, *path)
	if err != nil {
		log.Fatal(err)
	}

}
