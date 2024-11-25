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
package org.apache.plc4x.java.s7.readwrite;

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

public class S7PayloadUserDataItemCpuFunctionReadSzlRequest extends S7PayloadUserDataItem
    implements Message {

  // Accessors for discriminator values.
  public Byte getCpuFunctionGroup() {
    return (byte) 0x04;
  }

  public Byte getCpuFunctionType() {
    return (byte) 0x04;
  }

  public Short getCpuSubfunction() {
    return (short) 0x01;
  }

  // Properties.
  protected final SzlId szlId;
  protected final int szlIndex;

  public S7PayloadUserDataItemCpuFunctionReadSzlRequest(
      DataTransportErrorCode returnCode,
      DataTransportSize transportSize,
      int dataLength,
      SzlId szlId,
      int szlIndex) {
    super(returnCode, transportSize, dataLength);
    this.szlId = szlId;
    this.szlIndex = szlIndex;
  }

  public SzlId getSzlId() {
    return szlId;
  }

  public int getSzlIndex() {
    return szlIndex;
  }

  @Override
  protected void serializeS7PayloadUserDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest");

    // Simple Field (szlId)
    writeSimpleField("szlId", szlId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (szlIndex)
    writeSimpleField("szlIndex", szlIndex, writeUnsignedInt(writeBuffer, 16));

    writeBuffer.popContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadUserDataItemCpuFunctionReadSzlRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (szlId)
    lengthInBits += szlId.getLengthInBits();

    // Simple field (szlIndex)
    lengthInBits += 16;

    return lengthInBits;
  }

  public static S7PayloadUserDataItemBuilder staticParseS7PayloadUserDataItemBuilder(
      ReadBuffer readBuffer, Byte cpuFunctionGroup, Byte cpuFunctionType, Short cpuSubfunction)
      throws ParseException {
    readBuffer.pullContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    SzlId szlId =
        readSimpleField(
            "szlId",
            new DataReaderComplexDefault<>(() -> SzlId.staticParse(readBuffer), readBuffer));

    int szlIndex = readSimpleField("szlIndex", readUnsignedInt(readBuffer, 16));

    readBuffer.closeContext("S7PayloadUserDataItemCpuFunctionReadSzlRequest");
    // Create the instance
    return new S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilderImpl(szlId, szlIndex);
  }

  public static class S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilderImpl
      implements S7PayloadUserDataItem.S7PayloadUserDataItemBuilder {
    private final SzlId szlId;
    private final int szlIndex;

    public S7PayloadUserDataItemCpuFunctionReadSzlRequestBuilderImpl(SzlId szlId, int szlIndex) {
      this.szlId = szlId;
      this.szlIndex = szlIndex;
    }

    public S7PayloadUserDataItemCpuFunctionReadSzlRequest build(
        DataTransportErrorCode returnCode, DataTransportSize transportSize, int dataLength) {
      S7PayloadUserDataItemCpuFunctionReadSzlRequest
          s7PayloadUserDataItemCpuFunctionReadSzlRequest =
              new S7PayloadUserDataItemCpuFunctionReadSzlRequest(
                  returnCode, transportSize, dataLength, szlId, szlIndex);
      return s7PayloadUserDataItemCpuFunctionReadSzlRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadUserDataItemCpuFunctionReadSzlRequest)) {
      return false;
    }
    S7PayloadUserDataItemCpuFunctionReadSzlRequest that =
        (S7PayloadUserDataItemCpuFunctionReadSzlRequest) o;
    return (getSzlId() == that.getSzlId())
        && (getSzlIndex() == that.getSzlIndex())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getSzlId(), getSzlIndex());
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