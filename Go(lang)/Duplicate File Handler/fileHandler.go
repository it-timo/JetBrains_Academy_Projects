package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type fileInfo struct {
	path    string
	size    int64
	md5Hash string
}

var (
	deleteMap map[int]fileInfo
)

func calculateMD5Hash(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(f)

	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", md5Hash.Sum(nil)), nil
}

func getFiles(root string, extension string) (*[]fileInfo, error) {
	files := make([]fileInfo, 0)
	walkFunc := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if extension == filepath.Ext(info.Name()) || extension == "" {
				if md5Hash, err := calculateMD5Hash(path); err != nil {
					return err
				} else {
					files = append(files, fileInfo{path: path, size: info.Size(), md5Hash: md5Hash})
				}
			}
		}
		return nil
	}

	if err := filepath.Walk(root, walkFunc); err != nil {
		return nil, err
	} else {
		return &files, nil
	}
}

func sortFilesBySize(files *[]fileInfo, ascending bool) {
	sort.Slice(*files, func(i, j int) bool {
		if ascending {
			return (*files)[i].size < (*files)[j].size
		} else {
			return (*files)[i].size > (*files)[j].size
		}
	})
}

func printFileInfo(files *[]fileInfo) {
	var size int64 = -1
	for _, file := range *files {
		if file.size != size {
			size = file.size
			fmt.Printf("\n%v bytes\n", size)
		}
		fmt.Println(file.path)
	}
}

func printDuplicateFileInfo(files *[]fileInfo) {
	fileHashes := make(map[int64]map[string][]string)
	sizeOrder := make([]int64, 0)
	var size int64 = -1
	for _, file := range *files {
		if file.size != size {
			size = file.size
			sizeOrder = append(sizeOrder, size)
		}
		if _, ok := fileHashes[file.size]; ok {
			if _, ok := fileHashes[file.size][file.md5Hash]; ok {
				fileHashes[file.size][file.md5Hash] = append(fileHashes[file.size][file.md5Hash], file.path)
			} else {
				fileHashes[file.size][file.md5Hash] = make([]string, 0)
				fileHashes[file.size][file.md5Hash] = append(fileHashes[file.size][file.md5Hash], file.path)
			}
		} else {
			fileHashes[file.size] = make(map[string][]string)
			fileHashes[file.size][file.md5Hash] = make([]string, 0)
			fileHashes[file.size][file.md5Hash] = append(fileHashes[file.size][file.md5Hash], file.path)
		}
	}

	count := 1
	for _, size := range sizeOrder {
		duplicateFound := false
		for md5Hash := range fileHashes[size] {
			if len(fileHashes[size][md5Hash]) > 1 {
				if !duplicateFound {
					duplicateFound = true
					fmt.Printf("\n%v bytes\n", size)
				}
				fmt.Printf("Hash: %v\n", md5Hash)
				for _, fileName := range fileHashes[size][md5Hash] {
					fmt.Printf("%v. %v\n", count, fileName)
					deleteMap[count] = fileInfo{path: fileName, size: size}
					count++
				}
			}
		}
	}
}

func getUserInput() (string, bool) {
	var extension string
	var ascending bool

	fmt.Println("Enter file format:")
	if _, err := fmt.Scanln(&extension); err != nil {
		extension = ""
	} else {
		extension = "." + extension
	}

	fmt.Println("\nSize sorting options:")
	fmt.Println("1. Descending")
	fmt.Println("2. Ascending")
	for {
		var answer int
		fmt.Println("\nEnter a sorting option:")
		if _, err := fmt.Scanln(&answer); err == nil {
			if answer == 1 {
				ascending = false
				break
			} else if answer == 2 {
				ascending = true
				break
			}
		}
		fmt.Println("\nWrong option")
	}
	return extension, ascending
}

func checkForDuplicates() bool {
	for {
		var answer string
		fmt.Println("\nCheck for duplicates?")
		if _, err := fmt.Scanln(&answer); err == nil {
			answer = strings.ToLower(answer)
			switch answer {
			case "yes", "y":
				return true
			case "no", "n":
				return false
			}
		} else {
			fmt.Println("Please enter yes/no")
		}
	}
}

func checkForDeletes() {
	for {
		var answer string
		fmt.Println("\nDelete files?")
		if _, err := fmt.Scanln(&answer); err == nil {
			answer = strings.ToLower(answer)
			switch answer {
			case "yes", "y":
				entries := getEntriesToDelete()
				deleteEntries(entries)
				return
			case "no", "n":
				return
			}
		} else {
			fmt.Println("Please enter yes/no")
		}
	}
}

func getEntriesToDelete() []int {
	for {
		var wrongInput bool
		var intSlice []int
		fmt.Println("\nEnter file numbers to delete:")

		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if len(strings.TrimSpace(text)) > 0 {

			answerSlice := strings.Split(strings.TrimSpace(text), " ")

			for _, entry := range answerSlice {
				if len(entry) > 0 {
					c, err2 := strconv.Atoi(entry)

					if c > len(deleteMap)+1 || err2 != nil {
						wrongInput = true
						break
					}

					intSlice = append(intSlice, c)
				}
			}
			if wrongInput {
				fmt.Println("Wrong format inside")
				continue
			} else {
				return intSlice
			}
		}

		fmt.Println("Wrong format inside")
	}
}

func deleteEntries(entries []int) {
	var freedSpace int64

	for _, entry := range entries {
		_ = os.Remove(deleteMap[entry].path)
		freedSpace += deleteMap[entry].size
	}

	fmt.Printf("\nTotal freed up space: %d bytes", freedSpace)
}

func main() {
	deleteMap = make(map[int]fileInfo)

	if len(os.Args) == 2 {
		root := os.Args[1]
		extension, ascending := getUserInput()
		if files, err := getFiles(root, extension); err == nil {
			sortFilesBySize(files, ascending)
			printFileInfo(files)
			if checkForDuplicates() {
				printDuplicateFileInfo(files)
			}
			checkForDeletes()
		} else {
			log.Fatalln(err)
		}
	} else {
		fmt.Println("Directory is not specified")
	}
}
