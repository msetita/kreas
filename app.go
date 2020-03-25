package main

import (
	"log"
	"github.com/otiai10/copy"
	"os"
)

func InitProject(projectType string, projectName string){
	dest, _ := os.Getwd()
	err := copy.Copy(rootProjectPath() + "/resources/" + projectType, dest + "/" + projectName)
	if err != nil {
		log.Fatal(err.Error())
	}
}
