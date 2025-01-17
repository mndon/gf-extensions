package logx

func (l *Logger) Type(t string) *Logger {
	l.customFields.Type = t
	return l
}
