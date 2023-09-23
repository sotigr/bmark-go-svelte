package api

import (
	"main/store"
	"net/http"
	"strconv"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	path := c.Query("path")
	downloadStr := c.Query("download")
	download := false
	if downloadStr == "true" {
		download = true
	}

	bkt := c.MustGet("bucket").(*storage.BucketHandle)
	file, write, close, err := store.Read(path, bkt, c.Writer)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not_found",
		})
		return
	} else {
		defer close()
	}

	c.Writer.Header().Set("Content-Type", file.ContentType)
	c.Writer.Header().Set("Content-Length", strconv.FormatInt(file.Size, 10))

	if download {
		c.Writer.Header().Set("Content-Disposition", "attachment; filename=\""+file.Name+"\"")
	}

	write()

	if err != nil {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	}

}
