package dto

type RequestGarden struct {
	NamaNode    string  `json:"nama_node" validate:"required"`
	Kelembapan  float32 `json:"kelembapan" validate:"required"`
	TanggalNode string  `json:"tanggal_node" validate:"required,datetime=2006-01-02 15:04:05"`
}
