package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Resp struct {
	Something string
}

func ResolveGetBrands(w http.ResponseWriter, r *http.Request) {
	var vv Resp = Resp{"wrong"}
	fmt.Println(r)
	j, _ := json.Marshal(vv)
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(j))
}

func ResolveGetBrand(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	var vv Resp = Resp{fmt.Sprintf("here is your brand %v", args["brand_id"])}
	j, _ := json.Marshal(vv)
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, bytes.NewReader(j))
}
