package image

import (
	"bytes"
	"image"
	_ "image"       // required for imaging
	_ "image/color" // required for imaging
	"image/jpeg"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/golibri/text"
	"github.com/polds/imgbase64"
)

// Image stands for an image object
type Image struct {
	URL       string
	object    image.Image
	Base64    string        // jpeg
	Binary    *bytes.Buffer // jpeg
	Height    int
	Width     int
	Thumbnail *Image
}

// New is a default constructor
func New(URL string) *Image {
	return &Image{URL: URL}
}

// Fetch requests the stored URL and  processes the binary data. Chainable.
func (i *Image) Fetch() *Image {
	if !i.URLIsImage() {
		return i
	}
	img, err := i.fetchData()
	if err != nil {
		return i
	}
	i.object = img
	i.fetchDimensions()
	return i
}

// CreateBase64 creates and stores the Base64 string of the underlying image
func (i *Image) CreateBase64() *Image {
	i.CreateBinary()
	i.Base64 = imgbase64.FromBuffer(*i.Binary)
	return i
}

// CreateBinary creates and stores the byte representation of the image
func (i *Image) CreateBinary() *Image {
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, i.object, nil)
	i.Binary = buf
	return i
}

func (i *Image) fetchData() (image.Image, error) {
	res, err := http.Get(i.URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return imaging.Decode(res.Body)
}

func (i *Image) fetchDimensions() *Image {
	size := i.object.Bounds().Size()
	i.Height = size.Y
	i.Width = size.X
	return i
}

// CreateThumbnail shrinks the image to the specified sizes. Chainable.
func (i *Image) CreateThumbnail(width int, height int) *Image {
	thumb := imaging.Thumbnail(i.object, width, height, imaging.CatmullRom)
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, thumb, nil)
	if err != nil {
		return i
	}
	i.Thumbnail = &Image{
		Height: height,
		Width:  width,
		URL:    i.URL,
		Binary: buf,
		Base64: imgbase64.FromBuffer(*buf),
	}
	return i
}

// URLIsImage checks whether the stored URL has png/jpg/jpeg suffix
func (i *Image) URLIsImage() bool {
	u := text.New(i.URL)
	if u.IsEmpty() == true || u.DoesMatchPattern(`\.(png|jpg|jpeg)`) == false {
		return false
	}
	return true
}
