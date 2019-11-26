package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func createFile(name string) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}
	return os.Create(name)
}

func DeCompress(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := dest + hdr.Name
		file, err := createFile(filename)
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}

func tarImage(sandboxPath string, imagePath string, sandboxName string) {
	//strCmd := "mkdir " + sandboxPath + "/test"
	//fmt.Println(strCmd)
	err := os.Mkdir(sandboxPath + sandboxName, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	err = DeCompress(imagePath, sandboxPath + sandboxName)
	if err != nil {
		fmt.Println(err)
	}
}


func main() {
	sandboxPath := "/home/weilinchen/workspace/schoolfinal/sandboxes"
	imagePath := "/home/weilinchen/workspace/schoolfinal/sandbox/image.tar"
	tarImage(sandboxPath, imagePath, "/test")
}