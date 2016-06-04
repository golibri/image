package image

import "testing"

func TestThatItWorks(t *testing.T) {
	img := New("https://golang.org/doc/gopher/frontpage.png"). // init
									Fetch().                 // download and analyze
									CreateThumbnail(120, 70) // width, height

	if img.Height == 0 {
		t.Error("Image Fetching error")
	}

	if img.Thumbnail.Height != 70 {
		t.Error("Thumbnail Processing error")
	}
}
