package productcontroller

import (
	"go-jwt-mux/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"id":           1,
			"nama_product": "kemeja",
			"stok":         1000,
		},
		{
			"id":           2,
			"nama_product": "handphone",
			"stok":         200,
		}, {
			"id":           3,
			"nama_product": "dasi",
			"stok":         10000,
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)
}
