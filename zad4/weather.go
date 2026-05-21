package main

import "gorm.io/gorm"

type Weather struct {
	gorm.Model
	IdStacji           uint    `json:"id_stacji,string"`
	Stacja             string  `json:"stacja"`
	DataPomiaru        string  `json:"data_pomiaru"`
	GodzinaPomiaru     uint    `json:"godzina_pomiaru,string"`
	Temperatura        float64 `json:"temperatura,string"`
	PredkoscWiatru     float64 `json:"predkosc_wiatru,string"`
	KierunekWiatru     uint    `json:"kierunek_wiatru,string"`
	WilgotnoscWzgledna float64 `json:"wilgotnosc_wzgledna,string"`
	SumaOpadu          float64 `json:"suma_opadu,string"`
	Cisnienie          float64 `json:"cisnienie,string"`
}
