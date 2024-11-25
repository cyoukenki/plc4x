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
package org.apache.plc4x.java.modbus.readwrite;

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

public class ModbusPDUWriteFileRecordRequest extends ModbusPDU implements Message {

  // Accessors for discriminator values.
  public Boolean getErrorFlag() {
    return (boolean) false;
  }

  public Byte getFunctionFlag() {
    return (byte) 0x15;
  }

  public Boolean getResponse() {
    return (boolean) false;
  }

  // Properties.
  protected final List<ModbusPDUWriteFileRecordRequestItem> items;

  public ModbusPDUWriteFileRecordRequest(List<ModbusPDUWriteFileRecordRequestItem> items) {
    super();
    this.items = items;
  }

  public List<ModbusPDUWriteFileRecordRequestItem> getItems() {
    return items;
  }

  @Override
  protected void serializeModbusPDUChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("ModbusPDUWriteFileRecordRequest");

    // Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    short byteCount = (short) (ARRAY_SIZE_IN_BYTES(getItems()));
    writeImplicitField("byteCount", byteCount, writeUnsignedShort(writeBuffer, 8));

    // Array Field (items)
    writeComplexTypeArrayField("items", items, writeBuffer);

    writeBuffer.popContext("ModbusPDUWriteFileRecordRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ModbusPDUWriteFileRecordRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (byteCount)
    lengthInBits += 8;

    // Array field
    if (items != null) {
      for (Message element : items) {
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ModbusPDUBuilder staticParseModbusPDUBuilder(
      ReadBuffer readBuffer, Boolean response) throws ParseException {
    readBuffer.pullContext("ModbusPDUWriteFileRecordRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short byteCount = readImplicitField("byteCount", readUnsignedShort(readBuffer, 8));

    List<ModbusPDUWriteFileRecordRequestItem> items =
        readLengthArrayField(
            "items",
            new DataReaderComplexDefault<>(
                () -> ModbusPDUWriteFileRecordRequestItem.staticParse(readBuffer), readBuffer),
            byteCount);

    readBuffer.closeContext("ModbusPDUWriteFileRecordRequest");
    // Create the instance
    return new ModbusPDUWriteFileRecordRequestBuilderImpl(items);
  }

  public static class ModbusPDUWriteFileRecordRequestBuilderImpl
      implements ModbusPDU.ModbusPDUBuilder {
    private final List<ModbusPDUWriteFileRecordRequestItem> items;

    public ModbusPDUWriteFileRecordRequestBuilderImpl(
        List<ModbusPDUWriteFileRecordRequestItem> items) {
      this.items = items;
    }

    public ModbusPDUWriteFileRecordRequest build() {
      ModbusPDUWriteFileRecordRequest modbusPDUWriteFileRecordRequest =
          new ModbusPDUWriteFileRecordRequest(items);
      return modbusPDUWriteFileRecordRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ModbusPDUWriteFileRecordRequest)) {
      return false;
    }
    ModbusPDUWriteFileRecordRequest that = (ModbusPDUWriteFileRecordRequest) o;
    return (getItems() == that.getItems()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getItems());
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
