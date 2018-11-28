// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.1.0.0

package enums

import "errors"
import "../fbe"

// Fast Binary Encoding EnumInt8 field model class
type FieldModelEnumInt8 struct {
    buffer *fbe.Buffer  // Field model buffer
    offset int          // Field model buffer offset
}

// Get the field size
func (fm *FieldModelEnumInt8) FBESize() int { return 1 }
// Get the field extra size
func (fm *FieldModelEnumInt8) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelEnumInt8) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelEnumInt8) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelEnumInt8) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelEnumInt8) FBEUnshift(size int) { fm.offset -= size }

// Create a new field model
func NewFieldModelEnumInt8(buffer *fbe.Buffer, offset int) *FieldModelEnumInt8 {
    return &FieldModelEnumInt8{buffer: buffer, offset: offset}
}

// Check if the value is valid
func (fm *FieldModelEnumInt8) Verify() bool { return true }

// Get the value
func (fm *FieldModelEnumInt8) Get() (EnumInt8, error) {
    return fm.GetDefault(0)
}

// Get the value with provided default value
func (fm *FieldModelEnumInt8) GetDefault(defaults EnumInt8) (EnumInt8, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return EnumInt8(0), nil
    }

    return EnumInt8(fbe.ReadInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())), nil
}

// Set the value
func (fm *FieldModelEnumInt8) Set(value EnumInt8) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbe.WriteInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), int8(value))
    return nil
}
