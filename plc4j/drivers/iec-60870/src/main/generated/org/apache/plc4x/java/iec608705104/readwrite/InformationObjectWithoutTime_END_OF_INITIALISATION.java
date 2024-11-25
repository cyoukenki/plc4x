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

public class InformationObjectWithoutTime_END_OF_INITIALISATION extends InformationObjectWithoutTime
    implements Message {

  // Accessors for discriminator values.
  public TypeIdentification getTypeIdentification() {
    return TypeIdentification.END_OF_INITIALISATION;
  }

  // Properties.
  protected final CauseOfInitialization coi;

  public InformationObjectWithoutTime_END_OF_INITIALISATION(
      int address, CauseOfInitialization coi) {
    super(address);
    this.coi = coi;
  }

  public CauseOfInitialization getCoi() {
    return coi;
  }

  @Override
  protected void serializeInformationObjectWithoutTimeChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("InformationObjectWithoutTime_END_OF_INITIALISATION");

    // Simple Field (coi)
    writeSimpleField(
        "coi",
        coi,
        new DataWriterComplexDefault<>(writeBuffer),
        WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    writeBuffer.popContext("InformationObjectWithoutTime_END_OF_INITIALISATION");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    InformationObjectWithoutTime_END_OF_INITIALISATION _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (coi)
    lengthInBits += coi.getLengthInBits();

    return lengthInBits;
  }

  public static InformationObjectWithoutTimeBuilder staticParseInformationObjectWithoutTimeBuilder(
      ReadBuffer readBuffer, TypeIdentification typeIdentification, Byte numTimeByte)
      throws ParseException {
    readBuffer.pullContext("InformationObjectWithoutTime_END_OF_INITIALISATION");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    CauseOfInitialization coi =
        readSimpleField(
            "coi",
            new DataReaderComplexDefault<>(
                () -> CauseOfInitialization.staticParse(readBuffer), readBuffer),
            WithOption.WithByteOrder(ByteOrder.LITTLE_ENDIAN));

    readBuffer.closeContext("InformationObjectWithoutTime_END_OF_INITIALISATION");
    // Create the instance
    return new InformationObjectWithoutTime_END_OF_INITIALISATIONBuilderImpl(coi);
  }

  public static class InformationObjectWithoutTime_END_OF_INITIALISATIONBuilderImpl
      implements InformationObjectWithoutTime.InformationObjectWithoutTimeBuilder {
    private final CauseOfInitialization coi;

    public InformationObjectWithoutTime_END_OF_INITIALISATIONBuilderImpl(
        CauseOfInitialization coi) {
      this.coi = coi;
    }

    public InformationObjectWithoutTime_END_OF_INITIALISATION build(int address) {
      InformationObjectWithoutTime_END_OF_INITIALISATION
          informationObjectWithoutTime_END_OF_INITIALISATION =
              new InformationObjectWithoutTime_END_OF_INITIALISATION(address, coi);
      return informationObjectWithoutTime_END_OF_INITIALISATION;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof InformationObjectWithoutTime_END_OF_INITIALISATION)) {
      return false;
    }
    InformationObjectWithoutTime_END_OF_INITIALISATION that =
        (InformationObjectWithoutTime_END_OF_INITIALISATION) o;
    return (getCoi() == that.getCoi()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCoi());
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
