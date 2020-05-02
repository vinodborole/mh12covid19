package database

import (
	"github.com/jinzhu/gorm"
)

//User account info
type User struct {
	gorm.Model

	Name     string
	Phone    string
	Email    string
	Password string
	Token    string `sql:"-"`
}

//Ward ward
type Ward struct {
	gorm.Model

	Name string
	Code string
}

//DistrictCase district case
type DistrictCase struct {
	gorm.Model

	Confirmed int
	Death     int
	Recovered int
}

//AundhBaner aundh baner
type AundhBaner struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//BhawaniPeth bhawani peth
type BhawaniPeth struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//Bibwewadi bibwewadi
type Bibwewadi struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//DhankawadiSahakarnagar dhankawadi
type DhankawadiSahakarnagar struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//DholePatil dhole patil
type DholePatil struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//HadapsarMundhwa hadapsar mundhwa
type HadapsarMundhwa struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//KasbaVisharambaghwada kasba
type KasbaVisharambaghwada struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//KondhwaYewalewadi kondhwa
type KondhwaYewalewadi struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//KothrudBavdhan kothrud
type KothrudBavdhan struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//NagarRoadWadgoansheri nagar road
type NagarRoadWadgoansheri struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//ShivajiNagarGholeRoad shivaji nagar
type ShivajiNagarGholeRoad struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//SinghagadRoad singhagad
type SinghagadRoad struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//WanawadiRamtekdi wanawadi
type WanawadiRamtekdi struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//WarjeKarveNagar warje
type WarjeKarveNagar struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//YerwadaKalasDhanori yerwada
type YerwadaKalasDhanori struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}

//PuneRural pune rural
type PuneRural struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
}
