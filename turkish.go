package n2w

import (
	"fmt"
	"strconv"
	"strings"
)

func withTurkishStrategy(number int) string {
	result := ""
	units := []string{"", "BİR ", "İKİ ", "ÜÇ ", "DÖRT ", "BEŞ ", "ALTI ", "YEDİ ", "SEKİZ ", "DOKUZ "}
	tens := []string{"", "ON ", "YİRMİ ", "OTUZ ", "KIRK ", "ELLİ ", "ALTMIŞ ", "YETMİŞ ", "SEKSEN ", "DOKSAN "}
	others := []string{"", "BİN ", "MİLYON ", "MİLYAR ", "TRİLYON ", "KATTRİLYON ", "YÜZ "}
	hasParentDigit := false

	if number == 0 {
		return "SIFIR"
	}

	if number < 0 {
		result += "EKSİ "
		number *= -1
	}
	str_number := strconv.Itoa(number)

	for i := 0; i < len(str_number); i++ {
		order := len(str_number[i:]) % 3
		power := len(str_number[i:]) / 3
		number := str_number[i] - '0'
		switch order {
		// units
		case 1:
			if power == 1 && !hasParentDigit && number == 1 {
				result += units[0]
			} else {
				result += units[number]
			}
			if number == 0 && !hasParentDigit {
				result += others[0]
			} else {
				result += others[power]
			}
			hasParentDigit = false

		// tens
		case 2:
			result += tens[number]
			hasParentDigit = !(number == 0 && !hasParentDigit)

		// others
		case 0:
			if number == 1 {
				result += units[0]
			} else {
				result += units[number]
			}
			if number == 0 {
				result += others[0]
			} else {
				result += others[6]
			}
			hasParentDigit = !(number == 0 && !hasParentDigit)
		}
	}

	return strings.TrimSpace(result)
}

func ToTurkishWords(number float64) (string, error) {
	result := ""
	strNumber := fmt.Sprintf("%.2f", number)
	sections := strings.Split(strNumber, ".")

	if integral, err := strconv.Atoi(sections[0]); err != nil {
		return "", err
	} else {
		result += withTurkishStrategy(integral)
	}

	if fractional, err := strconv.Atoi(sections[1]); err != nil {
		return "", err
	} else if fractional > 0 {
		result += " NOKTA "
		result += withTurkishStrategy(fractional)
	}

	return result, nil
}
