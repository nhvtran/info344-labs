package main

import "testing"

func TestcountSwears(t *testing.T) {
	cases := []struct {
		name        string
		listOfWords []string
	}{
	// {
	// 	name:        "non-swear words",
	// 	listOfWords: [3]string{"lol", "rip", "roast"},
	// },
	}

	knownSwears := loadKnownSwears("known_swears.txt")
	for _, c := range cases {
		output := countSwears(knownSwears, c.listOfWords)
		for _, word := range c.listOfWords {
			if output[word] > 0 {
				t.Errorf("got %v, but expected", word)
			}
		}
	}
}

func TestloadKnownSwears(t *testing.T) {

}

func TestloadWords(t *testing.T) {

}

func TestopenFile(t *testing.T) {
	// hint os.Args = []string{"cmd", "path/to/file.txt"}
}
