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
package org.apache.plc4x.java.canopen.readwrite;

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

public class SDOAbort implements Message {

  // Properties.
  protected final IndexAddress address;
  protected final long code;

  public SDOAbort(IndexAddress address, long code) {
    super();
    this.address = address;
    this.code = code;
  }

  public IndexAddress getAddress() {
    return address;
  }

  public long getCode() {
    return code;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SDOAbort");

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 5));

    // Simple Field (address)
    writeSimpleField("address", address, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (code)
    writeSimpleField("code", code, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("SDOAbort");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    SDOAbort _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 5;

    // Simple field (address)
    lengthInBits += address.getLengthInBits();

    // Simple field (code)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static SDOAbort staticParse(ReadBuffer readBuffer, Object... args) throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static SDOAbort staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("SDOAbort");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 5), (byte) 0x00);

    IndexAddress address =
        readSimpleField(
            "address",
            new DataReaderComplexDefault<>(() -> IndexAddress.staticParse(readBuffer), readBuffer));

    long code = readSimpleField("code", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("SDOAbort");
    // Create the instance
    SDOAbort _sDOAbort;
    _sDOAbort = new SDOAbort(address, code);
    return _sDOAbort;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SDOAbort)) {
      return false;
    }
    SDOAbort that = (SDOAbort) o;
    return (getAddress() == that.getAddress()) && (getCode() == that.getCode()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getAddress(), getCode());
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