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

public abstract class InformationObjectWithSevenByteTime extends InformationObject
    implements Message {

  // Accessors for discriminator values.
  public Byte getNumTimeByte() {
    return (byte) 7;
  }

  // Abstract accessors for discriminator values.
  public abstract TypeIdentification getTypeIdentification();

  public InformationObjectWithSevenByteTime(int address) {
    super(address);
  }

  public abstract SevenOctetBinaryTime getCp56Time2a();

  protected abstract void serializeInformationObjectWithSevenByteTimeChild(WriteBuffer writeBuffer)
      throws SerializationException;

  @Override
  protected void serializeInformationObjectChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObjectWithSevenByteTime");

    // Switch field (Serialize the sub-type)
    serializeInformationObjectWithSevenByteTimeChild(writeBuffer);

    writeBuffer.popContext("InformationObjectWithSevenByteTime");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObjectWithSevenByteTime _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static InformationObjectBuilder staticParseInformationObjectBuilder(
      ReadBuffer readBuffer, TypeIdentification typeIdentification, Byte numTimeByte)
      throws ParseException {
    readBuffer.pullContext("InformationObjectWithSevenByteTime");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    InformationObjectWithSevenByteTimeBuilder builder = null;
    if (EvaluationHelper.equals(
        typeIdentification, TypeIdentification.SINGLE_POINT_INFORMATION_WITH_TIME_TAG_CP56TIME2A)) {
      builder =
          InformationObjectWithSevenByteTime_SINGLE_POINT_INFORMATION
              .staticParseInformationObjectWithSevenByteTimeBuilder(
                  readBuffer, typeIdentification, numTimeByte);
    } else if (EvaluationHelper.equals(
        typeIdentification, TypeIdentification.DOUBLE_POINT_INFORMATION_WITH_TIME_TAG_CP56TIME2A)) {
      builder =
          InformationObjectWithSevenByteTime_DOUBLE_POINT_INFORMATION
              .staticParseInformationObjectWithSevenByteTimeBuilder(
                  readBuffer, typeIdentification, numTimeByte);
    } else if (EvaluationHelper.equals(
        typeIdentification,
        TypeIdentification.STEP_POSITION_INFORMATION_WITH_TIME_TAG_CP56TIME2A)) {
      builder =
          InformationObjectWithSevenByteTime_STEP_POSITION_INFORMATION
              .staticParseInformationObjectWithSevenByteTimeBuilder(
                  readBuffer, typeIdentification, numTimeByte);
    } else if (EvaluationHelper.equals(
        typeIdentification, TypeIdentification.BITSTRING_OF_32_BIT_WITH_TIME_TAG_CP56TIME2A)) {
      builder =
          InformationObjectWithSevenByteTime_BITSTRING_OF_32_BIT
              .staticParseInformationObjectWithSevenByteTimeBuilder(
                  readBuffer, typeIdentification, numTimeByte);
    } else if (EvaluationHelper.equals(
        typeIdentification,
        TypeIdentification.MEASURED_VALUE_NORMALISED_VALUE_WITH_TIME_TAG_CP56TIME2A)) {
      builder =
          InformationObjectWithSevenByteTime_MEASURED_VALUE_NORMALISED_VALUE
              .staticParseInformationObjectWithSevenByteTimeBuilder(
                  readBuffer, typeIdentification, numTimeByte);
    } else if (EvaluationHelper.equals(
        typeIdentification,
        TypeIdentification.MEASURED_VALUE_SCALED_VALUE_WITH_TIME_TAG_CP56TIME2A)) {
      builder =
          InformationObjectWithSevenByteTime_MEASURED_VALUE_SCALED_VALUE
              .staticParseInformationObjectWithSevenByteTimeBuilder(
                  readBuffer, typeIdentification, numTimeByte);
    } else if (EvaluationHelper.equals(
        typeIdentification,
        TypeIdentification.MEASURED_VALUE_SHORT_FLOATING_POINT_NUMBER_WITH_TIME_TAG_CP56TIME2A)) {
      builder =
          InformationObjectWithSevenByteTime_MEASURED_VALUE_SHORT_FLOATING_POINT_NUMBER
              .staticParseInformationObjectWithSevenByteTimeBuilder(
                  readBuffer, typeIdentification, numTimeByte);
    } else if (EvaluationHelper.equals(
        typeIdentification, TypeIdentification.INTEGRATED_TOTALS_WITH_TIME_TAG_CP56TIME2A)) {
      builder =
          InformationObjectWithSevenByteTime_INTEGRATED_TOTALS
              .staticParseInformationObjectWithSevenByteTimeBuilder(
                  readBuffer, typeIdentification, numTimeByte);
    } else if (EvaluationHelper.equals(
        typeIdentification,
        TypeIdentification.EVENT_OF_PROTECTION_EQUIPMENT_WITH_TIME_TAG_CP56TIME2A)) {
      builder =
          InformationObjectWithSevenByteTime_EVENT_OF_PROTECTION_EQUIPMENT
              .staticParseInformationObjectWithSevenByteTimeBuilder(
                  readBuffer, typeIdentification, numTimeByte);
    } else if (EvaluationHelper.equals(
        typeIdentification,
        TypeIdentification.PACKED_START_EVENTS_OF_PROTECTION_EQUIPMENT_WITH_TIME_TAG_CP56TIME2A)) {
      builder =
          InformationObjectWithSevenByteTime_PACKED_START_EVENTS_OF_PROTECTION_EQUIPMENT
              .staticParseInformationObjectWithSevenByteTimeBuilder(
                  readBuffer, typeIdentification, numTimeByte);
    } else if (EvaluationHelper.equals(
        typeIdentification,
        TypeIdentification
            .PACKED_OUTPUT_CIRCUIT_INFORMATION_OF_PROTECTION_EQUIPMENT_WITH_TIME_TAG_CP56TIME2A)) {
      builder =
          InformationObjectWithSevenByteTime_PACKED_OUTPUT_CIRCUIT_INFORMATION_OF_PROTECTION_EQUIPMENT
              .staticParseInformationObjectWithSevenByteTimeBuilder(
                  readBuffer, typeIdentification, numTimeByte);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "typeIdentification="
              + typeIdentification
              + "]");
    }

    readBuffer.closeContext("InformationObjectWithSevenByteTime");
    // Create the instance
    return new InformationObjectWithSevenByteTimeBuilderImpl(builder);
  }

  public interface InformationObjectWithSevenByteTimeBuilder {
    InformationObjectWithSevenByteTime build(int address);
  }

  public static class InformationObjectWithSevenByteTimeBuilderImpl
      implements InformationObject.InformationObjectBuilder {
    private final InformationObjectWithSevenByteTimeBuilder builder;

    public InformationObjectWithSevenByteTimeBuilderImpl(
        InformationObjectWithSevenByteTimeBuilder builder) {
      this.builder = builder;
    }

    public InformationObjectWithSevenByteTime build(int address) {
      return builder.build(address);
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObjectWithSevenByteTime)) {
      return false;
    }
    InformationObjectWithSevenByteTime that = (InformationObjectWithSevenByteTime) o;
    return super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode());
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
