package main

import (
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {

	archivo, err := os.Create("qr.png")

	if err != nil {
		return
	}

	defer archivo.Close()

	err = qrcode.WriteFile("www.google.com", qrcode.Medium, 256, "qr.png")

	if err != nil {
		return
	}

	fmt.Println("QR HECHO")

}
