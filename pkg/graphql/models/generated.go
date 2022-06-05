// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Clothes struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	IDStore  int     `json:"id_store"`
	Category int     `json:"category"`
	Size     string  `json:"size"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type Store struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	StoreID     int    `json:"storeId"`
	Cnpj        string `json:"cnpj"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Category    int    `json:"category"`
}

type CreateCategoryInput struct {
	Name string `json:"name"`
}

type CreateClothInput struct {
	Name     string  `json:"name"`
	IDStore  int     `json:"id_store"`
	Category int     `json:"category"`
	Size     string  `json:"size"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type CreateStoreInput struct {
	Name        string `json:"name"`
	Cnpj        string `json:"cnpj"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Category    int    `json:"category"`
}

type UpdateCategoryInput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateClothInput struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	IDStore  int     `json:"id_store"`
	Category int     `json:"category"`
	Size     string  `json:"size"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

type UpdateStoreInput struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Cnpj        string `json:"cnpj"`
	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Category    int    `json:"category"`
}

type Weekdays string

const (
	WeekdaysMonday    Weekdays = "MONDAY"
	WeekdaysTuesday   Weekdays = "TUESDAY"
	WeekdaysWednesday Weekdays = "WEDNESDAY"
	WeekdaysThursday  Weekdays = "THURSDAY"
	WeekdaysFriday    Weekdays = "FRIDAY"
	WeekdaysSaturday  Weekdays = "SATURDAY"
	WeekdaysSunday    Weekdays = "SUNDAY"
)

var AllWeekdays = []Weekdays{
	WeekdaysMonday,
	WeekdaysTuesday,
	WeekdaysWednesday,
	WeekdaysThursday,
	WeekdaysFriday,
	WeekdaysSaturday,
	WeekdaysSunday,
}

func (e Weekdays) IsValid() bool {
	switch e {
	case WeekdaysMonday, WeekdaysTuesday, WeekdaysWednesday, WeekdaysThursday, WeekdaysFriday, WeekdaysSaturday, WeekdaysSunday:
		return true
	}
	return false
}

func (e Weekdays) String() string {
	return string(e)
}

func (e *Weekdays) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Weekdays(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Weekdays", str)
	}
	return nil
}

func (e Weekdays) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
