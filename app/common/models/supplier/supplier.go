package suppliermodel

type Supplier struct {
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	SupplierCode string `json:"supplier_code"`
}

// func GetSuppliers() map[int32]Supplier {
// 	return map[int32]Supplier{
// 		1: {1, "ZTE"},
// 		2: {2, "HKM"},
// 		3: {3, "HUAWEI"},
// 		4: {4, "FIBERHOME"},
// 		5: {5, "ALCATEL-LUCENT"},
// 		6: {6, "NOKIA"},
// 		7: {7, "TELKOMSEL"},
// 	}
// }
