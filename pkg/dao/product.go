package dao

import (
	DB "test/pkg/db"
	"test/src/constant"
)

func ProductList() [] constant.ItemJson {
	db := DB.ConnectDB()
	defer db.Close()

	rows, _ := db.Query("select pid, cid, cname, pname, price, energy, img, flag, tag, note from tbl_product")
	defer rows.Close()

	var products [] constant.ItemJson
	for rows.Next() {
		var item constant.ItemJson
		err := rows.Scan(&item.Cid, &item.Pid, &item.Cname, &item.Pname, &item.Price, &item.Energy, &item.Flag, &item.Img, &item.Tag, &item.Note)
		products = append(products, item)
		err = rows.Err()
		if err != nil {
			panic(err.Error())
		}
	}
	return products
}