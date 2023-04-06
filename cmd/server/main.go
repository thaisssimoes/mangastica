package main

import "Mangastica-server/pkg/files"

func main() {
	path := "Static"

	files.UnzipManga(path)
	//router.HandleRequest()

	//files.CreateMangaDirectoryByName("./Static/mangas/")

}
