[![Build Status](https://travis-ci.org/golibri/image.svg?branch=master)](https://travis-ci.org/golibri/text)
[![GoDoc](https://godoc.org/github.com/golibri/image?status.svg)](https://godoc.org/github.com/golibri/image)
[![Built with Spacemacs](https://cdn.rawgit.com/syl20bnr/spacemacs/442d025779da2f62fc86c2082703697714db6514/assets/spacemacs-badge.svg)](http://github.com/syl20bnr/spacemacs)

# golibri/image

Incremental image processing from web sources. Reads from png and jpg sources, but uses only jpeg internally and for exports.

# installation

`go get -u github.com/golibri/image`

# usage

1. Initialize with an image URL (`.png`, `.jpg` or `.jpeg`)
2. Download and analyze the actual image (bytes, height, width)
3. Create a thumbnail of arbitrary size off the image (jpg)
4. The Thumbnail also has a Base64 representation

````go
img := image.New("https://golang.org/doc/gopher/frontpage.png"). // init
  Fetch(). // download and analyze
  CreateThumbnail(120, 70) // width, height

thumb := img.Thumbnail // thumb also has type "Image"
````

Check the struct's data fields for any info you might need. You can also create a Base64 off the original image if needed.

# license
LGPLv3. (You can use it in commercial projects as you like, but improvements/bugfixes must flow back to this lib.)
