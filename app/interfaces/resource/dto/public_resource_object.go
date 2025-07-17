package dto

type ResourcePubObj struct {
	Applications []*ApplicationPubObj `json:"applications"`
}

type ApplicationPubObj struct {
	Id         string                `json:"id"`
	Name       string                `json:"name"`
	PublicName string                `json:"publicName"`
	Menus      []*MenuResourcePubObj `json:"menus"`
}

type MenuResourcePubObj struct {
	Id            string                   `json:"id"`
	ApplicationId string                   `json:"applicationId"`
	Name          string                   `json:"name"`
	PublicName    string                   `json:"publicName"`
	SubMenus      []*SubMenuResourcePubObj `json:"submenus"`
}

type SubMenuResourcePubObj struct {
	Id         string                    `json:"id"`
	MenuId     string                    `json:"menuId"`
	Name       string                    `json:"name"`
	PublicName string                    `json:"publicName"`
	Functions  []*FunctionResourcePubObj `json:"functions"`
}

type FunctionResourcePubObj struct {
	Id         string `json:"id"`
	SubmenuId  string `json:"submenuId"`
	Name       string `json:"name"`
	PublicName string `json:"publicName"`
}
