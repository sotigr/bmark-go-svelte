package api

import (
	"context"
	"main/lib/system"
	"main/store"
	"os"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func Write(c *gin.Context) {
	path := c.Query("path")
	prefix := os.Getenv("PATH_PREFIX")
	prefixPath := prefix + path

	bkt := c.MustGet("bucket").(*storage.BucketHandle)
	file_lock := c.MustGet("file_lock").(*system.SafeList[store.FileLock])

	if file_lock.Exists(func(item *store.FileLock) bool {
		if item.Name == path {
			return true
		} else {
			return false
		}
	}) {
		// TODO
	}

	obj := bkt.Object(prefixPath)

	ctx := context.Background()
	wr := obj.NewWriter(ctx)
	defer wr.Close()

	r, _, _ := c.Request.FormFile("file")

	buffer := []byte{}

	r.Read(buffer)

	wr.Write(buffer)

	c.JSON(200, gin.H{"message": "ok"})

}
