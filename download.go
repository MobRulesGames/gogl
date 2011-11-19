package main

import (
	"fmt"
	"http"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	KhronosRegistryBaseURL = "http://www.opengl.org/registry/api"
	OpenGLEnumSpecFile     = "enum.spec"
	OpenGLEnumExtSpecFile  = "enumext.spec"
	OpenGLSpecFile         = "gl.spec"
	OpenGLTypeMapFile      = "gl.tm"
)

func makeURL(base, file string) string {
	return fmt.Sprintf("%s/%s", base, file)
}

func DownloadFile(baseURL, fileName, outDir string) os.Error {
	fullURL := makeURL(baseURL, fileName)
	fmt.Printf("Downloading %s ...\n", fullURL)
	r, err := http.Get(fullURL)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	absPath, err := filepath.Abs(outDir)
	if err != nil {
		return err
	}
	err = os.MkdirAll(absPath, 0666)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(absPath, fileName), data, 0666)
	if err != nil {
		return err
	}
	return nil
}

func DownloadOpenGLSpecs(baseURL, outDir string) {
	DownloadFile(baseURL, OpenGLEnumExtSpecFile, outDir)
	DownloadFile(baseURL, OpenGLSpecFile, outDir)
	DownloadFile(baseURL, OpenGLTypeMapFile, outDir)
}

// TODO: download wgl, glx specs ...
