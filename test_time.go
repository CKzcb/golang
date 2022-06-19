// package main

// import (
// 	"image/color"
// 	"log"

// 	qrcode "github.com/skip2/go-qrcode"
// )

// func main() {
// 	qr, err := qrcode.New("http://192.168.0.102:8080", qrcode.Medium)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		qr.BackgroundColor = color.RGBA{50, 205, 50, 255}
// 		qr.ForegroundColor = color.White
// 		qr.WriteFile(256, "./golang_qrcode.png")
// 	}
// }
