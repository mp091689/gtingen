package parsers

import (
	"strings"
)

type VArgsParser struct{}

func (p VArgsParser) GetArgs(osArgs []string) map[string]string {
	args := make(map[string]string)

	prfixes := p.GetPrefixes()

	for _, a := range osArgs {
		for _, pref := range prfixes {
			if strings.HasPrefix(a, pref) {
				args[pref] = strings.TrimPrefix(a, pref)
			}
		}
	}

	return args
}

func (p VArgsParser) GetPrefixes() []string {
	return []string{"-source=", "-out="}
}
