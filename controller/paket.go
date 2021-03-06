package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"learn/REST/models"
	"learn/REST/repository"
	"learn/REST/utils"
	"net/http"
	"strconv"
)

func GetPaket(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		datas, err := repository.GetDataPaket(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, datas, http.StatusOK)
		return
	}
	http.Error(w, "Failed", http.StatusNotFound)
	return
}

func PostPaket(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Wrong Content Type", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var datas *models.Paket

		if err := json.NewDecoder(r.Body).Decode(&datas); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := repository.PostDataPaket(ctx, datas); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Success Insert Data",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}
	http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
	return
}

func UpdatePaket(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Wrong Content Type", http.StatusBadRequest)
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var datas *models.Paket

		if err := json.NewDecoder(r.Body).Decode(&datas); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := repository.UpdateDataPaket(ctx, datas); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status ": "Succes Update",
		}

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
	return
}

func DeletePaket(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// var datas models.Paket

		id := r.URL.Query().Get("id_paket")
		if id == "" {
			utils.ResponseJSON(w, "Please insert ID", http.StatusBadRequest)
		}
		var id_paket int
		id_paket, _ = strconv.Atoi(id)

		if err := repository.DeleteDataPaket(ctx, id_paket); err != nil {
			kesalahan := map[string]string{
				"errors": fmt.Sprintf("%v", err),
			}

			utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Data has been deleted",
		}

		utils.ResponseJSON(w, res, http.StatusOK)
		return
	}
	http.Error(w, "Not Allowed", http.StatusMethodNotAllowed)
}
