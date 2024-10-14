package dto


type LoginRq struct {
	Username  string `json:"username"`
	Password string `json:"password"`
}

type LoginRs struct {
	Token  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	Expiration int64 `json:"expiration"`
}

type RefreshTokenRq struct {
	RefreshToken string `json:"refreshToken"`
}

type MeRs struct {
	Token  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	Expiration int64 `json:"expiration"`
}

type CompanyListRs struct {
	ID  uint64 `json:"id"`
	Name string `json:"name"`

}

type LoginCompanyRs struct {
	CompanyID  uint64 `json:"id"`
}