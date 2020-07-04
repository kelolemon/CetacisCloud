package src

type (
	account struct {
		User string `json:"user"`
		Pass string `json:"pass"`
	}
)

type (
	flag struct {
		Success string `json:"success"`
	}
)

type (
	FileInfo struct {
		FileName string `json:"file_name"`
		FileSize int64 `json:"file_size"`
		FileType bool `json:"file_type"`
		IsShare bool `json:"is_share"`
	}
)

type (
	FileList struct {
		List []FileInfo `json:"list"`
	}
)

type (
	LinkInfo struct {
		Link string `json:"link"`
	}
)