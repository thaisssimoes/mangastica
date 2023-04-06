package files

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	Manga           string
	DestinationPath string
	ZipsPath        []string
}

//fileDestinationFolder PRECISA já ter o nome do mangá

func UnzipManga(rootPath string) {

	var newPath, mangaName, newMangaName string
	newMangaName = ""
	chapterNumber := ""
	newMangaFolder := ""

	zipPath := filepath.Join(rootPath, "zips")
	listFolders, err := GetFileList(zipPath)
	if err != nil {
	}

	for _, file := range listFolders {
		newPath = filepath.Join(zipPath, file.Name())
		mangaName, _ = regexGetBeforeExpression(file.Name(), "v")
		if err != nil {
		}

		if file.Name() == mangaName {
			mangaName = DigitPrefix(file.Name())
			newMangaName = strings.Trim(mangaName, " ")
			chapterNumber = DigitBetween(file.Name())
		} else if newMangaName == "" || !strings.Contains(mangaName, newMangaName) {
			newMangaName = strings.Trim(mangaName, " ")
		}

		if chapterNumber == "" {
			newMangaFolder = filepath.Join(rootPath, "mangas", newMangaName)
			createDirectory(newMangaFolder)
		} else {
			newMangaFolder = filepath.Join(rootPath, "mangas", newMangaName)
			createDirectory(newMangaFolder)
			newMangaFolder = filepath.Join(rootPath, "mangas", newMangaName, chapterNumber)
			createDirectory(newMangaFolder)
			chapterNumber = ""
		}

		chapterSeparation(newPath, newMangaFolder)

	}

}

func chapterSeparation(zipFilePath, fileDestinationFolder string) {
	//The first thing is to open the zipped file.
	openedFile, err := zip.OpenReader(zipFilePath)
	if err != nil {
		panic(err)
	}
	// postpone the closing of the file.
	defer openedFile.Close()

	for _, file := range openedFile.File {

		chapterId, err := regexGetBetweenExpression(" - c", " (", file.Name)
		if err != nil {
		}

		filePath := filepath.Join(fileDestinationFolder, file.Name)
		directoryPath := filepath.Join(fileDestinationFolder, chapterId)

		createDirectory(directoryPath)
		filePath = filepath.Join(directoryPath, file.Name)

		destinationFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, file.Mode())
		if err != nil {
			log.Println(err)
		}

		//Opening the file and copy it's contents
		fileInArchive, err := file.Open()
		if err != nil {
			log.Println(err)
		}
		if _, err := io.Copy(destinationFile, fileInArchive); err != nil {

		}
		destinationFile.Close()
		fileInArchive.Close()
	}
}
