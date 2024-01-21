package application_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/mig3177/travelmanagement/adapter/kafka"
	"github.com/mig3177/travelmanagement/adapter/kafka/dto"
	"github.com/mig3177/travelmanagement/application"
	"github.com/mig3177/travelmanagement/domain/model"
)

type HotelRepositoryMock struct {
	T *testing.T
}

func (repo HotelRepositoryMock) Create(*model.Hotel) (*model.Hotel, error) {
	return nil, nil
}
func (repo HotelRepositoryMock) Update(*model.Hotel) (*model.Hotel, error) {
	return nil, errors.New("DFDf")
}
func (repo HotelRepositoryMock) Delete(uint) error {
	return nil
}
func (repo HotelRepositoryMock) ListAll() ([]*model.Hotel, error) {
	hotels := []*model.Hotel{&Hotel1, &Hotel2, &Hotel3}
	return hotels, nil
}
func (repo HotelRepositoryMock) FindByID(uint) (*model.Hotel, error) {
	return nil, errors.New("Connection Lost")
}
func (repo HotelRepositoryMock) Count() (int64, error) {
	return 0, nil
}

type MessageServiceMock struct {
	T *testing.T
}

func (msg MessageServiceMock) PublishAsJSON(obj interface{}) error {
	event, _ := obj.(dto.HotelTravelEvent)
	if event.Type != "UpdateError" {
		msg.T.Fatalf("Expect len %s, got %s", "UpdateError", event.Type)
	}
	return nil
}

func TestHotelServiceTest(t *testing.T) {
	s := application.NewHotelService(HotelRepositoryMock{T: t}, kafka.TravelEventImpl{MsgService: MessageServiceMock{T: t}})
	res, _ := s.FindHotelTravel("", "", nil, nil, []uint{1})
	if len(res) != 2 {
		t.Fatalf("Expect len %d, got %d", 2, len(res))
	}
	res, _ = s.FindHotelTravel("", "", nil, nil, []uint{})
	if len(res) != 3 {
		t.Fatalf("Expect len %d, got %d", 3, len(res))
	}
	res, _ = s.FindHotelTravel("Barce", "", nil, nil, []uint{})
	if len(res) != 2 {
		t.Fatalf("Expect len %d, got %d", 2, len(res))
	}
	res, _ = s.FindHotelTravel("Barce", "Spa", nil, nil, []uint{})
	if len(res) != 1 {
		t.Fatalf("Expect len %d, got %d", 1, len(res))
	}
	from := time.Date(2022, 5, 1, 0, 0, 0, 0, time.Local)
	to := time.Date(2022, 5, 30, 0, 0, 0, 0, time.Local)
	res, _ = s.FindHotelTravel("", "", &from, &to, nil)
	if len(res) != 1 {
		t.Fatalf("Expect len %d, got %d", 1, len(res))
	}
	from = time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local)
	to = time.Date(2022, 1, 5, 0, 0, 0, 0, time.Local)
	res, _ = s.FindHotelTravel("", "", &from, &to, nil)
	fmt.Println(res)
	if len(res) != 0 {
		t.Fatalf("Expect len %d, got %d", 0, len(res))
	}
	_, err := s.GetHotel(1)
	if err == nil {
		t.Fatalf("Expect Error, got %s", err.Error())
	}

	s.UpdateHotel(&Hotel1)

}
