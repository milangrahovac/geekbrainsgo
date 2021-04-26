package confreader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type File struct {
	Path string
}

type FileData struct {
	Name     string   `json:"name" yaml:"name"`
	Course   string   `json:"course" yaml:"course"`
	Keywords []string `json:"keywords" yaml:"keywords"`
	Website  string   `json:"website" yaml:"website"`
}

func (f *File) IsFile() bool {
	if _, err := os.Stat(f.Path); err != nil {
		return false
	}
	return true
}

func (f *File) GetExtension() string {
	return filepath.Ext(f.Path)
}

func (f *File) OpenFile() *FileData {
	file, err := os.Open(f.Path)
	if err != nil {
		log.Printf("Canont open file: %v", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Cannot close file: %s.", err)
		}
	}()

	course := FileData{}
	ext := f.GetExtension()

	if ext == ".json" {
		err = json.NewDecoder(file).Decode(&course)

		if err != nil {
			log.Printf("Cannot decode json file: %v", err)
		}
	} else if ext == ".yaml" {

		yamlFile, err := ioutil.ReadFile("course.yaml")
		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, &course)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
	}

	return &course
}
