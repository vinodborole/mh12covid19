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

//AundhBaner aundh baner
type AundhBaner struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//BhawaniPeth bhawani peth
type BhawaniPeth struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//Bibwewadi bibwewadi
type Bibwewadi struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//DhankawadiSahakarnagar dhankawadi
type DhankawadiSahakarnagar struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//DholePatil dhole patil
type DholePatil struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//HadapsarMundhwa hadapsar mundhwa
type HadapsarMundhwa struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//KasbaVisharambaghwada kasba
type KasbaVisharambaghwada struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//KondhwaYewalewadi kondhwa
type KondhwaYewalewadi struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//KothrudBavdhan kothrud
type KothrudBavdhan struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//NagarRoadWadgoansheri nagar road
type NagarRoadWadgoansheri struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//ShivajiNagarGholeRoad shivaji nagar
type ShivajiNagarGholeRoad struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//SinghagadRoad singhagad
type SinghagadRoad struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//WanawadiRamtekdi wanawadi
type WanawadiRamtekdi struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//WarjeKarveNagar warje
type WarjeKarveNagar struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//YerwadaKalasDhanori yerwada
type YerwadaKalasDhanori struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}

//PuneRural pune rural
type PuneRural struct {
	gorm.Model

	Active    int
	Confirmed int
	Death     int
	Recovered int
}
