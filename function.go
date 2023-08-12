package ya2

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PrintFilesRecursive(path string, filterFn func(string) bool) {
	var walk func(string)

	walk = func(p string) {
		files, err := os.ReadDir(p)
		if err != nil {
			panic(err)
		}
	
		for _, f := range files {
			filename := filepath.Join(p, f.Name())
	
			if filterFn(filename) {
				fmt.Println(filename)
			}
	
			if f.IsDir() {
				walk(filename)
			}
		}
	}

	walk(path)
}

func PrintAllFiles(path string) {
	printAllFiles(path, "")
}

func PrintAllFilesWithFilter(path string, filter string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		panic(err)
	}

	for _, f := range files {
		filename := filepath.Join(path, f.Name())

		if strings.Contains(filename, filter) {
			fmt.Println(filename)
		}

		if f.IsDir() {
			PrintAllFilesWithFilter(filename, filter)
		}
	}
}

func printAllFiles(path string, indent string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		panic(err)
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())
		// печатаем имя элемента
		fmt.Println(indent, filename)
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			printAllFiles(filename, indent+"\t")
		}
	}
}
