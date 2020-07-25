package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Archiver Archive用
type Archiver interface {
	// DestFmt archiveフォーマット
	DestFmt() string
	// Archive src dir をarchiveし、dist dir に保存する
	Archive(src, dest string) error
}

type zipper struct{}

// ZIP zip-archiver
var ZIP Archiver = (*zipper)(nil)

func (z *zipper) Archive(src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	w := zip.NewWriter(out)
	defer w.Close()
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil // skip
		}
		if err != nil {
			return err
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		f, err := w.Create(path)
		if err != nil {
			return err
		}
		io.Copy(f, in)
		return nil
	})
}

func (z *zipper) DestFmt() string {
	return "%d.zip"
}
