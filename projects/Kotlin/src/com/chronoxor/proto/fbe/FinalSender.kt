// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: proto.fbe
// Version: 1.3.0.0

@file:Suppress("UnusedImport", "unused")

package com.chronoxor.proto.fbe

// Fast Binary Encoding com.chronoxor.proto final sender
@Suppress("MemberVisibilityCanBePrivate", "PropertyName")
open class FinalSender : com.chronoxor.fbe.Sender
{
    // Sender models accessors
    val OrderModel: OrderFinalModel
    val BalanceModel: BalanceFinalModel
    val AccountModel: AccountFinalModel
    val OrderMessageModel: OrderMessageFinalModel
    val BalanceMessageModel: BalanceMessageFinalModel
    val AccountMessageModel: AccountMessageFinalModel

    constructor() : super(true)
    {
        OrderModel = OrderFinalModel(buffer)
        BalanceModel = BalanceFinalModel(buffer)
        AccountModel = AccountFinalModel(buffer)
        OrderMessageModel = OrderMessageFinalModel(buffer)
        BalanceMessageModel = BalanceMessageFinalModel(buffer)
        AccountMessageModel = AccountMessageFinalModel(buffer)
    }

    constructor(buffer: com.chronoxor.fbe.Buffer) : super(buffer, true)
    {
        OrderModel = OrderFinalModel(buffer)
        BalanceModel = BalanceFinalModel(buffer)
        AccountModel = AccountFinalModel(buffer)
        OrderMessageModel = OrderMessageFinalModel(buffer)
        BalanceMessageModel = BalanceMessageFinalModel(buffer)
        AccountMessageModel = AccountMessageFinalModel(buffer)
    }

    @Suppress("JoinDeclarationAndAssignment")
    fun send(obj: Any): Long
    {
        when (obj)
        {
            is com.chronoxor.proto.Order -> if (obj.fbeType == OrderModel.fbeType) return send(obj)
            is com.chronoxor.proto.Balance -> if (obj.fbeType == BalanceModel.fbeType) return send(obj)
            is com.chronoxor.proto.Account -> if (obj.fbeType == AccountModel.fbeType) return send(obj)
            is com.chronoxor.proto.OrderMessage -> if (obj.fbeType == OrderMessageModel.fbeType) return send(obj)
            is com.chronoxor.proto.BalanceMessage -> if (obj.fbeType == BalanceMessageModel.fbeType) return send(obj)
            is com.chronoxor.proto.AccountMessage -> if (obj.fbeType == AccountMessageModel.fbeType) return send(obj)
        }

        return 0
    }

    fun send(value: com.chronoxor.proto.Order): Long
    {
        // Serialize the value into the FBE stream
        val serialized = OrderModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.proto.Order serialization failed!" }
        assert(OrderModel.verify()) { "com.chronoxor.proto.Order validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }
    fun send(value: com.chronoxor.proto.Balance): Long
    {
        // Serialize the value into the FBE stream
        val serialized = BalanceModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.proto.Balance serialization failed!" }
        assert(BalanceModel.verify()) { "com.chronoxor.proto.Balance validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }
    fun send(value: com.chronoxor.proto.Account): Long
    {
        // Serialize the value into the FBE stream
        val serialized = AccountModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.proto.Account serialization failed!" }
        assert(AccountModel.verify()) { "com.chronoxor.proto.Account validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }
    fun send(value: com.chronoxor.proto.OrderMessage): Long
    {
        // Serialize the value into the FBE stream
        val serialized = OrderMessageModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.proto.OrderMessage serialization failed!" }
        assert(OrderMessageModel.verify()) { "com.chronoxor.proto.OrderMessage validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }
    fun send(value: com.chronoxor.proto.BalanceMessage): Long
    {
        // Serialize the value into the FBE stream
        val serialized = BalanceMessageModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.proto.BalanceMessage serialization failed!" }
        assert(BalanceMessageModel.verify()) { "com.chronoxor.proto.BalanceMessage validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }
    fun send(value: com.chronoxor.proto.AccountMessage): Long
    {
        // Serialize the value into the FBE stream
        val serialized = AccountMessageModel.serialize(value)
        assert(serialized > 0) { "com.chronoxor.proto.AccountMessage serialization failed!" }
        assert(AccountMessageModel.verify()) { "com.chronoxor.proto.AccountMessage validation failed!" }

        // Log the value
        if (logging)
        {
            val message = value.toString()
            onSendLog(message)
        }

        // Send the serialized value
        return sendSerialized(serialized)
    }

    // Send message handler
    override fun onSend(buffer: ByteArray, offset: Long, size: Long): Long { throw UnsupportedOperationException("com.chronoxor.proto.fbe.Sender.onSend() not implemented!") }
}