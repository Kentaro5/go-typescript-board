package sexRepository

type Sexes struct {
	SexLists []Sex
}

type Sex struct {
	Id   uint8  `json:"sex_id"`
	Code uint8  `json:"sex_code"`
	Name string `json:"sex_name"`
}
