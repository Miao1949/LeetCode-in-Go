package Problem0754

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n, k int
	ans  string
}{

	{
		1,
		2,
		"10",
	},

	{
		2,
		2,
		"01100",
	},

	{
		4,
		5,
		"0004444344424441444044334432443144304423442244214420441344124411441044034402440144004343424341434043334332433143304323432243214320431343124311431043034302430143004242414240423342324231423042234222422142204213421242114210420342024201420041414041334132413141304123412241214120411341124111411041034102410141004040334032403140304023402240214020401340124011401040034002400140003333233313330332233213320331233113310330233013300323231323032223221322032123211321032023201320031313031223121312031123111311031023101310030302230213020301230113010300230013000222212220221122102201220021212021112110210121002020112010200120001111011001010000",
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, crackSafe(tc.n, tc.k), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			crackSafe(tc.n, tc.k)
		}
	}
}
