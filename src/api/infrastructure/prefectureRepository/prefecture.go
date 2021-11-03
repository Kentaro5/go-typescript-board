package prefectureRepository

type Prefectures struct {
	PrefectureLists []Prefecture
}

type Prefecture struct {
	Id   uint8  `json:"id" db:"id"`
	Code uint32 `json:"code" db:"pref_code"`
	Name string `json:"name" db:"pref"`
}
