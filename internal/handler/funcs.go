package handler

import (
	"strconv"
)

func FormatRupiah(price int) string {
	p := strconv.Itoa(price)
	return "Rp." + p
}