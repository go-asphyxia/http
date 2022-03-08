package codes

const (
	//Information
	Continue          = 100
	SwitchingProtocol = 101
	Processing        = 102
	EarlyHints        = 103

	//Successful
	OK                          = 200
	Created                     = 201
	Accepted                    = 202
	NonAuthoritativeInformation = 203
	NoContent                   = 204
	ResetContent                = 205
	PartialContent              = 206

	//Massages about redirect
	MultipleChoice    = 300
	MovedPermanently  = 301
	Found             = 302
	SeeOther          = 303
	NotModified       = 304
	UseProxy          = 305
	SwitchProxy       = 306
	TemporaryRedirect = 307
	PermanentRedirect = 308

	//Client
	BadRequest                   = 400
	Unauthorized                 = 401
	PaymentRequired              = 402
	Forbidden                    = 403
	NotFound                     = 404
	MethodNotAllowed             = 405
	NotAcceptable                = 406
	ProxyAuthenticationRequired  = 407
	RequestTimeout               = 408
	Conflict                     = 409
	Gone                         = 410
	LengthRequired               = 411
	PreconditionFailed           = 412
	RequestEntityTooLarge        = 413
	RequestURITooLong            = 414
	UnsupportedMediaType         = 415
	RequestedRangeNotSatisfiable = 416
	ExpectationFailed            = 417

	//Server
	InternalServerError     = 500
	NotImplemented          = 501
	BadGateway              = 502
	ServiceUnavailable      = 503
	GatewayTimeout          = 504
	HTTPVersionNotSupported = 505
)
