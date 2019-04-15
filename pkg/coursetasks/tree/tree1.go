package main

func dirTree1(output io.Writer, currDir string, printFiles bool)  error {
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
            dirTree1(output, newDir, printFiles)
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