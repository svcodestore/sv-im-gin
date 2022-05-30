package utils

import (
	"github.com/thanhpk/randstr"
	"math/rand"
	"time"
)

func RandRange(min, max int) (n int) {
	rand.Seed(time.Now().Unix())
	n = rand.Intn(max-min) + min
	return
}

func GenerateClientId() string {
	return randstr.Hex(10)
}

func GenerateClientSecret() string {
	return randstr.Hex(20)
}