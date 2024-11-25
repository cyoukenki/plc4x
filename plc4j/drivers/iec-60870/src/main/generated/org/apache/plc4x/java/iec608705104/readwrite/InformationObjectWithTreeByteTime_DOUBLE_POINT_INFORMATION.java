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
package org.apache.plc4x.java.iec608705104.readwrite;

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

public class InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION
    extends InformationObjectWithTreeByteTime implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.DOUBLE_POINT_INFORMATION_WITH_TIME_TAG;
  }

  // Properties.
  protected final DoublePointInformation diq;
  protected final ThreeOctetBinaryTime cp24Time2a;

  public InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION(
      int address, DoublePointInformation diq, ThreeOctetBinaryTime cp24Time2a) {
    super(address);
    this.diq = diq;
    this.cp24Time2a = cp24Time2a;
  }

  public DoublePointInformation getDiq() {
    return diq;
  }

  public ThreeOctetBinaryTime getCp24Time2a() {
    return cp24Time2a;
  }

  @Override
  protected void serializeInformationObjectWithTreeByteTimeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION");

    // Simple Field (diq)
    writeSimpleField(
        "diq",
        diq,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    // Simple Field (cp24Time2a)
    writeSimpleField(
        "cp24Time2a",
        cp24Time2a,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (diq)
    lengthInBits += diq.getLengthInBits();

    // Simple field (cp24Time2a)
    lengthInBits += cp24Time2a.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectWithTreeByteTimeBuilder
      staticParseInformationObjectWithTreeByteTimeBuilder(
          ReadBuffer readBuffer, TypeIdentification typeIdentification, Byte numTimeByte)
          throws ParseException {
    readBuffer.pullContext("InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    DoublePointInformation diq =
        readSimpleField(
            "diq",
            new DataReaderComplexDefault<>(
                () -> DoublePointInformation.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    ThreeOctetBinaryTime cp24Time2a =
        readSimpleField(
            "cp24Time2a",
            new DataReaderComplexDefault<>(
                () -> ThreeOctetBinaryTime.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION");
    // Create the instance
    return new InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATIONBuilderImpl(
        diq, cp24Time2a);
  }

  public static class InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATIONBuilderImpl
      implements InformationObjectWithTreeByteTime.InformationObjectWithTreeByteTimeBuilder {
    private final DoublePointInformation diq;
    private final ThreeOctetBinaryTime cp24Time2a;

    public InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATIONBuilderImpl(
        DoublePointInformation diq, ThreeOctetBinaryTime cp24Time2a) {
      this.diq = diq;
      this.cp24Time2a = cp24Time2a;
    }

    public InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION build(int address) {
      InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION
          informationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION =
              new InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION(
                  address, diq, cp24Time2a);
      return informationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION)) {
      return false;
    }
    InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION that =
        (InformationObjectWithTreeByteTime_DOUBLE_POINT_INFORMATION) o;
    return (getDiq() == that.getDiq())
        && (getCp24Time2a() == that.getCp24Time2a())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDiq(), getCp24Time2a());
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
