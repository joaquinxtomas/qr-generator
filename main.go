package main

import (
	"fmt"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {

	qrGenerator("WIFI USER", "WIFI PASSWORD")

}

func qrGenerator(user, password string) error {
	wifiData := fmt.Sprintf("WIFI:T:WPA;S:%s;P:%s;;", user, password)
	arch, err := os.Create("qr.png")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer arch.Close()

	err = qrcode.WriteFile(wifiData, qrcode.Medium, 256, "qr.png")

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("QR REALIZADO CORRECTAMENTE.")

	return nil
}
