package entities

type SmallUser struct {
	Login    string `json:"login"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename"`
	MacAddress string `json:"mac_address"`
	Role     int    `json:"privilege_level"`
}

type DefaultPhone struct {
	ID     int64  `json:"id"`
	Number string `json:"number"`
}

type YandexUser struct {
	FirstName       string       `json:"first_name"`
	LastName        string       `json:"last_name"`
	DisplayName     string       `json:"display_name"`
	Emails          []string     `json:"emails"`
	DefaultEmail    string       `json:"default_email"`
	DefaultPhone    DefaultPhone `json:"default_phone"`
	RealName        string       `json:"real_name"`
	IsAvatarEmpty   bool         `json:"is_avatar_empty"`
	DefaultAvatarID string       `json:"default_avatar_id"`
	Login           string       `json:"login"`
	OldSocialLogin  string       `json:"old_social_login"`
	Sex             string       `json:"sex"`
	ID              string       `json:"id"`
	ClientID        string       `json:"client_id"`
	Psuid           string       `json:"psuid"`
}
