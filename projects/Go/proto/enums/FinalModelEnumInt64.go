// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumInt64 final model class
type FinalModelEnumInt64 struct {
    buffer *fbe.Buffer  // Final model buffer
    offset int          // Final model buffer offset
}

// Get the allocation size
func (fm *FinalModelEnumInt64) FBEAllocationSize(value EnumInt64) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelEnumInt64) FBESize() int { return 8 }

// Get the final offset
func (fm *FinalModelEnumInt64) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumInt64) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumInt64) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumInt64) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelEnumInt64(buffer *fbe.Buffer, offset int) *FinalModelEnumInt64 {
    return &FinalModelEnumInt64{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm *FinalModelEnumInt64) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm *FinalModelEnumInt64) Get() (EnumInt64, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumInt64(0), 0, errors.New("model is broken")
    }

    return EnumInt64(fbe.ReadInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelEnumInt64) Set(value EnumInt64) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteInt64(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int64(value))
    return fm.FBESize(), nil
}
