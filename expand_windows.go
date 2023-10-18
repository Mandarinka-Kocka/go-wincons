package cons

import (
	"fmt"
	"unsafe"
)

// IsValidHandle checks if the current Handle is a valid handle.
//
// Returns:
//
//	bool: True if the Handle is valid, otherwise false.
func (nStdHandle Handle) IsValidHandle() bool {
	return nStdHandle != invalidHandle
}

// IsEnableMode checks if a specific console mode flag is enabled in a DWord mode.
//
// Parameters:
//
//	flag: The mode flag to check.
//
// Returns:
//
//	bool: True if the mode flag is enabled, otherwise false.
func (mode DWord) IsEnableMode(flag int) bool {
	return mode&DWord(flag) != 0
}

func IsEnableMode(hStdout Handle, flag int) bool {
	mode, err := GetMode(hStdout)
	if err != nil {
		return false
	}

	return mode&DWord(flag) != 0
}

// Enables or modifies console modes for the specified standard output handle in Windows.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which modes are to be enabled or modified.
//	flag: An optional variadic list of mode flags to enable or modify.
//
// Returns:
//
//	error: If the function successfully enables or modifies the console modes, it returns nil. Otherwise, it returns an error.
func EnableMode(hStdout Handle, flag ...int) error {
	mode, err := GetMode(hStdout)
	if err != nil {
		return err
	}

	for _, f := range flag {
		mode |= DWord(f)
	}

	return SetMode(hStdout, mode)
}

// Sets the visibility of the cursor for the specified standard output handle in Windows.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which cursor visibility is to be set.
//	visible: A boolean value indicating whether the cursor should be made visible (true) or hidden (false).
//
// Returns:
//
//	error: If the function successfully sets the cursor visibility, it returns nil. Otherwise, it returns an error.
func SetCursorVisible(hStdout Handle, visible bool) error {
	curinfo, err := GetCursorInfo(hStdout)
	if err != nil {
		return err
	}

	curinfo.Visible = visible
	return SetCursorInfo(hStdout, &curinfo)
}

// Sets the size of the cursor for the specified standard output handle in Windows.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which cursor size is to be set.
//	size: The new cursor size.
//
// Returns:
//
//	error: If the function successfully sets the cursor size, it returns nil. Otherwise, it returns an error.
func SetCursorSize(hStdout Handle, size uint32) error {
	curinfo, err := GetCursorInfo(hStdout)
	if err != nil {
		return err
	}

	curinfo.Size = size
	return SetCursorInfo(hStdout, &curinfo)
}

// retrieves the current cursor position within the screen buffer of the specified standard output handle in Windows.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which cursor position is to be obtained.
//
// Returns:
//
//	Coord: The current cursor position represented as a Coord structure.
//	error: If the function successfully retrieves the cursor position, it returns nil. Otherwise, it returns an error.
func GetCursorPosition(hStdout Handle) (Coord, error) {
	var scrbufinfo ScreenBufferInfo
	if err := GetScreenBufferInfo(hStdout, &scrbufinfo); err != nil {
		return Coord{}, err
	}

	return scrbufinfo.CursorPosition, nil
}

// updates the cursor position for the specified standard output handle by adding the given offset.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which the cursor position is to be updated.
//	offset: A Coord structure specifying the horizontal and vertical offset by which to move the cursor.
//
// Returns:
//
//	error: If the function successfully updates the cursor position, it returns nil. Otherwise, it returns an error.
func MoveCursor(hStdout Handle, offset Coord) error {
	curpos, err := GetCursorPosition(hStdout)
	if err != nil {
		return err
	}

	curpos.X += offset.X
	curpos.Y += offset.Y

	if err := SetCursorPosition(hStdout, curpos); err != nil {
		return err
	}

	return nil
}

// updates the cursor position for the specified standard output handle by moving it vertically.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which the cursor position is to be updated.
//	offset: The vertical offset by which to move the cursor.
//
// Returns:
//
//	error: If the function successfully updates the cursor position, it returns nil. Otherwise, it returns an error.
func MoveCursorVert(hStdout Handle, offset int16) error {
	return MoveCursor(hStdout, Coord{X: 0, Y: offset})
}

//	pdates the cursor position for the specified standard output handle by moving it horizontally.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which the cursor position is to be updated.
//	offset: The horizontal offset by which to move the cursor.
//
// Returns:
//
//	error: If the function successfully updates the cursor position, it returns nil. Otherwise, it returns an error.
func MoveCursorHorz(hStdout Handle, offset int16) error {
	return MoveCursor(hStdout, Coord{X: offset, Y: 0})
}

// retrieves a key press event from the specified standard input handle in Windows.
//
// Parameters:
//
//	hStdin: The handle to the standard input stream from which key press events will be retrieved.
//
// Returns:
//
//	uint16: The Unicode character representing the key press event.
//	error: If the function successfully retrieves a key press event, it returns nil. Otherwise, it returns an error.
func GetKeyValue(hStdin Handle) (uint16, error) {
	var (
		buffer  = InputRecord[KeyEventRecord]{}
		counter uint32
	)

	if err := ReadInput(hStdin, unsafe.Pointer(&buffer), 1, &counter); err != nil {
		return 0, err
	}

	if buffer.EventType == KeyEvent && buffer.Event.KeyDown != 0 {
		return buffer.Event.UnicodeChar, nil
	}

	return 0, nil
}

// Pause displays an optional message and waits for a key press event on the specified standard input handle in Windows.
//
// Parameters:
//
//	hStdin: The handle to the standard input stream to listen for a key press event.
//	msg: An optional message to display before waiting for input.
//
// Returns:
//
//	error: If the function successfully waits for a key press event, it returns nil. Otherwise, it returns an error.
func Pause(hStdin Handle, msg string) error {
	var (
		// buffer is a structure used to store console input records, including event type and KeyEventRecord.
		// It is used to read and store input records from the standard input handle in Windows.
		buffer = struct {
			EventType uint16
			KeyEvent  KeyEventRecord
		}{
			EventType: 0,
			KeyEvent:  KeyEventRecord{},
		}
		counter uint32
	)

	if msg != "" {
		fmt.Println(msg)
	}

	for {
		if err := ReadInput(hStdin, unsafe.Pointer(&buffer), 1, &counter); err != nil {
			return err
		}

		if buffer.EventType == KeyEvent && buffer.KeyEvent.KeyDown != 0 {
			return nil
		}
	}

}

// Fill updates a specified portion of the console screen buffer with the given character and attributes for the specified standard output handle in Windows.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which the console screen buffer is to be updated.
//	fill: A CharInfo structure specifying the character and attributes to be used for filling.
//	length: The number of character cells to update.
//	wcoord: The starting coordinates for updating the screen buffer.
//
// Returns:
//
//	error: If the function successfully updates the screen buffer, it returns nil. Otherwise, it returns an error.
func Fill(hStdout Handle, fill CharInfo, length int, wcoord Coord) error {
	if _, err := FillCharacter(hStdout, rune(fill.UnicodeChar), length, wcoord); err != nil {
		return err
	}

	if _, err := FillAttribute(hStdout, fill.Attributes, length, wcoord); err != nil {
		return err
	}

	return nil
}

// clears the entire console screen buffer by filling it with spaces and default attributes for the specified standard output handle in Windows.
//
// Parameters:
//
//	hStdout: The handle to the standard output stream for which the console screen buffer is to be cleared.
//
// Returns:
//
//	error: If the function successfully clears the screen buffer, it returns nil. Otherwise, it returns an error.
func ClearScreenBuffer(hStdout Handle) error {
	var (
		fill       CharInfo
		length     int
		scrbufinfo ScreenBufferInfo
	)

	if err := GetScreenBufferInfo(hStdout, &scrbufinfo); err != nil {
		return err
	}

	length = int(scrbufinfo.Size.X * scrbufinfo.Size.Y)
	fill.UnicodeChar = ' '
	fill.Attributes = ForegroundBlue | ForegroundGreen | ForegroundRed

	if err := Fill(hStdout, fill, length, Coord{}); err != nil {
		return err
	}

	return SetCursorPosition(hStdout, Coord{})
}
