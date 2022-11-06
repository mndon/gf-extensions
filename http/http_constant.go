package http

const (
	HttpHeaderRemoteIp = "X-Real-Ip"
	HttpHeaderUA       = "User-Agent"
)

// jwt相关配置
const (
	JwtKey         = "jwt.key"
	JwtTimeout     = "jwt.timeout"
	JwtMaxRefresh  = "jwt.max_refresh"
	JwtIdentityKey = "jwt.identity_key"
)
