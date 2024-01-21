package application_test

import (
	"time"

	"github.com/mig3177/travelmanagement/domain/model"
)

var Travel1 = model.Travel{Id: 2, From: time.Date(2022, 5, 20, 1, 0, 0, 0, time.Local), To: time.Date(2022, 5, 25, 1, 0, 0, 0, time.Local), Price: 587}
var Travel2 = model.Travel{Id: 11, From: time.Date(2022, 5, 10, 1, 0, 0, 0, time.Local), To: time.Date(2022, 5, 12, 1, 0, 0, 0, time.Local), Price: 888}
var Travel3 = model.Travel{Id: 14, From: time.Date(2022, 8, 10, 1, 0, 0, 0, time.Local), To: time.Date(2022, 8, 15, 1, 0, 0, 0, time.Local), Price: 693}
var Travel4 = model.Travel{Id: 28, From: time.Date(2023, 5, 23, 1, 0, 0, 0, time.Local), To: time.Date(2023, 5, 25, 1, 0, 0, 0, time.Local), Price: 355}
var Travel5 = model.Travel{Id: 38, From: time.Date(2023, 4, 1, 1, 0, 0, 0, time.Local), To: time.Date(2023, 4, 4, 1, 0, 0, 0, time.Local), Price: 748}

var Tag1 = model.Tag{Id: 1, Name: "City"}
var Tag2 = model.Tag{Id: 2, Name: "Beach"}
var Tag3 = model.Tag{Id: 3, Name: "Sport"}

var Hotel1 = model.Hotel{Id: 1, Name: "Hilton Diagonal Mar Barcelona", Address: model.Address{Land: "USA"},
	Tags: []*model.Tag{&Tag1, &Tag3}, Travels: []*model.Travel{&Travel1, &Travel2}}
var Hotel2 = model.Hotel{Id: 147, Name: "Europejski", Address: model.Address{Land: "Polen"},
	Tags: []*model.Tag{&Tag2}, Travels: []*model.Travel{&Travel5}}
var Hotel3 = model.Hotel{Id: 211, Name: "Serras Barcelona", Address: model.Address{Land: "Spanien"},
	Tags: []*model.Tag{&Tag1}, Travels: []*model.Travel{&Travel3, &Travel4}}
