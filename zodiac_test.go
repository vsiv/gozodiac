package gozodiac

import (
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
}
