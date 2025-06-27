package handlers

import (
	"encoding/json"
	"net/http"

	product_model "bewell-test/pkg/product/models"
	product_usecase "bewell-test/pkg/product/usecases"
)

func Call() {
	http.HandleFunc("/getProduct", getProduct)
}

func getProduct(writer http.ResponseWriter, req *http.Request) {
	var resp product_model.Response
	var body []product_model.ReqInputOrder
	writer.Header().Set("Content-Type", "application/json")
	// Check Method
	if req.Method != "POST" {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		resp.Status = false
		resp.Message = "Method Not Allowed"
		json.NewEncoder(writer).Encode(resp)
		return
	}
	// Check Body Req
	if true {
		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			resp.Status = false
			resp.Message = "Invalid JSON"
			json.NewEncoder(writer).Encode(resp)
			return
		}
	}
	result, err := product_usecase.GetProduct(body)
	if err != nil {
		resp := product_model.Response{
			Status:  false,
			Message: err.Error(),
		}
		json.NewEncoder(writer).Encode(resp)
		return
	}
	resp.Status = true
	resp.Message = "success"
	resp.Data = result
	json.NewEncoder(writer).Encode(resp)
}
