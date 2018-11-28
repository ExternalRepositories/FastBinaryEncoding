// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package fbe

import "errors"

// Fast Binary Encoding float32 final model class
type FinalModelFloat struct {
    buffer *Buffer  // Final model buffer
    offset int      // Final model buffer offset
}

// Get the allocation size
func (fm *FinalModelFloat) FBEAllocationSize(value float32) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelFloat) FBESize() int { return 4 }

// Get the final offset
func (fm *FinalModelFloat) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelFloat) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelFloat) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelFloat) FBEUnshift(size int) { fm.offset -= size }

// Create a new final model
func NewFinalModelFloat(buffer *Buffer, offset int) *FinalModelFloat {
    return &FinalModelFloat{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm *FinalModelFloat) Verify() (bool, int) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false, 0
    }

    return true, fm.FBESize()
}

// Get the value
func (fm *FinalModelFloat) Get() (float32, int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0.0, 0, errors.New("model is broken")
    }

    return ReadFloat(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), fm.FBESize(), nil
}

// Set the value
func (fm *FinalModelFloat) Set(value float32) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    WriteFloat(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return fm.FBESize(), nil
}
