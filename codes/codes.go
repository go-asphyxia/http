package codes

const (
	Continue          int = 100
	SwitchingProtocol int = 101
	Processing        int = 102
	EarlyHints        int = 103

	OK                          int = 200
	Created                     int = 201
	Accepted                    int = 202
	NonAuthoritativeInformation int = 203
	NoContent                   int = 204
	ResetContent                int = 205
	PartialContent              int = 206

	MultipleChoice    int = 300
	MovedPermanently  int = 301
	Found             int = 302
	SeeOther          int = 303
	NotModified       int = 304
	UseProxy          int = 305
	SwitchProxy       int = 306
	TemporaryRedirect int = 307
	PermanentRedirect int = 308

	BadRequest                   int = 400
	Unauthorized                 int = 401
	PaymentRequired              int = 402
	Forbidden                    int = 403
	NotFound                     int = 404
	MethodNotAllowed             int = 405
	NotAcceptable                int = 406
	ProxyAuthenticationRequired  int = 407
	RequestTimeout               int = 408
	Conflict                     int = 409
	Gone                         int = 410
	LengthRequired               int = 411
	PreconditionFailed           int = 412
	RequestEntityTooLarge        int = 413
	RequestURITooLong            int = 414
	UnsupportedMediaType         int = 415
	RequestedRangeNotSatisfiable int = 416
	ExpectationFailed            int = 417

	InternalServerError     int = 500
	NotImplemented          int = 501
	BadGateway              int = 502
	ServiceUnavailable      int = 503
	GatewayTimeout          int = 504
	HTTPVersionNotSupported int = 505
)
