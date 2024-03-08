package request

type CreateJournal struct {
	Event            string  `json:"event"`
	TanggalTransaksi string  `validate:"required" json:"tanggal_transaksi"`
	NoRekeningKredit string  `validate:"required" json:"no_rekening_kredit"`
	NoRekeningDebit  string  `validate:"required" json:"no_rekening_debit"`
	NominalKredit    float64 `validate:"required" json:"nominal_kredit"`
	NominalDebit     float64 `validate:"required" json:"nominal_debit"`
}
