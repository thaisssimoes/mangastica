package files

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetFileList(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(fmt.Sprintf("./Static%s", path))
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func GetFilePart(directory, cleaner string) ([]string, error) {
	var partList []string
	filesList, err := GetFileList(directory)
	if err != nil {
		return nil, err
	}

	for _, fileName := range filesList {
		file, err := regexGetBeforeExpression(fileName.Name(), cleaner)
		if err != nil {
			return nil, err
		}
		partList = append(partList, file)
	}
	return partList, nil
}

func regexGetBeforeExpression(expression, delimiter string) (string, error) {
	cleanExpression := regexp.MustCompile(fmt.Sprintf((`^[^%s]+`), delimiter)).FindString(expression)
	if cleanExpression == "" {
		return "", fmt.Errorf("Uncapable of finding expression")
	}
	return cleanExpression, nil
}

// Para pegar numero de capítulo e volume
func regexGetBetweenExpression(leftDelimiter, rightDelimiter, str string) (string, error) {
	r := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(leftDelimiter) + `(.*?)` + regexp.QuoteMeta(rightDelimiter))
	matches := r.FindAllStringSubmatch(str, -1)
	for _, v := range matches {
		return v[1], nil
	}
	return "", fmt.Errorf("Uncapable of finding expression")
}

func CreateMangaDirectoryByName(path string) {
	fileList, err := GetFilePart("/mangas", "(")
	if err != nil {
		log.Println(err)
		return
	}

	for _, fileName := range fileList {
		strings.Trim(fileName, fileName)
		completedPath := fmt.Sprintf("%s%s", path, fileName)
		createDirectory(completedPath)

	}
}

func createDirectory(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Directory already exists")
	}
}

// TODO
func GetChapterIdFromFile()        {} //Pegar o capitulo no nome do arquivo
func GetListChaptersFromManga()    {} //Pegar a lista de capitulos por mangá
func CreateChaptersDirectoryById() {} //Cria diretórios de capitulos de cada mangá
func GetListPagesChapter()         {} //Pegar a lista de páginas por capítulos

// Descomprimir arquivos .cbz
func DecompressMangaFile(zipFilePath, fileDestinationFolder string) {
	//The first thing is to open the zipped file.
	openedFile, err := zip.OpenReader(zipFilePath)
	if err != nil {
		panic(err)
	}
	// postpone the closing of the file.
	defer openedFile.Close()

	for _, file := range openedFile.File {
		filePath := filepath.Join(fileDestinationFolder, file.Name)
		fmt.Println("unzipping file", filePath)

		destinationFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, file.Mode())
		if err != nil {
			panic(err)
		}
		//Opening the file and copy it's contents
		fileInArchive, err := file.Open()
		if err != nil {
		}
		if _, err := io.Copy(destinationFile, fileInArchive); err != nil {
			panic(err)
		}
		destinationFile.Close()
		fileInArchive.Close()
	}
}
