package wardRepository

type Wards struct {
	WardLists []Ward
}

type Ward struct {
	Id       uint32 `json:"id" db:"id"`
	Code     uint32 `json:"code" db:"ward_code"`
	CityCode uint32 `json:"city_code" db:"city_code"`
	Name     string `json:"name" db:"ward"`
}
