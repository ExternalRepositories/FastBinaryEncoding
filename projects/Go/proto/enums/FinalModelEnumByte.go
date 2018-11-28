// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumByte final model class
type FinalModelEnumByte struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset
}

// Get the allocation size
func (fm *FinalModelEnumByte) FBEAllocationSize(value EnumByte) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelEnumByte) FBESize() int { return 1 }

// Get the final offset
func (fm *FinalModelEnumByte) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumByte) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumByte) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumByte) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelEnumByte(buffer *fbe.Buffer, offset int) *FinalModelEnumByte {
    return &FinalModelEnumByte{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm *FinalModelEnumByte) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm *FinalModelEnumByte) Get() (EnumByte, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumByte(0), 0, errors.New("model is broken")
    }

    return EnumByte(fbe.ReadByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelEnumByte) Set(value EnumByte) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteByte(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), byte(value))
    return fm.FBESize(), nil
}
