package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	//"sync"
	"os"
	//"strings"
	//	"time"
)

//var files []os.FileInfo
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
func LoopThrFiles(files []os.FileInfo, wg sync.WaitGroup) {
	for i := 0; i < len(files); i++ {
		go func() {
			wg.Add(1)
			f, err := os.Open(files[i].Name())
			defer func() {
				f.Close()
				wg.Done()
			}()
		}()
	}
}

var wg sync.WaitGroup

func main() {
	path := GetPath()
	files, err := ioutil.ReadDir(path)
	hPanic(err)
	fmt.Println(files)
	LoopThrFiles(files, wg)
	wg.Wait()
}

func hPanic(err error) {
	if err != nil {
		panic(err)
	}
}
func logErr(err error) { log.Fatal(err) }
