package parsers

type IArgsParser interface {
	GetArgs(osArgs []string) map[string]string
	GetPrefixes() []string
}
