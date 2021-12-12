package util

import (
	"io"
	"log"
	"os"
)

func CopyFiles(dir_src string, dir_dest string) error {
	log.Println(dir_src, dir_dest)

	os.MkdirAll(dir_dest, 0777)

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

	_, err = io.Copy(destination, source)
	return err
}