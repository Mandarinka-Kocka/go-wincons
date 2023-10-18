package cons

import "syscall"

var (
	kernel32                        = syscall.NewLazyDLL("kernel32.dll")
	procGetStdHandle                = kernel32.NewProc("GetStdHandle")
	procGetConsoleMode              = kernel32.NewProc("GetConsoleMode")
	procGetConsoleCursorInfo        = kernel32.NewProc("GetConsoleCursorInfo")
	procGetConsoleCP                = kernel32.NewProc("GetConsoleCP")
	procGetConsoleScrrenBufferInfo  = kernel32.NewProc("GetConsoleScreenBufferInfo")
	procGetConsoleOutputCP          = kernel32.NewProc("GetConsoleOutputCP")
	procReadConsoleInput            = kernel32.NewProc("ReadConsoleInputW")
	procSetConsoleMode              = kernel32.NewProc("SetConsoleMode")
	procSetConsoleOutputCP          = kernel32.NewProc("SetConsoleOutputCP")
	procSetCursorPosition           = kernel32.NewProc("SetConsoleCursorPosition")
	procSetConsoleCP                = kernel32.NewProc("SetConsoleCP")
	procSetConsoleTitle             = kernel32.NewProc("SetConsoleTitleW")
	procSetConsoleCursorInfo        = kernel32.NewProc("SetConsoleCursorInfo")
	procFillConsoleOutputCharacter  = kernel32.NewProc("FillConsoleOutputCharacterW")
	procFillConsoleOutputAttribute  = kernel32.NewProc("FillConsoleOutputAttribute")
	procWriteConsoleOutputAttribute = kernel32.NewProc("WriteConsoleOutputAttribute")
	procWriteConsoleOutputCharacter = kernel32.NewProc("WriteConsoleOutputCharacterW")
	procScrollConsoleScreenBuffer   = kernel32.NewProc("ScrollConsoleScreenBufferW")
)
