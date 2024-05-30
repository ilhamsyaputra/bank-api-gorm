package response

import "time"

type DaftarResponse struct {
	NoNasabah  string `json:"no_nasabah"`
	NoRekening string `json:"no_rekening"`
}

type TabungResponse struct {
	Saldo float64 `json:"saldo"`
}

type TarikResponse struct {
	Saldo float64 `json:"saldo"`
}

type TransferResponse struct {
	Saldo float64 `json:"saldo"`
}

type GetSaldo struct {
	Saldo float64 `json:"saldo"`
}

type GetMutasi struct {
	Nominal        float64   `json:"nominal"`
	TipeTransaksi  string    `json:"kode_transaksi"`
	WaktuTransaksi time.Time `json:"waktu"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
