package logapi

type Layout interface {
	SetCharset(cs string)

	// SetLayoutProps check the layout props is supported or not
	SetLayoutProps(props map[string]string) error

	// SetPattern use pattern format of log4j2
	SetPattern(pattern []byte)
}
