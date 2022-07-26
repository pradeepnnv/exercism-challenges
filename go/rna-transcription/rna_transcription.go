package strand

import "strings"

func ToRNA(dna string) string {
	dnaToRnaMap := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
	var rna strings.Builder

	for _, r := range dna {
		rna.WriteRune(dnaToRnaMap[r])
	}
	return rna.String()
}
