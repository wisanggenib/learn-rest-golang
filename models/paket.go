package models

type (
	Paket struct {
		ID_PAKET    int    `json:"id_paket"`
		NAMA_PAKET  string `json:"nama_paket"`
		HARGA_PAKET int    `json:"harga_paket"`
		ID_JENIS    int    `json:"id_jenis"`
	}

	UpdatePaket struct {
		NAMA_PAKET  string `json:"nama_paket"`
		HARGA_PAKET int    `json:"harga_paket"`
		ID_PAKET    int    `json:"id_paket"`
	}
)
