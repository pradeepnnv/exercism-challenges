package protein

import (
	"errors"
	"fmt"
)

var ErrStop error = errors.New("this is a STOP codon")
var ErrInvalidBase error = fmt.Errorf("this is an invalid codon")

var codonProteinMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UUG": "Leucine",
}

func FromRNA(rna string) ([]string, error) {
	var protiens []string
	for i := 0; i < len(rna)-2; i += 3 {
		codon := rna[i : i+3]
		v, err := FromCodon(codon)
		if err != nil {
			if err == ErrStop {
				break
			}
			return []string{}, err
		}
		protiens = append(protiens, v)
	}

	return protiens, nil
}

func FromCodon(codon string) (string, error) {

	if codon == "UAA" || codon == "UAG" || codon == "UGA" {
		return "", ErrStop
	}
	v, ok := codonProteinMap[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	return v, nil
}
