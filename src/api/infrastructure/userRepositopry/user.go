package userRepositopry

type User struct {
	Name         string
	Email        string
	PasswordHash string
	TokenHash    string
	SexCode      uint8
	PrefCode     uint32
	CityCode     uint32
	WardCode     uint32
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    *string // may be null
}
