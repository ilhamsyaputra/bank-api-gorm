package response

import "time"

type DaftarResponse struct {
	NoRekening string
}

type TabungResponse struct {
	Saldo float64
}

type TarikResponse struct {
	Saldo float64
}

type TransferResponse struct {
	Saldo float64
}

type GetSaldo struct {
	Saldo float64
}

type GetMutasi struct {
	Nominal        float64   `json:"nominal"`
	TipeTransaksi  string    `json:"kode_transaksi"`
	WaktuTransaksi time.Time `json:"waktu"`
}

type LoginResponse struct {
	Pin string
}
