package domain

import (
	"github.com/dgrijalva/jwt-go"
)

//Token auth token
type Token struct {
	UserID uint
	jwt.StandardClaims
}

//User account info
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

//Ward ward info
type Ward struct {
	Name string
	Code string
}

//DistrictCase district case
type DistrictCase struct {
	TotalPositive int
	Active        int
	Confirmed     int
	Death         int
	Recovered     int
	Date          string
}

//WardCase ward case
type WardCase struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//AundhBaner aundh baner
type AundhBaner struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//BhawaniPeth bhawani peth
type BhawaniPeth struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//Bibwewadi bibwewadi
type Bibwewadi struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//DhankawadiSahakarnagar dhankawadi
type DhankawadiSahakarnagar struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//DholePatil dhole patil
type DholePatil struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//HadapsarMundhwa hadapsar mundhwa
type HadapsarMundhwa struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//KasbaVisharambaghwada kasba
type KasbaVisharambaghwada struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//KondhwaYewalewadi kondhwa
type KondhwaYewalewadi struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//KothrudBavdhan kothrud
type KothrudBavdhan struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//NagarRoadWadgoansheri nagar road
type NagarRoadWadgoansheri struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//ShivajiNagarGholeRoad shivaji nagar
type ShivajiNagarGholeRoad struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//SinghagadRoad singhagad
type SinghagadRoad struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//WanawadiRamtekdi wanawadi
type WanawadiRamtekdi struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//WarjeKarveNagar warje
type WarjeKarveNagar struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//YerwadaKalasDhanori yerwada
type YerwadaKalasDhanori struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}

//PuneRural pune rural
type PuneRural struct {
	Active    int
	Confirmed int
	Death     int
	Recovered int
	Code      string
	Name      string
	Date      string
}
