package conf

//omitempy если ноль будет то ничего не запишет
type conf struct {
	Api    api    `json:"api,omitempty"`
	Server server `json:"server,omitempty"`
}

//структура конфига бд
type api struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}

//структура конфига сервера
type server struct {
	Host string `json:"host,omitempty"`
	Port string `json:"port,omitempty"`
}
