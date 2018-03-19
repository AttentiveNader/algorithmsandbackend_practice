package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"reflect"
	"strings"
	"sync"
	//	"time"
)

func GetPath() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the path of the directory : ")
	path, err := reader.ReadString('\n')
	if err != nil {
		logErr(err)
		GetPath()
	}
	return path[:len(path)-1]
}
func Getag() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the tag  : ")
	tag, err := reader.ReadString('\n')
	if err != nil {
		logErr(err)
		Getag()
	}
	return tag[:len(tag)-1]
}
func GeHelper() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the Helper  : ")
	Helper, err := reader.ReadString('\n')
	if err != nil {
		logErr(err)
		GeHelper()
	}
	return Helper[:len(Helper)-1]
}
func GetAttribute() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the Attribute  : ")
	Attribute, err := reader.ReadString('\n')
	if err != nil {
		logErr(err)
		GetAttribute()
	}
	return Attribute[:len(Attribute)-1]
}
func LoopThrFiles(files []os.FileInfo, wg *sync.WaitGroup, tag string, attr string, helper string) {
	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go func(file os.FileInfo, wg *sync.WaitGroup) {
			if file.Mode().IsDir() {
				wg.Done()
				return
			}
			f, err := ioutil.ReadFile(workingPath + "/" + file.Name())
			if err != nil {
				logErr(err)
			}
			defer wg.Done()
			content := string(f)
			editedFile := AddAttrToTag(content, tag, attr, helper)
			err = ioutil.WriteFile(workingPath+file.Name(), []byte(editedFile), file.Mode())
			fmt.Println("1 is finsished")
			logErr(err)
		}(files[i], wg)
	}
	wg.Wait()
}
func AddAttrToTag(file string, tag string, attribute string, helper string) string {
	inTagCondition := false
	lines := strings.Split(file, "\n")
	for i := 0; i < len(lines); i++ {
		line := strings.Split(lines[i], " ")
		for t := 0; t < len(line); t++ {
			if inTagCondition {
				if !strings.Contains(line[t], attribute) {
					fmt.Println("--------------------inner of tag----------------------")
					if strings.Contains(line[t], helper) {
						inTagCondition = false
						line[t] = " " + attribute + " " + line[t]
						fmt.Println(line)
						fmt.Println(line[t])
						fmt.Println("-----------------outside tag-----------")
					} else if strings.Contains(line[t], ">") {
						inTagCondition = false
						fmt.Println("-----------------outside tag-----------")
					}
				} else {
					inTagCondition = false
				}
			} else {
				if strings.Contains(line[t], "<"+tag) {
					inTagCondition = true
				}
			}
		}
		lines[i] = strings.Join(line, " ")
	}
	file = strings.Join(lines, "\n")
	return file
}

var workingPath string

func main() {
	var wg sync.WaitGroup
	workingPath = GetPath()
	tag := Getag()
	helper := GeHelper() //helper is just a sign that should be in the tag text like another attribute and don't have it include this attribute value
	attribute := GetAttribute()
	if string(workingPath[len(workingPath)-1]) != "/" {
		workingPath += "/"
	}
	files, err := ioutil.ReadDir(workingPath)
	hPanic(err)
	LoopThrFiles(files, &wg, tag, attribute, helper)
}

func hPanic(err error) {
	if err != nil {
		panic(err)
	}
}
func logErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
