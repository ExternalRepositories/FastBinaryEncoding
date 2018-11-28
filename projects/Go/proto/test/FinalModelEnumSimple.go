// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "errors"
import "../fbe"

// Fast Binary Encoding EnumSimple final model class
type FinalModelEnumSimple struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset
}

// Get the allocation size
func (fm *FinalModelEnumSimple) FBEAllocationSize(value EnumSimple) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelEnumSimple) FBESize() int { return 4 }

// Get the final offset
func (fm *FinalModelEnumSimple) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumSimple) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumSimple) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumSimple) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelEnumSimple(buffer *fbe.Buffer, offset int) *FinalModelEnumSimple {
    return &FinalModelEnumSimple{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm *FinalModelEnumSimple) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm *FinalModelEnumSimple) Get() (EnumSimple, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumSimple(0), 0, errors.New("model is broken")
    }

    return EnumSimple(fbe.ReadInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelEnumSimple) Set(value EnumSimple) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int32(value))
    return fm.FBESize(), nil
}
