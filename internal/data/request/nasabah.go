package request

type DaftarRequest struct {
	Nama       string `json:"nama"`
	Nik        string `json:"nik"`
	NoHp       string `json:"no_hp"`
	Pin        string `json:"pin"`
	KodeCabang string `json:"kode_cabang"`
}

type TabungRequest struct {
	NoRekening string  `json:"no_rekening"`
	Nominal    float64 `json:"nominal"`
}

type TarikRequest struct {
	NoRekening string  `json:"no_rekening"`
	Nominal    float64 `json:"nominal"`
}

type TransaksiRequest struct {
	NoRekeningAsal   string  `json:"no_rekening_asal"`
	NoRekeningTujuan string  `json:"no_rekening_tujuan"`
	Nominal          float64 `json:"nominal"`
}

type LoginRequest struct {
	NoNasabah string `json:"no_nasabah" validate:"required"`
	Pin       string `json:"pin" validate:"required"`
}
