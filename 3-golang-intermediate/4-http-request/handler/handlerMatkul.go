package handler

import (
	"damara_51418626_pert3/model" //sesuaikan dengan nama folder (case sensitive)
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HandlerMatkulGet(w http.ResponseWriter, r *http.Request) {
	var data interface{}
	var err error
	kd_matkul := r.URL.Query()["kd_mk"]
	if len(kd_matkul) != 0 {
		data, err = model.GetMatkul(db, kd_matkul[0])
	} else {
		data, err = model.GetAllMatkul(db)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func HandlerMatkulPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data model.Matkul
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = data.Insert(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(jsonData)
}

func HandlerMatkulDelete(w http.ResponseWriter, r *http.Request) {
	kd_matkul := r.URL.Query()["kd_mk"]
	if len(kd_matkul) != 0 {
		data := model.Matkul{Kd_mk: kd_matkul[0]}
		if err := data.Delete(db); err != nil {
			http.Error(w, "ID tidak ditemukan", http.StatusBadRequest)
			return
		}
		w.Write([]byte("Data telah dihapus"))
	} else {
		http.Error(w, "ID tidak ditemukan", http.StatusInternalServerError)
	}
}

func HandlerMatkulPut(w http.ResponseWriter, r *http.Request) {
	kd_matkul := r.URL.Query()["kd_mk"]
	if len(kd_matkul) == 0 {
		http.Error(w, "ID tidak ditemukan", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonmap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := model.Matkul{Kd_mk: kd_matkul[0]}
	err = data.Update(db, jsonmap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := model.GetMatkul(db, kd_matkul[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
