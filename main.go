package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var fileinfo os.FileInfo

func getDrive() (r []string) {
	var drive string
	fmt.Print("Please select your drive to search(C/D): ")
	fmt.Scan(&drive)
	f, err := os.Open(string(drive) + ":\\")
	if err == nil {
		d := string(drive) + ":/"
		r = append(r, d)
		f.Close()
	}
	return r
}
func FindFileFromExtension(extension []string, dir string, files *[]string) {
	fs, err := ioutil.ReadDir(dir)
	if err == nil {
		for _, f := range fs {
			for _, ex := range extension {
				if strings.HasSuffix(f.Name(), ex) {
					*files = append(*files, f.Name())
				}
			}
			if f.IsDir() {
				path := dir + "/" + f.Name()
				FindFileFromExtension(extension, path, files)
			}
		}
	}
}
func CreateNewFile() {
	drives := getDrive()
	files := []string{}
	newfile, err := os.Create("output.txt")
	for _, d := range drives {
		FindFileFromExtension([]string{".jpg"}, d, &files)
	}
	if err != nil {
		return
	}
	defer newfile.Close()
	for i := 0; i < len(files); i++ {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatal(err)
		}
		fileinfo, err = os.Stat(files[i])
		newfile.WriteString(dir)
		newfile.WriteString(string(fileinfo.Size()))
	}
}
func main() {
	drives := getDrive()
	CreateNewFile()
	fmt.Println(drives)
	fmt.Println("New file created successfully.")
}
