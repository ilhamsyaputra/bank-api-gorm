package enum

var Counter = struct {
	NoNasabah  string
	NoRekening string
}{
	NoNasabah:  "No Nasabah",
	NoRekening: "No Rekening",
}

var TipeTransaksi = struct {
	Kredit string
	Debit  string
}{
	Kredit: "C",
	Debit:  "D",
}
