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
