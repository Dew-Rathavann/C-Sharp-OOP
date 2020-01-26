package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var info os.FileInfo

func getDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			d := string(drive) + ":/"
			r = append(r, d)
			f.Close()
		}
	}
	return
}
func FindFileFromExtension(extension []string, dir string, files *[]string) {
	fs, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, f := range fs {
			for _, ex := range extension {
				if strings.HasSuffix(f.Name(), ex) {

					if f.ModTime().Year() == time.Now().Year() && f.ModTime().Month() == time.Now().Month() {
						*files = append(*files, f.Name())
					}
				}

			}

			if f.IsDir() {
				path := dir + "/" + f.Name()
				FindFileFromExtension(extension, path, files)
			}
		}
	}
}

func main() {
	start := time.Now()
	drives := getDrives()
	files := []string{}
	for _, d := range drives {
		FindFileFromExtension([]string{".png", ".jpg", ".jpeg", ".gif"}, d, &files)
	}
	fmt.Println(len(files))
	fmt.Println(time.Since(start))
}
