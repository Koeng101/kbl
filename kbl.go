package kbl

import (
	"github.com/koeng101/poly"
	"sync"
)

type Query struct {
	Select       []string        `json:"SELECT"`
	Optimize     poly.CodonTable `json:"OPTIMIZE"`
	SynthesisFix []string        `json:"SYNTHESIS_FIX"`
}

func (q *Query) Compile() []string {
	// Case Optimize
	codonTable := q.Optimize
	for i, sequence := range q.Select {
		q.Select[i] = poly.Optimize(sequence, codonTable)
	}
	// Case SynthesisFix
	if len(q.SynthesisFix) > 0 {
		var functions []func(string, chan poly.DnaSuggestion, *sync.WaitGroup)
		m := make(map[string]func(string, chan poly.DnaSuggestion, *sync.WaitGroup))
		m["TypeIIS"] = poly.FindTypeIIS
		for _, function := range q.SynthesisFix {
			functions = append(functions, m[function])
		}
		for i, sequence := range q.Select {
			q.Select[i], _ = poly.FixCds(":memory:", sequence, codonTable, functions)
		}
	}
	return q.Select
}
