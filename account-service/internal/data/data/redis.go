package data

import (
	"encoding/json"
)

type RedisPublish struct {
	Event            string            `json:"event"`
	NoRekeningKredit string            `json:"no_rekening_kredit"`
	NoRekeningDebit  string            `json:"no_rekening_debit"`
	NominalKredit    float64           `json:"nominal_kredit"`
	NominalDebit     float64           `json:"nominal_debit"`
	TanggalTransaksi string            `json:"tanggal_transaksi"`
	TraceContext     map[string]string `json:"trace_context"`
}

type RedisPublishMutasi struct {
	Event            string            `json:"event"`
	NoRekening       string            `json:"no_rekening"`
	JenisTransaksi   string            `json:"jenis_transaksi"`
	Nominal          float64           `json:"nominal"`
	TanggalTransaksi string            `json:"tanggal_transaksi"`
	TraceContext     map[string]string `json:"trace_context"`
}

func (i RedisPublish) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i RedisPublishMutasi) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
