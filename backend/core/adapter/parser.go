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
	// 1. i (index), 2. tmp (buffer), 3. n (name), 4. a (age), 5. c (city)
	var i int = len(input) - 1
	var tmp string
	var n, a, c string

	// Step 1: get CITY
	for ; i >= 0; i-- {
		if input[i] >= '0' && input[i] <= '9' {
			c = strings.TrimSpace(strings.ToUpper(tmp))
			// Cek Capital
			if prov, ok := u.Capitals[c]; ok {
				c = c + " " + strings.ToUpper(prov)
			}
			tmp = ""
			break
		}
		tmp = string(input[i]) + tmp
	}

	for ; i >= 0; i-- {
		if input[i] == ' ' && a != "" {
			break
		}
		if input[i] >= '0' && input[i] <= '9' {
			a = string(input[i]) + a
		}
	}

	// FORMATTING FIXED-WIDTH (Name: 30, Age: 3, City: 20)
	return fmt.Sprintf("%-30.30s", n),
		fmt.Sprintf("%-3.3s", a),
		fmt.Sprintf("%-20.20s", c)
}
