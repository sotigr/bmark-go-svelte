package api

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	path := c.Query("path")
	bkt := c.MustGet("bucket").(*storage.BucketHandle)
	obj := bkt.Object(path)
	ctx := context.Background()
	attrs, err := obj.Attrs(ctx)

	r, err := obj.NewReader(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer r.Close()

	name := attrs.Name[strings.LastIndex(attrs.Name, "/")+1:]
	cachePath := name

	c.Writer.Header().Set("Content-Type", attrs.ContentType)
	c.Writer.Header().Set("Content-Length", strconv.FormatInt(attrs.Size, 10))
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=\""+name+"\"")

	if _, err := os.Stat(cachePath); err == nil {
		f, err := os.Open(cachePath)
		defer f.Close()
		if err != nil {
			fmt.Println(err)
		}
		wr := c.Writer
		if _, err := io.Copy(wr, f); err != nil {
			// TODO: Handle error.

			fmt.Println(err)
		}
	} else {
		fi, err := os.Create(cachePath)
		if err != nil {
			fmt.Println(err)
		}
		defer fi.Close()
		wr := io.MultiWriter(c.Writer, fi)
		if _, err := io.Copy(wr, r); err != nil {
			// TODO: Handle error.
		}
	}

}
