package logx

import "context"

// ----------- output  api --------------

func (l *Logger) Print(ctx context.Context, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Print(ctx, v...)
}

// Printf prints `v` with format `format` using fmt.Sprintf.
// The parameter `v` can be multiple variables.
func (l *Logger) Printf(ctx context.Context, format string, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Printf(ctx, format, v...)
}

// Fatal prints the logging content with [FATA] header and newline, then exit the current process.
func (l *Logger) Fatal(ctx context.Context, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Fatal(ctx, v...)
}

// Fatalf prints the logging content with [FATA] header, custom format and newline, then exit the current process.
func (l *Logger) Fatalf(ctx context.Context, format string, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Fatalf(ctx, format, v...)
}

// Panic prints the logging content with [PANI] header and newline, then panics.
func (l *Logger) Panic(ctx context.Context, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Panic(ctx, v...)
}

// Panicf prints the logging content with [PANI] header, custom format and newline, then panics.
func (l *Logger) Panicf(ctx context.Context, format string, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Panicf(ctx, format, v...)
}

// Info prints the logging content with [INFO] header and newline.
func (l *Logger) Info(ctx context.Context, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Info(ctx, v...)
}

// Infof prints the logging content with [INFO] header, custom format and newline.
func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Infof(ctx, format, v...)
}

// Debug prints the logging content with [DEBU] header and newline.
func (l *Logger) Debug(ctx context.Context, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Debug(ctx, v...)
}

// Debugf prints the logging content with [DEBU] header, custom format and newline.
func (l *Logger) Debugf(ctx context.Context, format string, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Debugf(ctx, format, v...)
}

// Notice prints the logging content with [NOTI] header and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Notice(ctx context.Context, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Notice(ctx, v...)
}

// Noticef prints the logging content with [NOTI] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Noticef(ctx context.Context, format string, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Noticef(ctx, format, v...)
}

// Warning prints the logging content with [WARN] header and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Warning(ctx context.Context, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Warning(ctx, v...)
}

// Warningf prints the logging content with [WARN] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Warningf(ctx context.Context, format string, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Warningf(ctx, format, v...)
}

// Error prints the logging content with [ERRO] header and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Error(ctx context.Context, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Error(ctx, v...)
}

// Errorf prints the logging content with [ERRO] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Errorf(ctx, format, v...)
}

// Critical prints the logging content with [CRIT] header and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Critical(ctx context.Context, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Critical(ctx, v...)
}

// Criticalf prints the logging content with [CRIT] header, custom format and newline.
// It also prints caller stack info if stack feature is enabled.
func (l *Logger) Criticalf(ctx context.Context, format string, v ...interface{}) {
	ctx = context.WithValue(ctx, CustomFieldsKey, l.customFields)
	l.Logger.Criticalf(ctx, format, v...)
}
