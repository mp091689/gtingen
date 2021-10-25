package parsers

import (
	"strings"
)

type CArgsParser struct{}

func (p CArgsParser) GetArgs(osArgs []string) map[string]string {
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

func (p CArgsParser) GetPrefixes() []string {
	return []string{"-source=", "-out="}
}
