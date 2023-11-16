package api

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func Move(c *gin.Context) {
	path := c.Query("path")
	to := c.Query("to")
	prefix := os.Getenv("PATH_PREFIX")
	prefixPath := prefix + path
	prefixTo := prefix + to

	bkt := c.MustGet("bucket").(*storage.BucketHandle)

	src := bkt.Object(prefixPath)
	dst := bkt.Object(prefixTo)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	// Optional: set a generation-match precondition to avoid potential race
	// conditions and data corruptions. The request to copy the file is aborted
	// if the object's generation number does not match your precondition.
	// For a dst object that does not yet exist, set the DoesNotExist precondition.
	dst = dst.If(storage.Conditions{DoesNotExist: true})
	// If the destination object already exists in your bucket, set instead a
	// generation-match precondition using its generation number.
	// attrs, err := dst.Attrs(ctx)
	// if err != nil {
	//      return fmt.Errorf("object.Attrs: %w", err)
	// }
	// dst = dst.If(storage.Conditions{GenerationMatch: attrs.Generation})

	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		c.JSON(500, gin.H{"message": "err"})
	}
	if err := src.Delete(ctx); err != nil {
		c.JSON(500, gin.H{"message": "err"})
	}

	c.JSON(200, gin.H{"message": "ok"})
}
