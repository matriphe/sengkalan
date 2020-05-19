package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/matriphe/sengkala"
)

const versi = "0.2"

func main() {
	fmt.Println("Sengkalan versi ", versi)
	fmt.Println("")

	args := os.Args[1:]
	if len(args) == 0 {
		tampilkanBantuan("Tidak disertai tahun")
		return
	}

	tahun := args[0]
	sengkalan := sengkala.FromYear(tahun)

	suryaSengkala := sengkalan.GetSuryaSengkala()
	fmt.Println("📅 Tahun Masehi:", suryaSengkala.GetYear())
	fmt.Println("☀️ Surya Sengkala:", suryaSengkala.GetSengkala())
	fmt.Println("📜 Makna Surya Sengkala:")
	tampilkanArti(suryaSengkala.GetMeaning())

	fmt.Println("")

	candraSengkala := sengkalan.GetCandraSengkala()

	fmt.Println("📅 Tahun Jawa:", candraSengkala.GetYear())
	fmt.Println("🌙 Candra Sengkala:", candraSengkala.GetSengkala())
	fmt.Println("📜 Makna Candra Sengkala:")
	tampilkanArti(candraSengkala.GetMeaning())

	fmt.Println("")
}

func tampilkanKesalahan(pesan string) {
	msg := fmt.Sprintf("❌  %s", pesan)
	numMsg := len(msg)

	fmt.Println(strings.Repeat("=", numMsg))
	fmt.Println(msg)
	fmt.Println(strings.Repeat("-", numMsg))
}

func tampilkanBantuan(pesan string) {
	tampilkanKesalahan(pesan)
	fmt.Println("ℹ️ Penggunaan: sengkalan [tahun]")
	fmt.Println("🤖 Contoh: sengkalan", time.Now().Year())
}

func tampilkanArti(arti map[string]string) {
	for k, ar := range arti {
		fmt.Printf("   > %s: %s\n", k, ar)
	}
}
