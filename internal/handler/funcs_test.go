package handler

import "testing"

func TestFormatRupiah(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{0, "Rp0"},
		{500, "Rp500"},
		{2000, "Rp2.000"},
		{20000, "Rp20.000"},
		{28000, "Rp28.000"},
		{1000000, "Rp1.000.000"},
	}

	for _, tt := range tests {
		actual := formatRupiah(tt.input)
		if actual != tt.expected {
			t.Errorf("formatRupiah(%d) = %s; want %s", tt.input, actual, tt.expected)
		}
	}
}

