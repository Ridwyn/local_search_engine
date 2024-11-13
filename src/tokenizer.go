package src

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Term string
type TermCount int
type TermCountMap map[Term]TermCount

type Tokenizer struct {
	Filepath       string
	Content        []rune
	TermCountMap   TermCountMap
	TotalTermCount int
}

func NewTokenizer(c string, fp string) *Tokenizer {
	content := strings.TrimSpace(c)
	var rnes []rune = make([]rune, 0)
	for _, runeValue := range content {

		rnes = append(rnes, runeValue)

	}

	tokenizer := &Tokenizer{Content: rnes, TermCountMap: TermCountMap{}, TotalTermCount: 0, Filepath: fp}
	tokenizer.TokeniseContent()

	return tokenizer

}

func NewTokenizerQuery(c string) Tokenizer {
	content := strings.TrimSpace(c)
	var rnes []rune = make([]rune, 0)
	for _, runeValue := range content {
			rnes = append(rnes, runeValue)
	}

	tokenizer := Tokenizer{Content: rnes, TermCountMap: TermCountMap{}, TotalTermCount: 0}
	tokenizer.TokeniseContent()
	return tokenizer

}

func (t *Tokenizer) TokeniseContent() *Tokenizer {
	for _,runeValue := range t.Content {

		// rv := t.Content[0]
		if !utf8.ValidRune(runeValue) {
			fmt.Printf("Error invalid utf8 char %v", t.Content[0])
			continue
		}
		if unicode.IsDigit(runeValue) {
			t.chopWhileNumeric()
		}

		if unicode.IsLetter(runeValue) {
			t.chopWhileAlpabetic()
		}
		
		t.chop(t.Content, 1)

	}

	return t
}

func (t *Tokenizer) chop(c []rune, size int) {
	if len(c) <= 0 {
		return
	}
	runesTk := c[0:size]
	t.Content = t.Content[size:]

	str := string(runesTk)
	token := strings.ToLower(strings.TrimSpace(str))

	if ((token!="") || (len(token)!=0) ) {
		
		if _, ok := t.TermCountMap[Term(token)]; ok {
			t.TermCountMap[(Term(token))] = t.TermCountMap[Term(token)] + 1
			t.TotalTermCount += 1
		} else {
	
			t.TermCountMap[Term(token)] = 1
			t.TotalTermCount += 1
		}
	}


}
func (t *Tokenizer) chopWhileNumeric() {
	n := 0
	for n < len(t.Content) && unicode.IsDigit(rune(t.Content[n])) {
		n++
	}

	t.chop(t.Content, n)
}

func (t *Tokenizer) chopWhileAlpabetic() {
	n := 0
	for n < len(t.Content) && (unicode.IsLetter(rune(t.Content[n])) || unicode.IsDigit(rune(t.Content[n]))) {
		n++
	}
	t.chop(t.Content, n)
}
