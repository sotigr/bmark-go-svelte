package api

import (
	"context"
	"fmt"
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

	fmt.Println(prefixPath)
	obj := bkt.Object(prefixPath)

	ctx := context.Background()
	wr := obj.NewWriter(ctx)
	defer wr.Close()

	file, header, err := c.Request.FormFile("file")
	filename := header.Filename
	fmt.Println(filename)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(wr, file)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{"message": "ok"})
}
