/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.firmata.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class FirmataMessageAnalogIO extends FirmataMessage implements Message {

  // Accessors for discriminator values.
  public Byte getMessageType() {
    return (byte) 0xE;
  }

  // Properties.
  protected final byte pin;
  protected final List<Byte> data;

  public FirmataMessageAnalogIO(byte pin, List<Byte> data) {
    super();
    this.pin = pin;
    this.data = data;
  }

  public byte getPin() {
    return pin;
  }

  public List<Byte> getData() {
    return data;
  }

  @Override
  protected void serializeFirmataMessageChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("FirmataMessageAnalogIO");

    // Simple Field (pin)
    writeSimpleField(
        "pin",
        pin,
        writeUnsignedByte(writeBuffer, 4),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Array Field (data)
    writeSimpleTypeArrayField(
        "data",
        data,
        writeSignedByte(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("FirmataMessageAnalogIO");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    FirmataMessageAnalogIO _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (pin)
    lengthInBits += 4;

    // Array field
    if (data != null) {
      lengthInBits += 8 * data.size();
    }

    return lengthInBits;
  }

  public static FirmataMessageBuilder staticParseFirmataMessageBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("FirmataMessageAnalogIO");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte pin =
        readSimpleField(
            "pin", readUnsignedByte(readBuffer, 4), WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    List<Byte> data =
        readCountArrayField(
            "data",
            readSignedByte(readBuffer, 8),
            2,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("FirmataMessageAnalogIO");
    // Create the instance
    return new FirmataMessageAnalogIOBuilderImpl(pin, data);
  }

  public static class FirmataMessageAnalogIOBuilderImpl
      implements FirmataMessage.FirmataMessageBuilder {
    private final byte pin;
    private final List<Byte> data;

    public FirmataMessageAnalogIOBuilderImpl(byte pin, List<Byte> data) {
      this.pin = pin;
      this.data = data;
    }

    public FirmataMessageAnalogIO build() {
      FirmataMessageAnalogIO firmataMessageAnalogIO = new FirmataMessageAnalogIO(pin, data);
      return firmataMessageAnalogIO;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof FirmataMessageAnalogIO)) {
      return false;
    }
    FirmataMessageAnalogIO that = (FirmataMessageAnalogIO) o;
    return (getPin() == that.getPin())
        && (getData() == that.getData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getPin(), getData());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
