package request

type CreateMutasi struct {
	Event            string            `json:"event"`
	TanggalTransaksi string            `validate:"required" json:"tanggal_transaksi"`
	NoRekening       string            `validate:"required" json:"no_rekening"`
	JenisTransaksi   string            `validate:"required" json:"jenis_transaksi"`
	Nominal          float64           `validate:"required" json:"nominal"`
	TraceContext     map[string]string `json:"trace_context"`
}
