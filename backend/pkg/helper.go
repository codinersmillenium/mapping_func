package pkg

import (
	"fmt"
	"strings"
)

func ProcessCity(cityStr string, capitals map[string]string) string {
	city := strings.TrimSpace(strings.ToUpper(cityStr))
	if prov, ok := capitals[city]; ok {
		return fmt.Sprintf("%-20.20s", city+" "+strings.ToUpper(prov))
	}
	return fmt.Sprintf("%-20.20s", city)
}

func IsAgeSuffix(s string) bool {
	s = strings.ToUpper(strings.TrimSpace(s))
	return s == "TAHUN" || s == "THN" || s == "TH" ||
		strings.HasSuffix(s, "TAHUN") ||
		strings.HasSuffix(s, "THN") ||
		strings.HasSuffix(s, "TH")
}
