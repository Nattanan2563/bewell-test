package repositories

import (
	product_model "bewell-test/pkg/product/models"
	utils "bewell-test/utils"
)

func GetProducts(input []product_model.ResMapPrice) ([]product_model.ResCleanedOrder, int, error) {
	db := utils.DB

	var products []product_model.Product
	var productIDs []string
	for _, item := range input {
		productIDs = append(productIDs, item.ProductId)
	}
	if err := db.Select("p.*, p.product_id as master_id").
		Table("products p").
		Where("p.product_id IN (?)", productIDs).Find(&products).Error; err != nil {
		return nil, 0, err
	}
	orderNo := 1
	var complementarys []product_model.ResCleanedOrder
	for _, x := range products {
		for _, y := range input {
			if y.ProductId == x.MasterID {
				comcomplementary := product_model.ResCleanedOrder{
					No:         orderNo,
					ProductId:  x.ProductID,
					MaterialId: x.ProductTypeID + "-" + x.ProductTextureID,
					ModelId:    x.ModelID,
					UnitPrice:  y.UnitPrice / float64(y.TotalQty),
					TotalPrice: y.UnitPrice * float64(y.Qty) / float64(y.TotalQty),
					Qty:        y.Qty,
				}
				complementarys = append(complementarys, comcomplementary)
				orderNo++
			}
		}
	}
	return complementarys, orderNo, nil
}

func GetComplementarys(input []product_model.ResCleanedOrder, orderNo int) ([]product_model.ResCleanedComplementarysOrder, error) {
	db := utils.DB

	var products []product_model.Product
	var productIDs []string
	for _, item := range input {
		productIDs = append(productIDs, item.ProductId)
	}
	if err := db.Select("p2.*, p.product_id as master_id").
		Table("products p").
		Joins("LEFT JOIN products_complementaryitem pc ON p.id = pc.products_id").
		Joins("INNER JOIN products p2 ON pc.products_complementary_id = p2.id").
		Where("p.product_id IN (?)", productIDs).Find(&products).Error; err != nil {
		return nil, err
	}

	// var complementarys []product_model.ResCleanedComplementarysOrder
	// for _, x := range products {
	// 	for _, y := range input {
	// 		if y.ProductId == x.MasterID {
	// 			comcomplementary := product_model.ResCleanedComplementarysOrder{
	// 				No:        orderNo,
	// 				ProductId: x.ProductID,
	// 				// MaterialId: x.ProductTypeID + "-" + x.ProductTextureID,
	// 				// ModelId:    x.ModelID,
	// 				UnitPrice:  0,
	// 				TotalPrice: 0,
	// 				Qty:        y.Qty,
	// 			}
	// 			complementarys = append(complementarys, comcomplementary)
	// 			orderNo++
	// 		}
	// 	}
	// }
	var complementaryMap = make(map[string]product_model.ResCleanedComplementarysOrder)
	for _, x := range products {
		for _, y := range input {
			if y.ProductId == x.MasterID {
				key := x.ProductID
				if existing, ok := complementaryMap[key]; ok {
					// ถ้ามีอยู่แล้ว ให้บวก Qty
					existing.Qty += y.Qty
					complementaryMap[key] = existing
				} else {
					complementaryMap[key] = product_model.ResCleanedComplementarysOrder{
						No:         orderNo,
						ProductId:  x.ProductID,
						UnitPrice:  0,
						TotalPrice: 0,
						Qty:        y.Qty,
					}
					orderNo++
				}
			}
		}
	}

	// แปลง map กลับเป็น slice
	var complementarys []product_model.ResCleanedComplementarysOrder
	for _, v := range complementaryMap {
		complementarys = append(complementarys, v)
	}
	return complementarys, nil
}
