package gozodiac

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const shortForm = "2006-Jan-02"

func TestGetZodiacSign(t *testing.T) {
	// sign := GetZodiacSign(time.Now())
	// log.Printf("sign: %v", sign)

	testTime, err := time.Parse(shortForm, "2013-Feb-03")
	assert.NoError(t, err)
	sign := GetZodiacSign(testTime)
	assert.Equal(t, []ZodiacSign{Aquarius}, sign)

	testTime, err = time.Parse(shortForm, "2001-Nov-15")
	assert.NoError(t, err)
	sign = GetZodiacSign(testTime)
	assert.Equal(t, []ZodiacSign{Scorpio}, sign)

	testTime, err = time.Parse(shortForm, "1942-Oct-11")
	assert.NoError(t, err)
	sign = GetZodiacSign(testTime)
	assert.Equal(t, []ZodiacSign{Libra}, sign)
}

func TestGetChineseZodiacSign(t *testing.T) {
	// sign := GetZodiacSign(time.Now())
	// log.Printf("sign: %v", sign)

	testTime, err := time.Parse(shortForm, "1963-Dec-18")
	assert.NoError(t, err)
	sign := GetChineseZodiacSign(testTime)
	log.Println("chinese zodiac sign:", sign)
	// assert.Equal(t, []ZodiacSign{Aquarius}, sign)

	// testTime, err = time.Parse(shortForm, "2001-Nov-15")
	// assert.NoError(t, err)
	// sign = GetZodiacSign(testTime)
	// assert.Equal(t, []ZodiacSign{Scorpio}, sign)

	// testTime, err = time.Parse(shortForm, "1942-Oct-11")
	// assert.NoError(t, err)
	// sign = GetZodiacSign(testTime)
	// assert.Equal(t, []ZodiacSign{Libra}, sign)
}
