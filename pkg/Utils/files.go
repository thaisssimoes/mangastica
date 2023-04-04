package Utils

import (
	"fmt"
	"log"
	"os"
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

// Pegar o nome do mangá no nome do arquivo
func GetFilePart(directory, cleaner string) ([]string, error) {
	var partList []string
	filesList, err := GetFileList(directory)
	if err != nil {
		return nil, err
	}

	for _, fileName := range filesList {
		file := regexGetBeforeExpression(fileName.Name(), cleaner)
		partList = append(partList, file)
	}
	return partList, nil
}

func regexGetBeforeExpression(expression, cleaner string) string {
	cleanExpression := regexp.MustCompile(fmt.Sprintf((`^[^%s]+`), cleaner)).FindString(expression)
	return cleanExpression
}

// Cria diretório de mangá pelo nome do arquivo
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
func GetChapterIdFromFile()          {} //Pegar o capitulo no nome do arquivo
func GetListChaptersFromManga()      {} //Pegar a lista de capitulos por mangá
func CreateChaptersDirectoryByName() {} //Cria diretórios de capitulos de cada mangá
func GetListPagesChapter()           {} //Pegar a lista de páginas por capítulos
func DecompressMangaFile()           {} //Descomprimir arquivos .cbz
