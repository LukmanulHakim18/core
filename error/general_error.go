package error

import "net/http"

// ========================= general Error Data =========================
// general have prefix error code "BB-0***"

var UnknownErrorGateway = &Error{
	StatusCode:   http.StatusInternalServerError,
	ErrorCode:    "GW-9999",
	ErrorMessage: "Gateway Error",
	LocalizedMessage: Message{
		English:   "Gateway Error",
		Indonesia: "Masalah pada gateway",
	},
}

var UnknownError = &Error{
	StatusCode:   http.StatusInternalServerError,
	ErrorCode:    "BB-0001",
	ErrorMessage: "Unknown Error",
	LocalizedMessage: Message{
		English:   "Unknown Error",
		Indonesia: "Masalah tidak diketahui penyebabnya",
	},
}

var ErrorParse = &Error{
	StatusCode:   http.StatusBadRequest,
	ErrorCode:    "BB-0002",
	ErrorMessage: "Oops! Something went wrong!",
	LocalizedMessage: Message{
		English:   "Oops! Something went wrong!",
		Indonesia: "Terjadi kesalahan di sistem",
	},
}

var UnsupportedAppVersion = &Error{
	StatusCode:   http.StatusBadRequest,
	ErrorCode:    "BB-0003",
	ErrorMessage: "Cannot process request, please make sure you are using our latest version",
	LocalizedMessage: Message{
		English:   "Cannot process request, please make sure you are using our latest version",
		Indonesia: "Tidak dapat memperoses permintaan, mohon pastikan anda menggunakan aplikasi versi terbaru",
	},
}
var ErrorDatabase = &Error{
	StatusCode:   http.StatusInternalServerError,
	ErrorCode:    "BB-0004",
	ErrorMessage: "Sorry, We are unable to complete your request. Please try again.",
	LocalizedMessage: Message{
		English:   "Sorry, We are unable to complete your request. Please try again.",
		Indonesia: "Maaf, Kami tidak dapat memproses permintaan Anda. Mohon coba kembali",
	},
}

var MissingRequiredParam = &Error{
	StatusCode:   http.StatusBadRequest,
	ErrorCode:    "BB-0005",
	ErrorMessage: "%v is required",
	LocalizedMessage: Message{
		English:   "%v is required",
		Indonesia: "%v dibutuhkan",
	},
}

var InvalidParam = &Error{
	StatusCode:   http.StatusBadRequest,
	ErrorCode:    "BB-0006",
	ErrorMessage: "%v is invalid",
	LocalizedMessage: Message{
		English:   "%v is invalid",
		Indonesia: "%v tidak sesuai",
	},
}

var UnimplementedMethod = &Error{
	StatusCode:   http.StatusBadRequest,
	ErrorCode:    "BB-0007",
	ErrorMessage: "%v unimplemented",
	LocalizedMessage: Message{
		English:   "%v unimplemented",
		Indonesia: "%v tidak diimplementasikan",
	},
}
var UnknownMiddleware = &Error{
	StatusCode:   http.StatusInternalServerError,
	ErrorCode:    "BB-0008",
	ErrorMessage: "%v middleware not found",
	LocalizedMessage: Message{
		English:   "%v middleware not found",
		Indonesia: "middleware %v tidak ditemukan",
	},
}
