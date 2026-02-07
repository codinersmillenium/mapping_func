package adapter

import (
	"fmt"
	"strings"
)

type ParserAdapter struct {
	Capitals map[string]string
}

func NewParserAdapter(capitals map[string]string) *ParserAdapter {
	return &ParserAdapter{Capitals: capitals}
}

func (u *ParserAdapter) Parse(input string) (string, string, string) {
	var i int = len(input) - 1
	var tmp string
	var n, a, c string

	// Phase 1: Parse city
	for ; i >= 0; i-- {
		char := input[i]

		if char >= '0' && char <= '9' {
			suffix := strings.ToUpper(tmp)
			isAgeSuffix := false

			if suffix == "" ||
				strings.HasSuffix(suffix, "TAHUN") ||
				strings.HasSuffix(suffix, "THN") ||
				strings.HasSuffix(suffix, "TH") {
				isAgeSuffix = true

				if strings.HasSuffix(suffix, "TAHUN") {
					tmp = tmp[:len(tmp)-5]
				} else if strings.HasSuffix(suffix, "THN") {
					tmp = tmp[:len(tmp)-3]
				} else if strings.HasSuffix(suffix, "TH") {
					tmp = tmp[:len(tmp)-2]
				}
			}

			if !isAgeSuffix {
				c = strings.ToUpper(tmp)
			}
			tmp = ""
			a = string(char) + a
			continue
		}

		if a != "" && char == ' ' {
			hasMoreDigits := false
			for j := i - 1; j >= 0; j-- {
				if input[j] >= '0' && input[j] <= '9' {
					hasMoreDigits = true
					break
				} else if input[j] != ' ' {
					break
				}
			}

			if !hasMoreDigits {
				break
			}
		}

		tmp = string(char) + tmp
		if a != "" && (char < '0' || char > '9') {
			tmp = string(char) + tmp
			break
		}
	}

	// Phase 2: Parse name
	if i >= 0 {
		for j := i; j >= 0; j-- {
			tmp = string(input[j]) + tmp
		}
		n = strings.ToUpper(tmp)
	}

	// Validate and format output
	cleanAge := ""
	for _, ch := range a {
		if ch >= '0' && ch <= '9' {
			cleanAge += string(ch)
		}
	}
	a = cleanAge

	// clean last city
	replaceWords := []string{"TAHUN", "TH", "THN"}
	// replace
	for _, w := range replaceWords {
		c = strings.ReplaceAll(c, w, " ")
	}
	key := strings.TrimSpace(c)
	if prov, ok := u.Capitals[key]; ok {
		c = c + " " + strings.ToUpper(prov)
	}
	// Format fixed-width
	return fmt.Sprintf("%-30.30s", n),
		fmt.Sprintf("%-3.3s", a),
		fmt.Sprintf("%-20.20s", c)
}
