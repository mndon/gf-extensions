package logx

import "time"

func (l *Logger) Type(t string) *Logger {
	l.customFields.Type = t
	return l
}

func (l *Logger) Uid(t string) *Logger {
	l.customFields.Uid = t
	return l
}

func (l *Logger) AccessTime(t time.Duration) *Logger {
	l.customFields.AccessTime = t
	return l
}

func (l *Logger) ResStatus(t int) *Logger {
	l.customFields.ResStatus = t
	return l
}

func (l *Logger) ReqMethod(t string) *Logger {
	l.customFields.ReqMethod = t
	return l
}

func (l *Logger) ReqUri(t string) *Logger {
	l.customFields.ReqUri = t
	return l
}

func (l *Logger) ReqUrl(t string) *Logger {
	l.customFields.ReqUrl = t
	return l
}

func (l *Logger) ReqBody(t string) *Logger {
	l.customFields.ReqBody = t
	return l
}

func (l *Logger) ReqIp(t string) *Logger {
	l.customFields.ReqIp = t
	return l
}

func (l *Logger) UA(t string) *Logger {
	l.customFields.UA = t
	return l
}
