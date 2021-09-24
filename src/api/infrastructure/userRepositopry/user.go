package userRepositopry

type User struct {
	Id            int
	Name          string
	Email         string
	PasswordHash  string
	TokenHash     string
	SexCode       uint8
	PrefCode      uint32
	CityCode      uint32
	WardCode      uint32
	RememberToken *string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     *string // may be null
}
