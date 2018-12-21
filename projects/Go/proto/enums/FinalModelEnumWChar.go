// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.2.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumWChar final model
type FinalModelEnumWChar struct {
    // Final model buffer
    buffer *fbe.Buffer
    // Final model buffer offset
    offset int
}

// Create a new EnumWChar final model
func NewFinalModelEnumWChar(buffer *fbe.Buffer, offset int) *FinalModelEnumWChar {
    return &FinalModelEnumWChar{buffer: buffer, offset: offset}
}

// Get the allocation size
func (fm *FinalModelEnumWChar) FBEAllocationSize(value *EnumWChar) int { return fm.FBESize() }

// Get the final size
func (fm *FinalModelEnumWChar) FBESize() int { return 4 }

// Get the final offset
func (fm *FinalModelEnumWChar) FBEOffset() int { return fm.offset }
// Set the final offset
func (fm *FinalModelEnumWChar) SetFBEOffset(value int) { fm.offset = value }

// Shift the current final offset
func (fm *FinalModelEnumWChar) FBEShift(size int) { fm.offset += size }
// Unshift the current final offset
func (fm *FinalModelEnumWChar) FBEUnshift(size int) { fm.offset -= size }

// Check if the value is valid
func (fm *FinalModelEnumWChar) Verify() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return fbe.MaxInt
    }

    return fm.FBESize()
}

// Get the value
func (fm *FinalModelEnumWChar) Get() (*EnumWChar, int, error) {
    var value EnumWChar
    size, err := fm.GetValueDefault(&value, EnumWChar(0))
    return &value, size, err
}

// Get the value with provided default value
func (fm *FinalModelEnumWChar) GetDefault(defaults EnumWChar) (*EnumWChar, int, error) {
    var value EnumWChar
    size, err := fm.GetValueDefault(&value, defaults)
    return &value, size, err
}

// Get the value by the given pointer
func (fm *FinalModelEnumWChar) GetValue(value *EnumWChar) (int, error) {
    return fm.GetValueDefault(value, EnumWChar(0))
}

// Get the value by the given pointer with provided default value
func (fm *FinalModelEnumWChar) GetValueDefault(value *EnumWChar, defaults EnumWChar) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        *value = defaults
        return 0, errors.New("model is broken")
    }

    *value = EnumWChar(fbe.ReadUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset()))
    return fm.FBESize(), nil
}

// Set the value by the given pointer
func (fm *FinalModelEnumWChar) Set(value *EnumWChar) (int, error) {
    return fm.SetValue(*value)
}

// Set the value
func (fm *FinalModelEnumWChar) SetValue(value EnumWChar) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbe.WriteUInt32(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), uint32(value))
    return fm.FBESize(), nil
}