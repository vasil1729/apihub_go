package kitchensink

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math/rand"
	"time"
)

type ImagesService struct{}

func NewImagesService() *ImagesService {
	return &ImagesService{}
}

func (s *ImagesService) GenerateJPEG() ([]byte, error) {
	img := s.generateRandomImage(200, 200)
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 90}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *ImagesService) GeneratePNG() ([]byte, error) {
	img := s.generateRandomImage(200, 200)
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *ImagesService) GenerateSVG() string {
	return `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg width="200" height="200" xmlns="http://www.w3.org/2000/svg">
  <rect width="100%" height="100%" fill="blue" />
  <circle cx="100" cy="100" r="50" fill="green" stroke="yellow" stroke-width="4" />
  <text x="50%" y="50%" dominant-baseline="middle" text-anchor="middle" fill="white" font-size="20">Kitchen Sink</text>
</svg>`
}

func (s *ImagesService) GetWebP() []byte {
	// Minimal 1x1 transparent WebP
	// RIFF + WEBP + VP8L chunk
	return []byte{
		0x52, 0x49, 0x46, 0x46, // RIFF
		0x1a, 0x00, 0x00, 0x00, // Size (26 bytes)
		0x57, 0x45, 0x42, 0x50, // WEBP
		0x56, 0x50, 0x38, 0x4c, // VP8L
		0x0d, 0x00, 0x00, 0x00, // Chunk size (13 bytes)
		0x2f, 0x00, 0x00, 0x00, // Signature
		0x10, 0x07, 0x10, 0x11, // Dimensions (1x1)
		0x11, 0x88, 0x88, 0xfe, // Data
		0x07, 0x00,             // Filler
	}
}

func (s *ImagesService) generateRandomImage(w, h int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	
	// Fill background
	draw.Draw(img, img.Bounds(), &image.Uniform{s.randomColor()}, image.Point{}, draw.Src)
	
	// Draw random rectangles
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		x1 := rand.Intn(w)
		y1 := rand.Intn(h)
		x2 := rand.Intn(w)
		y2 := rand.Intn(h)
		rect := image.Rect(x1, y1, x2, y2)
		draw.Draw(img, rect, &image.Uniform{s.randomColor()}, image.Point{}, draw.Src)
	}
	return img
}

func (s *ImagesService) randomColor() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}
