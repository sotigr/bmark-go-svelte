package api

import (
	"context"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	path := c.Query("path")
	prefix := os.Getenv("PATH_PREFIX")
	prefixPath := prefix + path

	bkt := c.MustGet("bucket").(*storage.BucketHandle)

	obj := bkt.Object(prefixPath)

	ctx := context.Background()
	obj.Delete(ctx)

	c.JSON(200, gin.H{"message": "ok"})
}
