package logx

import (
	"bytes"
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

func HandlerPlain(ctx context.Context, in *glog.HandlerInput) {
	in.Buffer = formatPlainOutput(ctx, *in)
	in.Next(ctx)
}

func formatPlainOutput(ctx context.Context, in glog.HandlerInput) *bytes.Buffer {
	buffer := bytes.NewBuffer(nil)
	if in.Logger.GetConfig().HeaderPrint {
		if in.TimeFormat != "" {
			buffer.WriteString(in.TimeFormat)
		}
		if in.Logger.GetConfig().LevelPrint && in.LevelFormat != "" {
			var levelStr = "[" + in.LevelFormat + "]"
			if in.Color > 0 {
				levelStr = getColoredStr(getColorByLevel(in.Level), levelStr)
			}
			addStringToBuffer(buffer, levelStr)
		}
	}
	if in.TraceId != "" {
		addStringToBuffer(buffer, "{"+in.TraceId+"}")
	}

	// 自定义字段处理
	v, ok := ctx.Value(CustomFieldsKey).(*CustomFields)
	if ok {
		addStringToBuffer(buffer, gconv.String(v))
	}

	if in.Logger.GetConfig().HeaderPrint {
		if in.Prefix != "" {
			addStringToBuffer(buffer, in.Prefix)
		}
		if in.CallerFunc != "" {
			addStringToBuffer(buffer, in.CallerFunc)
		}
		if in.CallerPath != "" {
			addStringToBuffer(buffer, in.CallerPath)
		}
	}

	if in.Content != "" {
		addStringToBuffer(buffer, in.Content)
	}

	// Convert values string content.
	var valueContent string
	for _, v := range in.Values {
		valueContent = gconv.String(v)
		if len(valueContent) == 0 {
			continue
		}
		if buffer.Len() > 0 {
			if buffer.Bytes()[buffer.Len()-1] == '\n' {
				// Remove one blank line(\n\n).
				if valueContent[0] == '\n' {
					valueContent = valueContent[1:]
				}
				buffer.WriteString(valueContent)
			} else {
				buffer.WriteString(" " + valueContent)
			}
		} else {
			buffer.WriteString(valueContent)
		}
	}

	if in.Stack != "" {
		addStringToBuffer(buffer, "\nStack:\n"+in.Stack)
	}
	// avoid a single space at the end of a line.
	buffer.WriteString("\n")
	return buffer
}

func addStringToBuffer(buffer *bytes.Buffer, strings ...string) {
	for _, s := range strings {
		if buffer.Len() > 0 {
			buffer.WriteByte(' ')
		}
		buffer.WriteString(s)
	}
}
