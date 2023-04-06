package router

import (
	"Mangastica-server/internal/pkg/service"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/chapterslist/:title", service.GetChapterList)
	r.GET("/mangasList", service.GetMangaList)

	r.Run()
}
