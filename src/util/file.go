package util

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func CopyFiles(dir_src string, dir_dest string) error {
	log.Println(dir_src, dir_dest)

	err := os.MkdirAll(dir_dest, 0777)
	if err != nil {
		log.Println(err)
	}

	source, err := os.Open(dir_src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Open(dir_dest)
    if err != nil {
        return err
    }
	defer destination.Close()

	files, err := ioutil.ReadDir(dir_src)
	for _, file := range files {
		log.Println(path.Join(dir_src, file.Name(), "index.html"))
		file_src, err := os.Open(path.Join(dir_src, file.Name(), "index.html"))
		if err != nil {
			return err
		}

		os.MkdirAll(path.Join(dir_dest, file.Name()), 0777)
		_, err = os.Create(path.Join(dir_dest, file.Name(), "index.html"))
		if err != nil {
			return err
		}

		file_dest, err := os.Create(path.Join(dir_dest, file.Name(), "index.html"))
		if err != nil {
			return err
		}

		_, err = io.Copy(file_dest, file_src)
		if err != nil {
			return err
		}
	}

	return err
}