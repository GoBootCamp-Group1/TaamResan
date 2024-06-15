package config

type Config struct {
	Server Server `json:"server"`
}

type Server struct {
	HttpPort               int    `json:"http_port"`
	Host                   string `json:"host"`
	TokenExpMinutes        uint   `json:"token_exp_minutes"`
	RefreshTokenExpMinutes uint   `json:"refresh_token_exp_minute"`
	TokenSecret            string `json:"token_secret"`
}
