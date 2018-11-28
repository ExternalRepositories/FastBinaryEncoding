// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.1.0.0

package test

import "errors"
import "../fbe"

// Fast Binary Encoding EnumTyped final model class
type FinalModelEnumTyped struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset
}

// Get the allocation size
func (fm *FinalModelEnumTyped) FBEAllocationSize(value EnumTyped) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelEnumTyped) FBESize() int { return 1 }

// Get the final offset
func (fm *FinalModelEnumTyped) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumTyped) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumTyped) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumTyped) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelEnumTyped(buffer *fbe.Buffer, offset int) *FinalModelEnumTyped {
    return &FinalModelEnumTyped{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm *FinalModelEnumTyped) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm *FinalModelEnumTyped) Get() (EnumTyped, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumTyped(0), 0, errors.New("model is broken")
    }

    return EnumTyped(fbe.ReadUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelEnumTyped) Set(value EnumTyped) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint8(value))
    return fm.FBESize(), nil
}
