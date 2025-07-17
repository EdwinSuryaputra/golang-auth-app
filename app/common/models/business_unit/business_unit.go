package businessunitmodel

type BusinessUnit struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Level string `json:"level"`
	
}

func GetArea() map[int32]BusinessUnit {
	return map[int32]BusinessUnit{
		1: {1, "AREA 1", "Area"},
		2: {2, "AREA 2", "Area"},
		3: {3, "AREA 3", "Area"},
	}
}

func GetRegions() map[int32]BusinessUnit {
	return map[int32]BusinessUnit{
		1: {1, "Medan", "Region"},
		2: {2, "Jakarta", "Region"},
	}
}

func GetWitels() map[int32]BusinessUnit {
	return map[int32]BusinessUnit{
		1: {1, "WITEL A", "Witel"},
		2: {2, "WITEL B", "Witel"},
		3: {3, "WITEL C", "Witel"},
		4: {4, "WITEL D", "Witel"},
	}
}

func GetWarehouses() map[int32]BusinessUnit {
	return map[int32]BusinessUnit{
		1: {1, "Warehouse 1", "Warehouse"},
		2: {2, "Warehouse 2", "Warehouse"},
	}
}
