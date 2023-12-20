package builder

import "github.com/thirteenths/WEB_BMSTU23/backend/internal/model"

type PlaceBuilder struct {
	Id          int
	Name        string
	CountTicket int
}

func (b *PlaceBuilder) Build() *model.Place {
	return &model.Place{
		Id:          b.Id,
		Name:        b.Name,
		CountTicket: b.CountTicket,
	}
}

type PlaceMother struct {
}

func (m *PlaceMother) Obj0() *model.Place {
	builder := PlaceBuilder{
		Id:          1,
		Name:        "МЗЦКИ",
		CountTicket: 50}
	return builder.Build()
}

func (m *PlaceMother) Obj1() *model.Place {
	builder := PlaceBuilder{
		Id:          2,
		Name:        "БЗЦКИ",
		CountTicket: 100}
	return builder.Build()
}
