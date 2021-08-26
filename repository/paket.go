package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"learn/REST/config"
	"learn/REST/models"
	"log"
)

const (
	table = "paket"
)

func GetDataPaket(ctx context.Context) ([]models.Paket, error) {
	var Datas []models.Paket
	db, err := config.GetConnection()

	if err != nil {
		log.Fatal("Error : ", err)
	}

	queryPaket := fmt.Sprintf("SELECT * FROM %v", table)
	rowQuery, err := db.QueryContext(ctx, queryPaket)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var MyPaket models.Paket

		if err = rowQuery.Scan(&MyPaket.ID_PAKET,
			&MyPaket.NAMA_PAKET,
			&MyPaket.HARGA_PAKET,
			&MyPaket.ID_JENIS); err != nil {
			return nil, err
		}
		Datas = append(Datas, MyPaket)
	}
	return Datas, nil
}

func PostDataPaket(ctx context.Context, dts models.Paket) error {
	db, err := config.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	queryPaket := fmt.Sprintf("INSERT INTO %v VALUES (%v,'%v',%v,%v)",
		table,
		dts.ID_PAKET,
		dts.NAMA_PAKET,
		dts.HARGA_PAKET,
		dts.ID_JENIS)

	_, err = db.ExecContext(ctx, queryPaket)

	if err != nil {
		return err
	}
	return nil

}

func UpdateDataPaket(ctx context.Context, dts models.Paket) error {
	db, err := config.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	queryPaket := fmt.Sprintf("UPDATE %v set nama_paket='%v', harga_paket=%v WHERE id_paket=%v ",
		table,
		dts.NAMA_PAKET,
		dts.HARGA_PAKET,
		dts.ID_PAKET)

	_, err = db.ExecContext(ctx, queryPaket)

	if err != nil {
		return err
	}

	return nil
}

func DeleteDataPaket(ctx context.Context, dts models.Paket) error {
	db, err := config.GetConnection()
	if err != nil {
		log.Fatal(err)
	}

	queryPaket := fmt.Sprintf("DELETE FROM %v WHERE id_paket = %v",
		table,
		dts.ID_PAKET)

	val, err := db.ExecContext(ctx, queryPaket)
	if err != nil {
		return err
	}

	check, err := val.RowsAffected()
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	fmt.Print(check)
	if check == 0 {
		return errors.New("id not found")
	}

	return nil
}
