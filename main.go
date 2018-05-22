package main

import (
    "image"
    "image/draw"
    "image/png"
    "os"
    "log"
    "github.com/nfnt/resize"
)

func main() {

    left := 200
    top := 200
    width := 200
    height := 200

    rectangle := "rectangle.png"


    srcImage, _ := os.Open("b.png")
	defer srcImage.Close()

    imageData, _, _ := image.Decode(srcImage)

    rectImage := image.NewRGBA(image.Rect(0, 0, width, height))


    imgResized := resize.Resize(500, 0, imageData, resize.NearestNeighbor)


    points := image.Pt(left,top)
    draw.Draw(rectImage, rectImage.Bounds(), imgResized, points, draw.Src)

    file, err := os.Create(rectangle)
    if err != nil {
        log.Fatalf("failed create file: %s", err)
    }
    png.Encode(file, rectImage)
}
