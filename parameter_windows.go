package cons

import "syscall"

const (
	StdInputHandle  = ^uint32(10) + 1
	StdOutputHandle = ^uint32(11) + 1
	StdErrorHandle  = ^uint32(12) + 1
	invalidHandle   = Handle(0)
)

const (
	errorSuccess = syscall.Errno(0)
)

const (
	EnableProcessedInput            = 0x0001
	EnableLineInput                 = 0x0002
	EnableEchoInput                 = 0x0004
	EnableWindowInput               = 0x0008
	EnableMouseInput                = 0x0010
	EnableInsertMode                = 0x0020
	EnableQuickEditMode             = 0x0040
	EnableVirtualTerminalInput      = 0x0200
	EnableProcessedOutput           = 0x0001
	EnableWrapAtEolOutput           = 0x0002
	EnableVirtualTerminalProcessing = 0x0004
	DisableNewlineAutoReturn        = 0x0008
	EnableLvbGridWorldwide          = 0x0010
)

const (
	ForegroundBlue      = 0x0001
	ForegroundGreen     = 0x0002
	ForegroundRed       = 0x0004
	ForegroundIntensity = 0x0008
	BackgroundBlue      = 0x0010
	BackgroundGreen     = 0x0020
	BackgroundRed       = 0x0040
)

const (
	Acp       = 0
	Oemcp     = 1
	Maccp     = 2
	ThreadAcp = 3
	Symbol    = 42
	Utf7      = 65000
	Utf8      = 65001
)

const (
	KeyEvent              = 0x0001
	MouseEvent            = 0x0002
	WindowBufferSizeEvent = 0x0004
	MenuEvent             = 0x0008
	FocusEvent            = 0x0010
)
