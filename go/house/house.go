// Package house is fun.
package house

const testVersion = 1

type Phrase struct {
	Beginning string
	Ending    string
}

var phrases = []Phrase{
	{},
	{Beginning: "that lay in ", Ending: "the house that Jack built."},
	{Beginning: "that ate ", Ending: "the malt\n"},
	{Beginning: "that killed ", Ending: "the rat\n"},
	{Beginning: "that worried ", Ending: "the cat\n"},
	{Beginning: "that tossed ", Ending: "the dog\n"},
	{Beginning: "that milked ", Ending: "the cow with the crumpled horn\n"},
	{Beginning: "that kissed ", Ending: "the maiden all forlorn\n"},
	{Beginning: "that married ", Ending: "the man all tattered and torn\n"},
	{Beginning: "that woke ", Ending: "the priest all shaven and shorn\n"},
	{Beginning: "that kept ", Ending: "the rooster that crowed in the morn\n"},
	{Beginning: "that belonged to ", Ending: "the farmer sowing his corn\n"},
	{Beginning: "", Ending: "the horse and the hound and the horn\n"},
}


func Song() string {
	var song string
	for i := 1; i <= 12; i++ {
		song += Verse(i)
		if i != 12 {
			song += "\n\n"
		}
	}
	return song
}


func Verse(i int) string {
	verse := "This is " + phrases[i].Ending + build(i-1)
	return verse
}


func build(i int) string {
	if i == 0 {
		return ""
	}
	return phrases[i].Beginning + phrases[i].Ending + build(i-1)
}
