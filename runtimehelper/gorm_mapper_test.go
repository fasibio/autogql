package runtimehelper_test

import (
	"testing"

	"github.com/fasibio/autogql/runtimehelper"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
)

func TestCombineSimpleQuery(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		elements []runtimehelper.ConditionElement
		relation runtimehelper.Relation
		want     string
		// want2    []interface{}
	}{
		{
			name: "simple one",
			elements: []runtimehelper.ConditionElement{
				runtimehelper.Equal("A", "eqvalue"),
			},
			relation: runtimehelper.RelationAnd,
			want:     "A = ?",
		},
		{
			name: "simple two values",
			elements: []runtimehelper.ConditionElement{
				runtimehelper.Equal("A", "eqvalue"),
				runtimehelper.Equal("B", "eqbvalue"),
			},
			relation: runtimehelper.RelationAnd,
			want:     "A = ? AND B = ?",
		},
		{
			name: "simple two values child",
			elements: []runtimehelper.ConditionElement{

				runtimehelper.Equal("A", "eqvalue"),
				runtimehelper.Complex(runtimehelper.RelationOr, runtimehelper.Equal("B", "eqbvalue"), runtimehelper.Equal("B", "eqotherbvalue")),
			},
			relation: runtimehelper.RelationAnd,
			want:     "A = ? AND (B = ? OR B = ?)",
		},
		{
			name: "simple two values child secend not ",
			elements: []runtimehelper.ConditionElement{

				runtimehelper.Equal("A", "eqvalue"),
				runtimehelper.Complex(runtimehelper.RelationOr, runtimehelper.Equal("B", "eqbvalue"), runtimehelper.NotEqual("B", "eqotherbvalue")),
			},
			relation: runtimehelper.RelationAnd,
			want:     "A = ? AND (B = ? OR B <> ?)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := runtimehelper.CombineSimpleQuery(tt.elements, tt.relation)
			assert.Equal(t, tt.want, got)
			snaps.MatchSnapshot(t, got2)
		})
	}
}
