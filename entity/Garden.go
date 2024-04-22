package entity

type Garden struct {
	ID          uint
	NamaNode    string  `json:"nama_node"`
	Kelembapan  float32 `json:"kelembapan"`
	TanggalNode string  `json:"tanggal_node"`
}
