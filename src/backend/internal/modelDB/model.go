package modelDB

type Comic struct {
	Id       int    `json:"id"`
	IdImage  int    `json:"id-image"`
	Name     string `json:"name"`
	City     string `json:"city"`
	Sentence string `json:"sentence"`
	CountKek int    `json:"count_kek"`
}

type Image struct {
	Id   int    `json:"id"`
	Path string `json:"path"`
}

type Person struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
	Salt  string `json:"salt"`
	Hash  string `json:"hash"`
	Role  int    `json:"role"`
}

type Poster struct {
	Id      int    `json:"id"`
	IdImage int    `json:"id_image"`
	IdEvent int    `json:"id_event"`
	IdPlace int    `json:"id_place"`
	Date    string `json:"date"`
}

type Event struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	About string `json:"about"`
}

type Place struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	CountTicket int    `json:"count_ticket"`
}

type Ticket struct {
	Id       int `json:"id"`
	IdPoster int `json:"id_poster"`
	IdPerson int `json:"id_person"`
}
