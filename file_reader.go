package main

import "io/ioutil"

type FileReader int

func (r *FileReader) Read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
