// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package fbe

import "errors"

// Fast Binary Encoding rune field model class
type FieldModelWChar struct {
    buffer *Buffer  // Field model buffer
    offset int      // Field model buffer offset
}

// Get the field size
func (fm *FieldModelWChar) FBESize() int { return 4 }
// Get the field extra size
func (fm *FieldModelWChar) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelWChar) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelWChar) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelWChar) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelWChar) FBEUnshift(size int) { fm.offset -= size }

// Create a new field model
func NewFieldModelWChar(buffer *Buffer, offset int) *FieldModelWChar {
    return &FieldModelWChar{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm *FieldModelWChar) Verify() bool { return true }

// Get the value
func (fm *FieldModelWChar) Get() (rune, error) {
    return fm.GetDefault('\000')
}

// Get the value with provided default value
func (fm *FieldModelWChar) GetDefault(defaults rune) (rune, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return defaults, nil
    }

    return ReadWChar(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()), nil
}

// Set the value
func (fm *FieldModelWChar) Set(value rune) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    WriteWChar(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), value)
    return nil
}
