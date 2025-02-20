package epub_test

import (
	"os"
	"testing"

	"github.com/jrschumacher/epub"
)

func TestEpub(t *testing.T) {
	bk, err := open(t, "fixtures/test.epub")
	if err != nil {
		t.Fatal(err)
	}
	defer bk.Close()
}

func TestEpubBytes(t *testing.T) {
	bkBytes, err := os.ReadFile("fixtures/test.epub")
	if err != nil {
		t.Fatal(err)
	}
	bk, err := epub.OpenBytes(bkBytes)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("files: %+v", bk.Files())
	t.Logf("book: %+v", bk)
}

func open(t *testing.T, f string) (*epub.Book, error) {
	bk, err := epub.Open(f)
	if err != nil {
		return nil, err
	}
	defer bk.Close()

	t.Logf("files: %+v", bk.Files())
	t.Logf("book: %+v", bk)

	return bk, nil
}
