package utils

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// AppRootPath Application Root Path
	AppRootPath = filepath.Join(filepath.Dir(b), "../")
)
