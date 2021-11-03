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
	WardCode      *uint32
	Sex           []UserSex
	Prefecture    []UserPrefecture
	City          []UserCity
	Ward          []UserWard
	RememberToken *string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     *string // may be null
}

type UserSex struct {
	Code uint8
	Name string
}

type UserPrefecture struct {
	PrefCode uint32
	Name     string
}

type UserCity struct {
	CityCode uint32
	Name     string
}

type UserWard struct {
	WardCode *uint32
	Name     *string
}

type UpdateUser struct {
	Name     string  `json:"name"`
	SexCode  uint8   `json:"sex_code"`
	Email    string  `json:"email"`
	PrefCode uint32  `json:"pref_code"`
	CityCode uint32  `json:"city_code"`
	WardCode *uint32 `json:"ward_code"`
}

type UpdatePassword struct {
	OldPassword  string `json:"old_password"`
	NewPassword  string `json:"new_password"`
	RefreshToken string `json:"refresh_token"`
	GrantType    string `json:"grant_type"`
}
