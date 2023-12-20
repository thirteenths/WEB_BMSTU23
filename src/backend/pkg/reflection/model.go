package reflection

type Public struct {
	Id   int    `json:"_"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type Queue struct {
	Id         int  `json:"_"`
	IdStudent  int  `json:"idstudent"`
	IdSchedule int  `json:"idschedule"`
	Status     bool `json:"status"`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// И Методы делающие Date строкой для постгрес

type Schedule struct {
	Id      int    `json:"_"`
	Subject string `json:"subject"`
	Date    string `json:"date"`
}

type Subscription struct {
	Id        int `json:"_"`
	IdStudent int `json:"id_student"`
	IdPublic  int `json:"id_public"`
}

type Student struct {
	Id    int    `json:"_"`
	Name  string `json:"name"`
	Class string `json:"class"` // В скрипте grope но я так понимаю это очепятка
	Role  string `json:"role"`  // = "student"
}
