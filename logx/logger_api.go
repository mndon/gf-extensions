package logx

// ----------- output  api --------------

func (l *Logger) Print(v ...interface{}) {
	l.Logger.Print(l.Ctx, v...)
}

// Printf prints `v` with format `format` using fmt.Sprintf.
// The parameter `v` can be multiple variables.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Logger.Printf(l.Ctx, format, v...)
}

// Fatal prints the logging content with [FATA] header and newline, then exit the current process.
func (l *Logger) Fatal(v ...interface{}) {
	l.Logger.Fatal(l.Ctx, v...)
}

// Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Logger.Fatalf(l.Ctx, format, v...)
}

// Panic prints the logging content with [PANI] header and newline, then panics.
func (l *Logger) Panic(v ...interface{}) {
	l.Logger.Panic(l.Ctx, v...)
}

// Panicf prints the logging content with [PANI] header, custom format and newline, then panics.
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Logger.Panicf(l.Ctx, format, v...)
}

// Info prints the logging content with [INFO] header and newline.
func (l *Logger) Info(v ...interface{}) {
	l.Logger.Info(l.Ctx, v...)
}

// Infof prints the logging content with [INFO] header, custom format and newline.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Logger.Infof(l.Ctx, format, v...)
}

// Debug prints the logging content with [DEBU] header and newline.
func (l *Logger) Debug(v ...interface{}) {
	l.Logger.Debug(l.Ctx, v...)
}

// Debugf prints the logging content with [DEBU] header, custom format and newline.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Logger.Debugf(l.Ctx, format, v...)
}

// Notice prints the logging content with [NOTI] header and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Notice(v ...interface{}) {
	l.Logger.Notice(l.Ctx, v...)
}

// Noticef prints the logging content with [NOTI] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Noticef(format string, v ...interface{}) {
	l.Logger.Noticef(l.Ctx, format, v...)
}

// Warning prints the logging content with [WARN] header and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Warning(v ...interface{}) {
	l.Logger.Warning(l.Ctx, v...)
}

// Warningf prints the logging content with [WARN] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.Logger.Warningf(l.Ctx, format, v...)
}

// Error prints the logging content with [ERRO] header and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Error(v ...interface{}) {
	l.Logger.Error(l.Ctx, v...)
}

// Errorf prints the logging content with [ERRO] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Logger.Errorf(l.Ctx, format, v...)
}

// Critical prints the logging content with [CRIT] header and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Critical(v ...interface{}) {
	l.Logger.Critical(l.Ctx, v...)
}

// Criticalf prints the logging content with [CRIT] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	l.Logger.Criticalf(l.Ctx, format, v...)
}
