package resize

import (
	"bytes"
	"image"

	"github.com/disintegration/imaging"
)

func ResizeImgBuf(imgBuf []byte, width int, height int, format imaging.Format, opt imaging.EncodeOption) (*bytes.Buffer, error) {
	decodedImage, err := imaging.Decode(bytes.NewReader(imgBuf))
	if err != nil {
		return nil, err
	}

	resized := imaging.Resize(decodedImage, width, height, imaging.Lanczos)

	var newbuf bytes.Buffer
	err = imaging.Encode(&newbuf, resized, format, opt)
	if err != nil {
		return nil, err
	}
	return &newbuf, nil
}

func CropAndResizeImgBuf(imgBuf []byte, width int, height int, format imaging.Format, opt imaging.EncodeOption) (*bytes.Buffer, error) {
	decodedImage, err := imaging.Decode(bytes.NewReader(imgBuf))
	if err != nil {
		return nil, err
	}

	var resized image.Image
	if width >= height {
		resized = imaging.Resize(decodedImage, width, 0, imaging.Lanczos)
	} else {
		resized = imaging.Resize(decodedImage, 0, height, imaging.Lanczos)
	}

	filled := imaging.Fill(resized, width, height, imaging.Center, imaging.Lanczos)

	var newbuf bytes.Buffer
	err = imaging.Encode(&newbuf, filled, format, opt)
	if err != nil {
		return nil, err
	}
	return &newbuf, nil
}
