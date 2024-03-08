package data

import "encoding/json"

type RedisPublish struct {
	Event            string  `json:"event"`
	NoRekeningKredit string  `json:"no_rekening_kredit"`
	NoRekeningDebit  string  `json:"no_rekening_debit"`
	NominalKredit    float64 `json:"nominal_kredit"`
	NominalDebit     float64 `json:"nominal_debit"`
	TanggalTransaksi string  `json:"tanggal_transaksi"`
}

func (i RedisPublish) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
