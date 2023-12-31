package api

import (
	"context"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func Write(c *gin.Context) {
	path := c.Query("path")
	prefix := os.Getenv("PATH_PREFIX")
	prefixPath := prefix + path

	bkt := c.MustGet("bucket").(*storage.BucketHandle)

	obj := bkt.Object(prefixPath)

	ctx := context.Background()
	wr := obj.NewWriter(ctx)
	defer wr.Close()

	file, _, err := c.Request.FormFile("file")

	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(wr, file)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"message": "ok"})
}
