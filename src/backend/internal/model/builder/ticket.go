package builder

import "github.com/thirteenths/WEB_BMSTU23/backend/internal/model"

type TicketBuilder struct {
	Id       int
	IdPoster int
	IdPerson int
}

func (b *TicketBuilder) Build() *model.Ticket {
	return &model.Ticket{
		Id:       b.Id,
		IdPoster: b.IdPoster,
		IdPerson: b.IdPerson,
	}
}

type TicketMother struct {
}

func (m *TicketMother) Obj0() *model.Ticket {
	builder := TicketBuilder{Id: 1, IdPoster: 1, IdPerson: 1}
	return builder.Build()
}

func (m *TicketMother) Obj1() *model.Ticket {
	builder := TicketBuilder{Id: 2, IdPoster: 1, IdPerson: 2}
	return builder.Build()
}

func (m *TicketMother) Obj2() *model.Ticket {
	builder := TicketBuilder{Id: 3, IdPoster: 1, IdPerson: 3}
	return builder.Build()
}

func (m *TicketMother) Obj3() *model.Ticket {
	builder := TicketBuilder{Id: 4, IdPoster: 2, IdPerson: 1}
	return builder.Build()
}

func (m *TicketMother) Obj4() *model.Ticket {
	builder := TicketBuilder{Id: 5, IdPoster: 2, IdPerson: 4}
	return builder.Build()
}
