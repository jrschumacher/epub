package epub

import (
	"archive/zip"
	"bytes"
)

func openBook(fd *zip.Reader, close func()) (*Book, error) {
	bk := Book{fd: fd, close: close}
	mt, err := bk.readBytes("mimetype")
	if err == nil {
		bk.Mimetype = string(mt)
		err = bk.readXML("META-INF/container.xml", &bk.Container)
	}
	if err == nil {
		err = bk.readXML(bk.Container.Rootfile.Path, &bk.Opf)
	}

	for _, mf := range bk.Opf.Manifest {
		if mf.ID == bk.Opf.Spine.Toc {
			err = bk.readXML(bk.filename(mf.Href), &bk.Ncx)
			break
		}
	}

	return &bk, err
}

//Open open a epub file
func Open(fn string) (*Book, error) {
	fd, err := zip.OpenReader(fn)
	if err != nil {
		return nil, err
	}

	bk, err := openBook(&fd.Reader, func() { fd.Close() })
	if err != nil {
		fd.Close()
		return nil, err
	}
	return bk, nil
}

func OpenBytes(zipBytes []byte) (*Book, error) {
	fd, err := zip.NewReader(bytes.NewReader(zipBytes), int64(len(zipBytes)))
	if err != nil {
		return nil, err
	}

	return openBook(fd, nil)
}
