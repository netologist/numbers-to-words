package n2w

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ToEnglishWords(number float64) (string, error) {
	result := ""
	strNumber := fmt.Sprintf("%.2f", number)
	sections := strings.Split(strNumber, ".")

	if integral, err := strconv.Atoi(sections[0]); err != nil {
		return "", err
	} else {
		result += withEnglishStrategy(integral)
	}

	if fractional, err := strconv.Atoi(sections[1]); err != nil {
		return "", err
	} else if fractional > 0 {
		result += " POINT "
		result += withEnglishStrategy(fractional)
	}

	return result, nil

}

func withEnglishStrategy(number int) string {
	result := ""
	units := []string{"", "ONE ", "TWO ", "THREE ", "FOUR ", "FIVE ", "SIX ", "SEVEN ", "EIGHT ", "NINE "}
	tens := []string{"", "TEN ", "TWENTY ", "THIRTY ", "FORTY ", "FIFTY ", "SIXTY ", "SEVENTY ", "EIGHTY ", "NINETY "}
	teens := []string{"TEN ", "ELEVEN ", "TWELVE ", "THIRTEEN ", "FOURTEEN ", "FIFTEEN ", "SIXTEEN ", "SEVENTEEN ", "EIGHTEEN ", "NINETEEN "}
	others := []string{"", "HUNDERED ", "THOUSAND ", "MILLION ", "BILLION ", "TRILLION ", "QUADRILION "}
	hasParentDigit := false
	hasTeens := false

	if number == 0 {
		return "ZERO"
	}

	if number < 0 {
		result += "NEGATIVE "
		number *= -1
	}
	str_number := strconv.Itoa(number)

	power := 1

	for i := 0; i < len(str_number); i++ {
		order := len(str_number[i:]) % 3
		number := str_number[i] - '0'

		switch order {
		// units
		case 1:
			power = int(math.Ceil(float64(len(str_number[i:])) / 3))

			if hasTeens {
				result += teens[number]
				hasTeens = false
				hasParentDigit = true
				continue
			}

			result += units[number]
			hasParentDigit = !(number == 0 && !hasParentDigit)

		// tens
		case 2:

			power = int(math.Ceil(float64(len(str_number[i:])) / 3))
			if number == 1 {
				hasTeens = true
				continue
			}
			result += tens[number]
			hasParentDigit = !(number == 0 && !hasParentDigit)

		// other
		case 0:
			if power > 1 && hasParentDigit {
				result += others[power]
			}

			if number > 0 {
				result += units[number]
				newPower := int(math.Ceil(float64(len(str_number[i:])) / 3))
				if newPower > power {
					result += others[power]
				} else {
					result += others[1]
				}

				power = newPower
			}
			hasParentDigit = number > 0

		}
	}

	return strings.TrimSpace(result)
}
