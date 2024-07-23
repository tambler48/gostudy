package wordsCounter

import (
	"reflect"
	"testing"
)

func TestWordsCounter(t *testing.T) {
	cases := []struct {
		input string
		want  map[string]int
	}{
		{
			"One two two three three Three four four Four four",
			map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
				"four":  4,
			},
		},
		{
			"Harry Potter is a series of seven fantasy novels written by British author J. K. Rowling. The novels chronicle the lives of a young wizard, Harry Potter, and his friends, Hermione Granger and Ron Weasley, all of whom are students at Hogwarts School of Witchcraft and Wizardry. The main story arc concerns Harry's conflict with Lord Voldemort, a dark wizard who intends to become immortal, overthrow the wizard governing body known as the Ministry of Magic, and subjugate all wizards and Muggles (non-magical people).\n\nThe series was originally published in English by Bloomsbury in the United Kingdom and Scholastic Press in the United States. A series of many genres, including fantasy, drama, coming-of-age fiction, and the British school story (which includes elements of mystery, thriller, adventure, horror, and romance), the world of Harry Potter explores numerous themes and includes many cultural meanings and references.[1] Major themes in the series include prejudice, corruption, madness, and death.[2][3]\n\nSince the release of the first novel, Harry Potter and the Philosopher's Stone, on 26 June 1997, the books have found immense popularity, positive reviews, and commercial success worldwide. They have attracted a wide adult audience as well as younger readers and are widely considered cornerstones of modern literature.[4][5] As of February 2023, the books have sold more than 600 million copies worldwide, making them the best-selling book series in history, and have been available in 85 languages.[6] The last four books consecutively set records as the fastest-selling books in history, with the final instalment selling roughly 2.7 million copies in the United Kingdom and 8.3 million copies in the United States within twenty-four hours of its release.\n\nWarner Bros. Pictures adapted the original seven books into an eight-part namesake film series. In 2016, the total value of the Harry Potter franchise was estimated at $25 billion,[7] making it one of the highest-grossing media franchises of all time. Harry Potter and the Cursed Child is a play based on a story co-written by Rowling.",
			map[string]int{
				"the":    27,
				"and":    17,
				"of":     16,
				"in":     10,
				"harry":  7,
				"a":      7,
				"potter": 6,
				"series": 6,
				"books":  5,
				"as":     5,
			},
		},
	}

	for _, c := range cases {
		got := WordsCounter(c.input)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("input:%v, got: %v, want: %v", c.input, got, c.want)
		}
	}
}
