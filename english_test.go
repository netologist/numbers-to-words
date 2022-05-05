package n2w_test

import (
	"github.com/hasanozgan/numbers-to-words"
	"testing"
)

func TestEnglishWords(t *testing.T) {
	tests := testcases{
		{name: "number -1", number: -1, expected: "NEGATIVE ONE"},
		{name: "number 0", number: 0, expected: "ZERO"},
		{name: "number 0.05", number: 0.05, expected: "ZERO POINT FIVE"},
		{name: "number 1", number: 1, expected: "ONE"},
		{name: "number 2", number: 2, expected: "TWO"},
		{name: "number 3", number: 3, expected: "THREE"},
		{name: "number 4", number: 4, expected: "FOUR"},
		{name: "number 5", number: 5, expected: "FIVE"},
		{name: "number 6", number: 6, expected: "SIX"},
		{name: "number 7", number: 7, expected: "SEVEN"},
		{name: "number 8", number: 8, expected: "EIGHT"},
		{name: "number 9", number: 9, expected: "NINE"},
		{name: "number 10", number: 10, expected: "TEN"},
		{name: "number 11", number: 11, expected: "ELEVEN"},
		{name: "number 12", number: 12, expected: "TWELVE"},
		{name: "number 13", number: 13, expected: "THIRTEEN"},
		{name: "number 14", number: 14, expected: "FOURTEEN"},
		{name: "number 15", number: 15, expected: "FIFTEEN"},
		{name: "number 16", number: 16, expected: "SIXTEEN"},
		{name: "number 17", number: 17, expected: "SEVENTEEN"},
		{name: "number 18", number: 18, expected: "EIGHTEEN"},
		{name: "number 19", number: 19, expected: "NINETEEN"},
		{name: "number 20", number: 20, expected: "TWENTY"},
		{name: "number 30", number: 30, expected: "THIRTY"},
		{name: "number 40", number: 40, expected: "FORTY"},
		{name: "number 50", number: 50, expected: "FIFTY"},
		{name: "number 60", number: 60, expected: "SIXTY"},
		{name: "number 70", number: 70, expected: "SEVENTY"},
		{name: "number 80", number: 80, expected: "EIGHTY"},
		{name: "number 90", number: 90, expected: "NINETY"},
		{name: "number 39", number: 39, expected: "THIRTY NINE"},
		{name: "number 48", number: 48, expected: "FORTY EIGHT"},
		{name: "number 100", number: 100, expected: "ONE HUNDERED"},
		{name: "number 123", number: 123, expected: "ONE HUNDERED TWENTY THREE"},
		{name: "number 968", number: 968, expected: "NINE HUNDERED SIXTY EIGHT"},
		{name: "number 850", number: 850, expected: "EIGHT HUNDERED FIFTY"},
		{name: "number 1.000", number: 1000, expected: "ONE THOUSAND"},
		{name: "number 1.624", number: 1624, expected: "ONE THOUSAND SIX HUNDERED TWENTY FOUR"},
		{name: "number 8.596", number: 8596, expected: "EIGHT THOUSAND FIVE HUNDERED NINETY SIX"},
		{name: "number 10.000", number: 10000, expected: "TEN THOUSAND"},
		{name: "number 115.000", number: 115000, expected: "ONE HUNDERED FIFTEEN THOUSAND"},
		{name: "number 100.000", number: 100000, expected: "ONE HUNDERED THOUSAND"},
		{name: "number 850.000", number: 850000, expected: "EIGHT HUNDERED FIFTY THOUSAND"},
		{name: "number 1.000.000", number: 1000000, expected: "ONE MILLION"},
		{name: "number 1.000.000.000", number: 1000000000, expected: "ONE BILLION"},
		{name: "number 1.000.000.000.000", number: 1000000000000, expected: "ONE TRILLION"},
		{name: "number 1.000.000.000.000.000", number: 1000000000000000, expected: "ONE QUADRILION"},
		{name: "number 1.001.001.001.001.001", number: 1001001001001001, expected: "ONE QUADRILION ONE TRILLION ONE BILLION ONE MILLION ONE THOUSAND ONE"},
		{name: "number 9.876.543.210.987.654", number: 9876543210987654, expected: "NINE QUADRILION EIGHT HUNDERED SEVENTY SIX TRILLION FIVE HUNDERED FORTY THREE BILLION TWO HUNDERED TEN MILLION NINE HUNDERED EIGHTY SEVEN THOUSAND SIX HUNDERED FIFTY FOUR"},
	}

	for _, test := range tests {
		t.Run(test.name+" should be "+test.expected, func(t *testing.T) {
			if actual, err := n2w.ToEnglishWords(test.number); err != nil {
				t.Errorf("Failed %s with %v", actual, err)
			} else if test.expected != actual {
				t.Errorf("Failed %s", actual)
			}
		})
	}
}
