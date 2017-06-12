package gozodiac

import (
	"log"
	"testing"
	"time"
)

func TestGetZodiacSign(t *testing.T) {

	sign := GetZodiacSign(time.Now())
	log.Printf("sign: %v", sign)
}
