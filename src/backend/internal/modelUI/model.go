package modelUI

type Comic struct {
	Id         int    `json:"_"`
	Name       string `json:"name"`
	City       string `json:"city"`
	Sentence   string `json:"sentence"`
	CountOfKek int    `json:"countOfKek"`
	Image      string `json:"image"`
}

type Poster struct {
	Id      int    `json:"_"`
	Image   string `json:"image"`
	IdEvent int    `json:"id_event"`
	IdPlace int    `json:"id_place"`
	Date    string `json:"date"`
}

type SingIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SingUp struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Person struct {
	Id    int    `json:"_"`
	Login string `json:"login"`
	Email string `json:"email"`
	Salt  string `json:"salt"`
	Hash  string `json:"hash"`
	Role  int    `json:"role"`
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
