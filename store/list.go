package store

import (
	"context"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func List(bkt *storage.BucketHandle, dir string) DirectoryListing {
	var directory string
	prefix := os.Getenv("PATH_PREFIX")
	if strings.HasSuffix(dir, "/") || dir == "" {
		directory = prefix + dir
	} else {
		directory = prefix + dir + "/"
	}

	query := &storage.Query{Prefix: directory, Delimiter: "/", IncludeTrailingDelimiter: true}

	ctx := context.Background()

	it := bkt.Objects(ctx, query)

	folders := []Dir{}
	files := []File{}

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		name_ln := len(attrs.Name)
		pfx_ln := len(prefix)
		if name_ln == 0 || name_ln-pfx_ln == 0 || attrs.Name == dir {
			continue
		}
		full_name := attrs.Name[pfx_ln:]
		if strings.HasSuffix(full_name, "/") {
			name := full_name[:len(full_name)-1]
			slashIdx := strings.LastIndex(name, "/")
			if slashIdx >= 0 {
				name = strings.ReplaceAll(name[slashIdx:], "/", "")
			}
			if full_name != dir {

				folders = append(folders, Dir{
					Name:     name,
					FullName: full_name,
				})
			}
		} else {
			slashIdx := strings.LastIndex(full_name, "/")
			var name string
			if slashIdx >= 0 {
				name = strings.ReplaceAll(full_name[strings.LastIndex(full_name, "/"):], "/", "")
			} else {
				name = full_name
			}

			files = append(files, File{
				Name:        name,
				FullName:    full_name,
				Size:        attrs.Size,
				Created:     attrs.Created,
				ContentType: attrs.ContentType,
			})
		}

	}
	return DirectoryListing{
		Folders: &folders,
		Files:   &files,
	}
}
