package foodchain

import "strings"

const TestVersion = 1

var verse = []struct{ thing, reason, exclaim string }{
	{"", "", ""},
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die.", ""},
	{"spider", "She swallowed the spider to catch the fly.", "It wriggled and jiggled and tickled inside her."},
	{"bird", "She swallowed the bird to catch the spider that wriggled and jiggled and tickled inside her.", "How absurd to swallow a bird!"},
	{"cat", "She swallowed the cat to catch the bird.", "Imagine that, to swallow a cat!"},
	{"dog", "She swallowed the dog to catch the cat.", "What a hog, to swallow a dog!"},
	{"goat", "She swallowed the goat to catch the dog.", "Just opened her throat and swallowed a goat!"},
	{"cow", "She swallowed the cow to catch the goat.", "I don't know how she swallowed a cow!"},
	{"horse", "", "She's dead, of course!"},
}

func Verse(v int) (s string) {
	if v > len(verse) || v < 1 {
		return ""
	}
	s = "I know an old lady who swallowed a " + verse[v].thing + "."

	if verse[v].exclaim != "" {
		s += "\n" + verse[v].exclaim
	}

	if verse[v].reason != "" {
		for ; v > 0; v-- {
			s += "\n" + verse[v].reason
		}
	}

	return
}

func Verses(f, l int) string {
	verses := make([]string, 0, l-f+1)

	for i := f; i <= l; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}

func Song() (s string) {
	return Verses(1, 8)
}