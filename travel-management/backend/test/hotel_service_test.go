package test

import (
	"testing"

	"github.com/mig3177/travelmanagement/domain/model"
)

type HotelRepositoryMock struct {
	T *testing.T
}

func (repo HotelRepositoryMock) Create(*model.Hotel) (*model.Hotel, error) {
	return nil, nil
}
func (repo HotelRepositoryMock) Update(*model.Hotel) (*model.Hotel, error) {
	return nil, nil
}
func (repo HotelRepositoryMock) Delete(uint) error {
	return nil
}
func (repo HotelRepositoryMock) ListAll() ([]*model.Hotel, error) {
	return nil, nil
}
func (repo HotelRepositoryMock) FindByID(uint) (*model.Hotel, error) {
	return nil, nil
}
func (repo HotelRepositoryMock) Count() (int64, error) {
	return 0, nil
}

type MessageServiceMock struct {
	T *testing.T
}

func (msg MessageServiceMock) PublishAsJSON(interface{}) error {
	msg.T.Fatalf("ICH BIN DA - WER NOCH?")
	return nil
}

func HotelServiceTest(t *testing.T) {
	//s := application.NewHotelService(HotelRepositoryMock{T: t}, kafka.TravelEventImpl{MsgService: MessageServiceMock{T: t}})

	//	s.RemoveHotel(5, "d")
	t.Fatalf("ajshdas %s", "dafadf")

}
