package treedir

import (
	"fmt"
	"io"
	"sort"
	"io/ioutil"
	"os"
	"path/filepath"
	"log"
)

func DirTree(output io.Writer, currDir string, printFiles bool) error {
    recursionPrintService("", output, currDir, printFiles)
    return nil
}

func recursionPrintService(prependingString string, output io.Writer, currDir string, printFiles bool) {
    fileObj, err := os.Open(currDir)
    defer fileObj.Close()
    if err != nil {
        log.Fatalf("Could not open %s: %s", currDir, err.Error())
    }
    fileName := fileObj.Name()
    files, err := ioutil.ReadDir(fileName)
    if err != nil {
        log.Fatalf("Could not read dir names in %s: %s", currDir, err.Error())
    }
    var filesMap map[string]os.FileInfo = map[string]os.FileInfo{}
    var unSortedFilesNameArr []string = []string{}
    for _, file := range files {
        unSortedFilesNameArr = append(unSortedFilesNameArr, file.Name())
        filesMap[file.Name()] = file
    }
    sort.Strings(unSortedFilesNameArr)
    var sortedFilesArr []os.FileInfo = []os.FileInfo{}
    for _, stringName := range unSortedFilesNameArr {
        sortedFilesArr = append(sortedFilesArr, filesMap[stringName])
    }
    files = sortedFilesArr
    var newFileList []os.FileInfo = []os.FileInfo{}
    var length int
    if !printFiles {
        for _, file := range files {
            if file.IsDir() {
                newFileList = append(newFileList, file)
            }
        }
        files = newFileList
    }
    length = len(files)
    for i, file := range files {
        if file.IsDir() {
            var stringPrepender string
            if length > i+1 {
                fmt.Fprintf(output, prependingString+"├───"+"%s\n", file.Name())
                stringPrepender = prependingString + "│\t"
            } else {
                fmt.Fprintf(output, prependingString+"└───"+"%s\n", file.Name())
                stringPrepender = prependingString + "\t"
            }
            newDir := filepath.Join(currDir, file.Name())
            recursionPrintService(stringPrepender, output, newDir, printFiles)
        } else if printFiles {
            if file.Size() > 0 {
                if length > i+1 {
                    fmt.Fprintf(output, prependingString+"├───%s (%vb)\n", file.Name(), file.Size())
                } else {
                    fmt.Fprintf(output, prependingString+"└───%s (%vb)\n", file.Name(), file.Size())
                }
            } else {
                if length > i+1 {
                    fmt.Fprintf(output, prependingString+"├───%s (empty)\n", file.Name())
                } else {
                    fmt.Fprintf(output, prependingString+"└───%s (empty)\n", file.Name())
                }
            }
        }
    }
}

func DirTree1(output io.Writer, currDir string, printFiles bool) error {
		dirTreeLocal(output, currDir, printFiles)
		return nil
}

func dirTreeLocal(output io.Writer, currDir string, printFiles bool)  error {
    fileObj, err := os.Open(currDir)
    if err != nil {
        log.Fatalf("Could not open %s: %s", currDir, err.Error())
    }
    defer fileObj.Close()
    fileName := fileObj.Name()
    files, err := ioutil.ReadDir(fileName)
    if err != nil {
        log.Fatalf("Could not read dir names in %s: %s", currDir, err.Error())
    }
    for _, file := range files {
        if file.IsDir() {
            fmt.Fprintf(output, "%s\n", file.Name())
            newDir := filepath.Join(currDir, file.Name())
            dirTreeLocal(output, newDir, printFiles)
        } else if printFiles {
            if file.Size() > 0 {
                fmt.Fprintf(output, "%s (%vb)\n", file.Name(), file.Size())
            } else {
                fmt.Fprintf(output, "%s (empty)\n", file.Name())
            }
        }
    }
    return nil
}
