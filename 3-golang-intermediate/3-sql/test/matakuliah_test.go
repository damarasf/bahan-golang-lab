package test

import (
	"fmt"
	"pert3_damara/model"
	"testing"
)

func TestMatkul(t *testing.T) {

	var dataInsertMatkul = []model.Matkul{
		model.Matkul{
			Kd_mk:      "KA01",
			Matakuliah: "MTK Dasar",
		},
		model.Matkul{
			Kd_mk:      "KA02",
			Matakuliah: "B.Indo Dasar",
		},
		model.Matkul{
			Kd_mk:      "KA3",
			Matakuliah: "IPA Dasar",
		},
	}

	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("Testing insert  mahasiswwa", func(t *testing.T) {
		for _, dataInsert := range dataInsertMatkul {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Testing update  mahasiswwa", func(t *testing.T) {
		var updateData = map[string]interface{}{
			"matakuliah": "Algoritma 1",
		}

		data := dataInsertMatkul[0]
		if err := data.Update(db, updateData); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing Get matkul", func(t *testing.T) {
		// data := dataInsertMatkul[0]
		_, err := model.GetMatkul(db, "KA3")
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println(result)
	})

	t.Run("Testing GetAll matkul", func(t *testing.T) {
		result, err := model.GetAllMatkul(db)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(result[0])
	})

	t.Run("Testing delete matkul", func(t *testing.T) {
		data := dataInsertMatkul[0]
		if err := data.Delete(db); err != nil {
			t.Fatal(err)
		}
	})

	defer db.Close()
}
