package logapi

// Layout layout is api for display layout format in output
type Layout interface {
	SetCharset(cs string)

	// SetLayoutProps check the layout props is supported or not
	SetLayoutProps(props map[string]string) error

	// SetPattern use pattern format of log4j2
	SetPattern(pattern []byte)
}
