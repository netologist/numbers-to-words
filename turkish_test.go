package n2w_test

import (
	"github.com/hasanozgan/numbers-to-words"
	"testing"
)

func TestTurkishWords(t *testing.T) {
	tests := testcases{
		{name: "number -1", number: -1, expected: "EKSİ BİR"},
		{name: "number 0", number: 0, expected: "SIFIR"},
		{name: "number 0.05", number: 0.05, expected: "SIFIR NOKTA BEŞ"},
		{name: "number 1", number: 1, expected: "BİR"},
		{name: "number 2", number: 2, expected: "İKİ"},
		{name: "number 3", number: 3, expected: "ÜÇ"},
		{name: "number 4", number: 4, expected: "DÖRT"},
		{name: "number 5", number: 5, expected: "BEŞ"},
		{name: "number 6", number: 6, expected: "ALTI"},
		{name: "number 7", number: 7, expected: "YEDİ"},
		{name: "number 8", number: 8, expected: "SEKİZ"},
		{name: "number 9", number: 9, expected: "DOKUZ"},
		{name: "number 10", number: 10, expected: "ON"},
		{name: "number 20", number: 20, expected: "YİRMİ"},
		{name: "number 30", number: 30, expected: "OTUZ"},
		{name: "number 40", number: 40, expected: "KIRK"},
		{name: "number 50", number: 50, expected: "ELLİ"},
		{name: "number 60", number: 60, expected: "ALTMIŞ"},
		{name: "number 70", number: 70, expected: "YETMİŞ"},
		{name: "number 80", number: 80, expected: "SEKSEN"},
		{name: "number 90", number: 90, expected: "DOKSAN"},
		{name: "number 39", number: 39, expected: "OTUZ DOKUZ"},
		{name: "number 48", number: 48, expected: "KIRK SEKİZ"},
		{name: "number 24", number: 24, expected: "YİRMİ DÖRT"},
		{name: "number 100", number: 100, expected: "YÜZ"},
		{name: "number 123", number: 123, expected: "YÜZ YİRMİ ÜÇ"},
		{name: "number 968", number: 968, expected: "DOKUZ YÜZ ALTMIŞ SEKİZ"},
		{name: "number 1.000", number: 1000, expected: "BİN"},
		{name: "number 1.624", number: 1624, expected: "BİN ALTI YÜZ YİRMİ DÖRT"},
		{name: "number 8.596", number: 8596, expected: "SEKİZ BİN BEŞ YÜZ DOKSAN ALTI"},
		{name: "number 10.000", number: 10000, expected: "ON BİN"},
		{name: "number 100.000", number: 100000, expected: "YÜZ BİN"},
		{name: "number 1.000.000", number: 1000000, expected: "BİR MİLYON"},
		{name: "number 1.000.000.000", number: 1000000000, expected: "BİR MİLYAR"},
		{name: "number 1.000.000.000.000", number: 1000000000000, expected: "BİR TRİLYON"},
		{name: "number 1.000.000.000.000.000", number: 1000000000000000, expected: "BİR KATTRİLYON"},
		{name: "number 9.876.543.210.987.654", number: 9876543210987654, expected: "DOKUZ KATTRİLYON SEKİZ YÜZ YETMİŞ ALTI TRİLYON BEŞ YÜZ KIRK ÜÇ MİLYAR İKİ YÜZ ON MİLYON DOKUZ YÜZ SEKSEN YEDİ BİN ALTI YÜZ ELLİ DÖRT"},
	}

	for _, test := range tests {
		t.Run(test.name+" should be "+test.expected, func(t *testing.T) {
			if actual, err := n2w.ToTurkishWords(test.number); err != nil {
				t.Errorf("Failed %s with %v", actual, err)
			} else if test.expected != actual {
				t.Errorf("Failed %s", actual)
			}
		})
	}
}
