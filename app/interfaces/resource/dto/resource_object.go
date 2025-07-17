package dto

type ResourceObj struct {
	Applications map[int32]*ApplicationResourceObj
}

type ApplicationResourceObj struct {
	Id         int32
	Name       string
	PublicName string
	Menus      map[int32]*MenuResourceObj
}

type MenuResourceObj struct {
	Id            int32
	ApplicationId int32
	Name          string
	PublicName    string
	SubMenus      map[int32]*SubMenuResourceObj
}

type SubMenuResourceObj struct {
	Id         int32
	MenuId     int32
	Name       string
	PublicName string
	Functions  map[int32]*FunctionResourceObj
}

type FunctionResourceObj struct {
	Id           int32
	SubMenuId    int32
	Name         string
	PublicName   string
	ResourceId   int32
	ResourceName string
}
