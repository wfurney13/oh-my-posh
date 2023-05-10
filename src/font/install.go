// Derived from https://github.com/Crosse/font-install
// Copyright 2020 Seth Wright <seth@crosse.org>
package font

import (
	"archive/zip"
	"bytes"
	"io"
)

func InstallZIP(data []byte) (err error) {
	bytesReader := bytes.NewReader(data)

	zipReader, err := zip.NewReader(bytesReader, int64(bytesReader.Len()))
	if err != nil {
		return
	}

	var fonts []*Font

	for _, zf := range zipReader.File {
		rc, err := zf.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		data, err := io.ReadAll(rc)
		if err != nil {
			return err
		}

		fontData, err := newFont(zf.Name, data)
		if err != nil {
			continue
		}

		fonts = append(fonts, fontData)
	}

	for _, font := range fonts {
		if err = install(font); err != nil {
			return err
		}
	}

	return nil
}
