package service

import (
	"Mangastica-server/internal/pkg/files"
	"Mangastica-server/internal/pkg/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMangaList(g *gin.Context) {
	var mangaList []model.Manga
	var manga model.Manga

	mangas, err := files.GetFileList("")
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	for _, mangaItem := range mangas {
		manga.Title = mangaItem.Name()
		mangaList = append(mangaList, manga)
	}

	g.JSON(http.StatusOK, gin.H{
		"zips": mangaList,
	})
}

func GetChapterList(g *gin.Context) {
	var chapterList []int
	mangaName := g.Params.ByName("title")

	path := fmt.Sprintf("/mangas/%s", mangaName)
	chapters, err := files.GetFileList(path)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	for _, chapter := range chapters {
		numChapter, err := strconv.Atoi(chapter.Name())
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}
		chapterList = append(chapterList, numChapter)
	}

	g.JSON(http.StatusOK, gin.H{
		"title":    mangaName,
		"chapters": chapterList,
	})

}
