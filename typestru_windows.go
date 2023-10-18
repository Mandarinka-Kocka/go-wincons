package cons

import "unsafe"

type (
	Handle uintptr
	DWord  uint32
)

type CursorInfo struct {
	Size    uint32 // The size of the cursor.
	Visible bool   // Indicates whether the cursor is visible (true) or not (false).
}

type Coord struct {
	X int16 // The X-coordinate.
	Y int16 // The Y-coordinate.
}

type SmallRect struct {
	Left   int16 // The left boundary.
	Top    int16 // The top boundary.
	Right  int16 // The right boundary.
	Bottom int16 // The bottom boundary.
}

type ScreenBufferInfo struct {
	Size              Coord     // The size of the screen buffer.
	CursorPosition    Coord     // The current cursor position.
	Attributes        uint16    // The attributes of the screen buffer.
	Window            SmallRect // The window area within the screen buffer.
	MaximumWindowSize Coord     // The maximum window size.
}

type KeyEventRecord struct {
	KeyDown         int32  // Indicates if a key is pressed down (1) or released (0).
	RepeatCount     uint16 // Number of times the key was repeated.
	VirtualKeyCode  uint16 // Virtual key code of the key.
	VirtualScanCode uint16 // Virtual scan code of the key.
	UnicodeChar     uint16 // Unicode character produced by the key.
	ControlKeyState uint32 // State of control keys (e.g., SHIFT, CTRL, ALT).
}

type MouseEventRecord struct {
	MousePosition   Coord  // The position of the mouse cursor.
	ButtonState     uint32 // State of mouse buttons (e.g., left, right, middle).
	ControlKeyState uint32 // State of control keys (e.g., SHIFT, CTRL, ALT).
	EventFlags      uint32 // Flags indicating the type of mouse event.
}

type WindowBufferSizeRecord struct {
	Size Coord // The new size of the window's buffer.
}

type MenuEventRecord struct {
	CommandId uint // Identifier of the menu command.
}

type FocusEventRecord struct {
	SetFocus bool // Indicates whether focus is set (true) or lost (false).
}

type CharInfo struct {
	UnicodeChar uint16 // The Unicode character to be displayed.
	Attributes  uint16 // Display attributes (e.g., foreground and background colors).
}

type InputRecord[T KeyEventRecord | MouseEventRecord | WindowBufferSizeRecord | MenuEventRecord | FocusEventRecord] struct {
	EventType uint16
	Event     T
}

// Converts a typed pointer to an unsafe pointer and then to a uintptr.
//
// Parameters:
//   p: A pointer to a value of any type.
//
// Returns:
//   uintptr: The uintptr representation of the pointer p.
func touintptr[T any](p *T) uintptr {
	return uintptr(unsafe.Pointer(p))
}

// Converts a typed pointer to an unsafe pointer and then to a uintptr, treating it as an int32.
//
// Parameters:
//   p: A pointer to a value of any type.
//
// Returns:
//   uintptr: The uintptr representation of the pointer p, interpreted as an int32.
func strutouintptr[T any](p *T) uintptr {
	return uintptr(*(*int32)(unsafe.Pointer(p)))
}
