package store

import (
	"context"
	"io"
	"main/lib/system"
	"mime"
	"os"

	pth "path"

	"cloud.google.com/go/storage"
)

func Read(path string, bkt *storage.BucketHandle, out io.Writer) (*File, func(), func(), error) {
	prefix := os.Getenv("PATH_PREFIX")
	prefixPath := prefix + path

	name := system.FilenameFromPath(path)
	cachePath := "/tmp/" + system.ShortHash(prefixPath) + name // set temp path

	if system.LocalExists(cachePath) {
		stats, _ := os.Stat(cachePath)
		f, err := os.Open(cachePath)
		if err != nil {
			return nil, nil, nil, err
		}
		file := File{
			Name:        name,
			FullName:    path,
			Size:        stats.Size(),
			Created:     stats.ModTime(),
			ContentType: mime.TypeByExtension(pth.Ext(cachePath)),
		}
		return &file, func() {
			if _, err := io.Copy(out, f); err != nil {
			}
		}, func() { f.Close() }, nil
	} else {

		obj := bkt.Object(prefixPath)

		ctx := context.Background()

		attrs, err := obj.Attrs(ctx)
		if err != nil {
			return nil, nil, nil, err
		}

		r, err := obj.NewReader(ctx)

		if err != nil {
			return nil, nil, nil, err
		}

		var fi *os.File
		var wr io.Writer
		if attrs.Size < 10000000 /* 10 mb */ {
			fi, err = os.Create(cachePath)
			if err != nil {
				return nil, nil, nil, err
			}

			wr = io.MultiWriter(out, fi)
		} else {
			wr = out
		}
		file := File{
			Name:        name,
			FullName:    path,
			Size:        attrs.Size,
			Created:     attrs.Created,
			ContentType: attrs.ContentType,
		}

		return &file, func() {
			if _, err := io.Copy(wr, r); err != nil {
			}
		}, func() { fi.Close(); r.Close() }, nil
	}
}
