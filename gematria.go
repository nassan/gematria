package gematria

import (
	"errors"
	"strings"
)

//Gematria contains the necessary data structures for
//gematria calculation
type Gematria struct {
	//letterValueMap holds each letter of the Hebrew
	//alphabet as a key, and the corresponding numerical
	//value as letterValueMap[key].
	letterValueMap map[string]int

	//orderedLetterArray is an array containing the Hebrew
	//alphabet. It is needed to define the order of the Hebrew
	//alphabet during iteration.
	orderedLetterArray [22]string
}

//Load returns a Gematria object ready for use
func Load() Gematria {
	g := Gematria{}
	g.letterValueMap = make(map[string]int)

	g.letterValueMap["א"] = 1
	g.letterValueMap["ב"] = 2
	g.letterValueMap["ג"] = 3
	g.letterValueMap["ד"] = 4
	g.letterValueMap["ה"] = 5
	g.letterValueMap["ו"] = 6
	g.letterValueMap["ז"] = 7
	g.letterValueMap["ח"] = 8
	g.letterValueMap["ט"] = 9
	g.letterValueMap["י"] = 10
	g.letterValueMap["כ"] = 20
	g.letterValueMap["ל"] = 30
	g.letterValueMap["מ"] = 40
	g.letterValueMap["נ"] = 50
	g.letterValueMap["ס"] = 60
	g.letterValueMap["ע"] = 70
	g.letterValueMap["פ"] = 80
	g.letterValueMap["צ"] = 90
	g.letterValueMap["ק"] = 100
	g.letterValueMap["ר"] = 200
	g.letterValueMap["ש"] = 300
	g.letterValueMap["ת"] = 400

	g.orderedLetterArray[0] = "א"
	g.orderedLetterArray[1] = "ב"
	g.orderedLetterArray[2] = "ג"
	g.orderedLetterArray[3] = "ד"
	g.orderedLetterArray[4] = "ה"
	g.orderedLetterArray[5] = "ו"
	g.orderedLetterArray[6] = "ז"
	g.orderedLetterArray[7] = "ח"
	g.orderedLetterArray[8] = "ט"
	g.orderedLetterArray[9] = "י"
	g.orderedLetterArray[10] = "כ"
	g.orderedLetterArray[11] = "ל"
	g.orderedLetterArray[12] = "מ"
	g.orderedLetterArray[13] = "נ"
	g.orderedLetterArray[14] = "ס"
	g.orderedLetterArray[15] = "ע"
	g.orderedLetterArray[16] = "פ"
	g.orderedLetterArray[17] = "צ"
	g.orderedLetterArray[18] = "ק"
	g.orderedLetterArray[19] = "ר"
	g.orderedLetterArray[20] = "ש"
	g.orderedLetterArray[21] = "ת"

	return g
}

// Get returns a string Gematria representation of a given integer.
func (g Gematria) Get(m int) (string, error) {
	_, result, err := g.getNextLetter(0, m, "")

	//Sanitize exceptions even on an error
	result = sanatize(result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (g Gematria) getNextLetter(depth, m int, a string) (int, string, error) {
	for i := len(g.orderedLetterArray) - 1; i >= 0; i-- {
		if m >= g.letterValueMap[g.orderedLetterArray[i]] {
			a += g.orderedLetterArray[i]
			m -= g.letterValueMap[g.orderedLetterArray[i]]
			break
		}
	}

	if m <= 0 {
		return m, a, nil
	}
	depth++
	if depth > 10 {
		return m, a, errors.New("Gematria phrase length exceeds 10")
	}
	return g.getNextLetter(depth, m, a)
}

func sanatize(a string) string {
	a = strings.Replace(a, "יה", "טו", -1)
	a = strings.Replace(a, "יו", "טז", -1)
	// b := a
	// r, size := utf8.DecodeLastRuneInString(b)
	// c, _ := utf8.DecodeLastRuneInString("ה")

	// if r == c {
	// 	b = b[:len(b)-size]
	// 	r, _ := utf8.DecodeLastRuneInString(b)
	// 	c, _ := utf8.DecodeLastRuneInString("י")
	// 	if r == c {

	// 	}

	// }
	return a
}
