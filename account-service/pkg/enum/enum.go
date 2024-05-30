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

var Status = struct {
	Error   string
	Success string
}{
	Error:   "error",
	Success: "success",
}

var Remark = struct {
	InternalServerError string
}{
	InternalServerError: "terjadi kesalahan pada sistem, harap hubungi technical support",
}
