package gozodiac

import "time"

// ZodiacSign represents the zodiac signs
type ZodiacSign int

func (s ZodiacSign) String() string {
	switch s {
	case Aries:
		return "Aries"
	case Taurus:
		return "Taurus"
	case Gemini:
		return "Gemini"
	case Cancer:
		return "Cancer"
	case Leo:
		return "Leo"
	case Virgo:
		return "Virgo"
	case Libra:
		return "Libra"
	case Scorpio:
		return "Scorpio"
	case Sagittarius:
		return "Sagittarius"
	case Capricorn:
		return "Capricorn"
	case Aquarius:
		return "Aquarius"
	case Pisces:
		return "Pisces"
	}
	return ""
}

const (
	// Aries zodiac sign
	Aries ZodiacSign = iota
	// Taurus zodiac sign
	Taurus
	// Gemini zodiac sign
	Gemini
	// Cancer zodiac sign
	Cancer
	// Leo zodiac sign
	Leo
	// Virgo zodiac sign
	Virgo
	// Libra zodiac sign
	Libra
	// Scorpio zodiac sign
	Scorpio
	// Sagittarius zodiac sign
	Sagittarius
	// Capricorn zodiac sign
	Capricorn
	// Aquarius zodiac sign
	Aquarius
	// Pisces zodiac sign
	Pisces
)

// GetZodiacSign returns the matching zodiac sign
func GetZodiacSign(date time.Time) []ZodiacSign {
	switch date.Month() {
	case time.March:
		if date.Day() < 20 {
			return []ZodiacSign{Pisces}
		} else if date.Day() == 20 {
			return []ZodiacSign{Pisces, Aries}
		}
		return []ZodiacSign{Aries}
	case time.April:
		if date.Day() < 19 {
			return []ZodiacSign{Aries}
		} else if date.Day() == 19 {
			return []ZodiacSign{Aries, Taurus}
		}
		return []ZodiacSign{Taurus}
	case time.May:
		if date.Day() < 20 {
			return []ZodiacSign{Taurus}
		} else if date.Day() == 20 {
			return []ZodiacSign{Taurus, Gemini}
		}
		return []ZodiacSign{Gemini}
	case time.June:
		if date.Day() < 21 {
			return []ZodiacSign{Gemini}
		} else if date.Day() == 21 {
			return []ZodiacSign{Gemini, Cancer}
		}
		return []ZodiacSign{Cancer}
	case time.July:
		if date.Day() < 22 {
			return []ZodiacSign{Cancer}
		} else if date.Day() == 22 {
			return []ZodiacSign{Cancer, Leo}
		}
		return []ZodiacSign{Leo}
	case time.August:
		if date.Day() < 22 {
			return []ZodiacSign{Leo}
		} else if date.Day() == 22 {
			return []ZodiacSign{Leo, Virgo}
		}
		return []ZodiacSign{Virgo}
	case time.September:
		if date.Day() < 22 {
			return []ZodiacSign{Virgo}
		} else if date.Day() == 22 {
			return []ZodiacSign{Virgo, Libra}
		}
		return []ZodiacSign{Libra}
	case time.October:
		if date.Day() < 23 {
			return []ZodiacSign{Libra}
		} else if date.Day() == 23 {
			return []ZodiacSign{Libra, Scorpio}
		}
		return []ZodiacSign{Scorpio}
	case time.November:
		if date.Day() < 22 {
			return []ZodiacSign{Scorpio}
		} else if date.Day() == 22 {
			return []ZodiacSign{Scorpio, Sagittarius}
		}
		return []ZodiacSign{Sagittarius}
	case time.December:
		if date.Day() < 21 {
			return []ZodiacSign{Sagittarius}
		} else if date.Day() == 21 {
			return []ZodiacSign{Sagittarius, Capricorn}
		}
		return []ZodiacSign{Capricorn}
	case time.January:
		if date.Day() < 20 {
			return []ZodiacSign{Capricorn}
		} else if date.Day() == 20 {
			return []ZodiacSign{Capricorn, Aquarius}
		}
		return []ZodiacSign{Aquarius}
	case time.February:
		if date.Day() < 18 {
			return []ZodiacSign{Aquarius}
		} else if date.Day() == 18 {
			return []ZodiacSign{Aquarius, Pisces}
		}
		return []ZodiacSign{Pisces}
	}
	return []ZodiacSign{Pisces}
}
