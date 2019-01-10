package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/pkg/errors"
)

// ImageRequest --
type ImageRequest struct {
	Images []Image
}

// Image --
type Image struct {
	UUID    string
	Encoded string
}

func getFiles(dirPath string, filesLimit int) (map[string][]byte, error) {
	infos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, errors.Wrap(err, "faild to read directory")
	}

	files := make(map[string][]byte, len(infos))

	for _, info := range infos {
		path := filepath.Join(dirPath, info.Name())
		// 一気に読むのは良くないけど、今日だけ...
		d, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, errors.Wrap(err, "faild to read file")
		}
		files[path] = d
	}
	return files, nil
}

func main() {
	counts := []int{4, 10, 20}
	for _, count := range counts {
		m, err := getFiles("./images", count)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m)
	}
}
