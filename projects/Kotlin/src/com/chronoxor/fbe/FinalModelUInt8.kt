// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: FBE
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.fbe

// Fast Binary Encoding UByte final model
class FinalModelUInt8(buffer: Buffer, offset: Long) : FinalModel(buffer, offset)
{
    // Get the allocation size
    @Suppress("UNUSED_PARAMETER")
    fun fbeAllocationSize(value: UByte): Long = fbeSize

    // Final size
    override val fbeSize: Long = 1

    // Check if the value is valid
    override fun verify(): Long
    {
        if ((_buffer.offset + fbeOffset + fbeSize) > _buffer.size)
            return Long.MAX_VALUE

        return fbeSize
    }

    // Get the value
    fun get(size: Size): UByte
    {
        if ((_buffer.offset + fbeOffset + fbeSize) > _buffer.size)
            return 0.toUByte()

        size.value = fbeSize
        return readUInt8(fbeOffset)
    }

    // Set the value
    fun set(value: UByte): Long
    {
        assert((_buffer.offset + fbeOffset + fbeSize) <= _buffer.size) { "Model is broken!" }
        if ((_buffer.offset + fbeOffset + fbeSize) > _buffer.size)
            return 0

        write(fbeOffset, value)
        return fbeSize
    }
}