package cityRepository

type Cities struct {
	CityLists []City
}

type City struct {
	Id       uint32 `json:"id" db:"id"`
	Code     uint32 `json:"code" db:"city_code"`
	PrefCode uint32 `json:"pref_code" db:"pref_code"`
	Name     string `json:"name" db:"city"`
}
