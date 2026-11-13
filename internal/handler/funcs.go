package handler

import (
	"strconv"
)

func formatRupiah(price int) string {
	teks := strconv.Itoa(price)
	hasil := ""

	for i := len(teks) - 1; i>=0; i--{
		hasil = string(teks[i]) + hasil

		posisi := len(teks) - i
		if posisi % 3 == 0 && i > 0{
			hasil = "." + hasil
		}
	}
	return "Rp" + hasil
}	