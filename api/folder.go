package api

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func Folder(c *gin.Context) {
	path := c.Query("path")
	prefix := os.Getenv("PATH_PREFIX")
	prefixPath := prefix + path

	if prefixPath[len(prefixPath)-1:] != "/" {
		c.JSON(500, gin.H{"message": "invalid_name"})
		return
	}

	bkt := c.MustGet("bucket").(*storage.BucketHandle)

	obj := bkt.Object(prefixPath)

	ctx := context.Background()
	wr := obj.NewWriter(ctx)
	defer wr.Close()

	_, err := wr.Write([]byte{})

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"message": "ok"})
}
