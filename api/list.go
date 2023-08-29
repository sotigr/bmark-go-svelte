package api

import (
	"main/io"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	bkt := c.MustGet("bucket").(*storage.BucketHandle)
	path := c.Query("path")

	listing := io.List(bkt, path)
	c.JSON(http.StatusOK, listing)
}
