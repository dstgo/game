package repository

import "github.com/dstgo/kratosx"

type File interface {
	// GetFileURL 获取指定文件的url
	GetFileURL(ctx kratosx.Context, sha string) string
}
