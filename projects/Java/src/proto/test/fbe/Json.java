// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding

package test.fbe;

import java.io.*;
import java.lang.*;
import java.lang.reflect.*;
import java.math.*;
import java.nio.charset.*;
import java.time.*;
import java.util.*;

import fbe.*;
import test.*;

import com.google.gson.*;

// Fast Binary Encoding test JSON class
public final class Json
{
    private static final Gson _engine;

    // Get the JSON engine
    public static Gson getEngine() { return _engine; }

    static
    {
        _engine = register(new GsonBuilder()).create();
    }

    private Json() {}

    public static GsonBuilder register(GsonBuilder builder)
    {
        fbe.Json.register(builder);
        proto.fbe.Json.register(builder);
        builder.registerTypeAdapter(EnumSimple.class, new EnumSimpleJson());
        builder.registerTypeAdapter(EnumTyped.class, new EnumTypedJson());
        builder.registerTypeAdapter(FlagsSimple.class, new FlagsSimpleJson());
        builder.registerTypeAdapter(FlagsTyped.class, new FlagsTypedJson());
        return builder;
    }
}