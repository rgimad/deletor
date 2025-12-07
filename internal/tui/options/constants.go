package options

// Option names as constants to avoid string literals
const (
	//options for clean view
	ShowHiddenFiles       = "Show hidden files"
	ConfirmDeletion       = "Confirm deletion"
	IncludeSubfolders     = "Include subfolders"
	DeleteEmptySubfolders = "Delete empty subfolders"
	SendFilesToTrash      = "Send files to trash"
	SecureDeleteFiles     = "Secure delete files"
	LogOperations         = "Log operations"
	LogToFile             = "Log to file"
	ShowStatistics        = "Show statistics"
	ExitAfterDeletion     = "Exit after deletion"
	//options for cache view
	SystemCache = "System cache"
)

const (
	SecureDeletionZero    = "zero"
	SecureDeletionDoD     = "dod"
	SecureDeletionGutmann = "gutmann"
)

// If you change the bool in these options, you must also change the values in the default rules json (rules/manager.go).
var DefaultCleanOptionState = map[string]bool{
	ShowHiddenFiles:       false,
	ConfirmDeletion:       false,
	IncludeSubfolders:     false,
	DeleteEmptySubfolders: false,
	SendFilesToTrash:      false,
	SecureDeleteFiles:     false,
	LogOperations:         false,
	LogToFile:             false,
	ShowStatistics:        true,
	ExitAfterDeletion:     false,
}

var DefaultSecureDeletionAlgo = SecureDeletionZero

var DefaultCleanOption = []string{
	ShowHiddenFiles,
	ConfirmDeletion,
	IncludeSubfolders,
	DeleteEmptySubfolders,
	SendFilesToTrash,
	SecureDeleteFiles,
	LogOperations,
	LogToFile,
	ShowStatistics,
	ExitAfterDeletion,
}

var SecureDeletionAlgos = []string{
	SecureDeletionZero,
	SecureDeletionDoD,
	SecureDeletionGutmann,
}

var DefaultCacheOptionState = map[string]bool{
	SystemCache: true,
}

var DefaultCacheOption = []string{
	SystemCache,
}
