package usecases

import (
	product_model "bewell-test/pkg/product/models"
	product_repo "bewell-test/pkg/product/repositories"
	"strconv"
	"strings"
)

func GetProduct(req []product_model.ReqInputOrder) ([]any, error) {
	var result []any
	productsMap, err := convertInputOrder(req)
	if err != nil {
		return nil, err
	}
	products, orderNo, err := product_repo.GetProducts(productsMap)
	if err != nil {
		return nil, err
	}
	complementarys, err := product_repo.GetComplementarys(products, orderNo)
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		result = append(result, p)
	}

	for _, c := range complementarys {
		result = append(result, c)
	}
	return result, nil
}

func convertInputOrder(req []product_model.ReqInputOrder) ([]product_model.ResMapPrice, error) {
	products, err := splitInputOrder(req)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func splitInputOrder(req []product_model.ReqInputOrder) ([]product_model.ResMapPrice, error) {
	var products []product_model.ResMapPrice
	for _, x := range req {
		var totalQty int = 0
		product := strings.Split(x.PlatformProductId, "/")
		for _, raw := range product {
			var qty int = 1
			parts := strings.SplitN(raw, "*", 2)
			if len(parts) > 1 {
				if parsedQty, err := strconv.ParseInt(parts[1], 10, 64); err == nil {
					qty = int(parsedQty)
				} else {
					qty = 1
				}
				totalQty = totalQty + (x.Qty * qty)
			} else {
				if len(product) == 1 {
					totalQty = 1
				} else {
					totalQty = totalQty + (x.Qty * qty)
				}
			}
		}
		for _, raw := range product {
			raw := checkWrongPrefix(raw)
			var qty int = 1
			parts := strings.SplitN(raw, "*", 2)
			if len(parts) > 1 {
				raw = parts[0]
				if parsedQty, err := strconv.ParseInt(parts[1], 10, 64); err == nil {
					qty = int(parsedQty)
				} else {
					qty = 1
				}
			}
			newproduct := product_model.ResMapPrice{
				ProductId: raw,
				Qty:       x.Qty * qty,
				UnitPrice: x.UnitPrice,
				TotalQty:  totalQty,
			}
			products = append(products, newproduct)
		}
	}
	return products, nil
}

func checkWrongPrefix(productId string) string { // คงต้องมีแก้ logic ตรงนี้
	if strings.HasPrefix(productId, "x2-3&") {
		return strings.TrimPrefix(productId, "x2-3&")
	}
	if strings.HasPrefix(productId, "%20x") {
		return strings.TrimPrefix(productId, "%20x")
	}
	if strings.HasPrefix(productId, "--") {
		return strings.TrimPrefix(productId, "--")
	}
	return productId
}
