package src

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"sort"
)

type DocFrequency map[Term]float32

type Doc struct {
	Name     string
	TermFreq TermCountMap
	DocFreq  DocFrequency
	Count    int
}
type Docs []Doc

type Model struct {
	Docs []Doc
}

// var documents Docs = make(Docs, 4)

func NewModel() *Model {

	return &Model{Docs: make(Docs, 4)}

}

func (m *Model) saveDocument(d Doc) {

	m.Docs = append(m.Docs, d)
}

func (m *Model) SaveAllDocuments() {
	b, err := json.Marshal(m.Docs)
	if err != nil {
		fmt.Printf("Error: couldn't Json Serialize doc: %v\n", err)
	}
	dirPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: could get working directory: %v\n", err)
	}
	// write to file
	os.WriteFile(dirPath+"/index.json", b, os.ModePerm)

}

func (m *Model) Query_terms(t Tokenizer) {
	res := make(map[string]float32)
	for token, _ := range t.TermCountMap {
		for i := 0; i < len(m.Docs); i++ {
			if val, ok := m.Docs[i].DocFreq[token]; ok {
				res[m.Docs[i].Name] = val
			}

		}

	}
	//sort result by val
	keys := make([]string, 0, len(res))

	for key := range res {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return res[keys[i]] > res[keys[j]]
	})

	//Print only top to
	fmt.Println("Getting Results of Query from top 10")
	for _, k := range keys[0:10] {
		fmt.Println(k, res[k])
	}

}

func (m *Model) NewDoc(t *Tokenizer) Doc {
	d := Doc{TermFreq: t.TermCountMap, DocFreq: make(DocFrequency), Count: t.TotalTermCount, Name: t.Filepath}

	for term, tc := range t.TermCountMap {

		// d.TermFreq[term] = computeTf(tc, d) *computeIdf(term)

		d.DocFreq[term] = m.computeTf(tc, d, term) * m.computeIdf(term)

	}

	m.saveDocument(d)

	return d
}

func (m *Model) computeTf(tc TermCount, doc Doc, term Term) float32 {
	// return (float32(tc) / float32(doc.Count))
	fmt.Printf("%v tf (%v) termcount:%v , totaltermsInDoc:%v === %v \n",doc.Name,term,tc,doc.Count, (float32(tc)/float32(doc.Count)))
	return (float32(tc)/float32(doc.Count))
}

func (m *Model) computeIdf(term Term) float32 {
	n := len(m.Docs)
	sum := 1
	for i := 0; i < n; i++ {
		if _, ok := m.Docs[i].TermFreq[term]; ok {
			sum++
		}
	}
	fmt.Printf("%v idf docfreq:%v ===%v \n",term,sum, float32(math.Log10(float64(n)+1 / (float64(sum))))+1)
	return float32(math.Log(float64(n) / (float64(sum)+1)))

}
