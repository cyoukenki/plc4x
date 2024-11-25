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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class LBusmonInd extends CEMI implements Message {

  // Accessors for discriminator values.
  public Short getMessageCode() {
    return (short) 0x2B;
  }

  // Properties.
  protected final short additionalInformationLength;
  protected final List<CEMIAdditionalInformation> additionalInformation;
  protected final LDataFrame dataFrame;
  protected final Short crc;

  public LBusmonInd(
      short additionalInformationLength,
      List<CEMIAdditionalInformation> additionalInformation,
      LDataFrame dataFrame,
      Short crc) {
    super();
    this.additionalInformationLength = additionalInformationLength;
    this.additionalInformation = additionalInformation;
    this.dataFrame = dataFrame;
    this.crc = crc;
  }

  public short getAdditionalInformationLength() {
    return additionalInformationLength;
  }

  public List<CEMIAdditionalInformation> getAdditionalInformation() {
    return additionalInformation;
  }

  public LDataFrame getDataFrame() {
    return dataFrame;
  }

  public Short getCrc() {
    return crc;
  }

  @Override
  protected void serializeCEMIChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("LBusmonInd");

    // Simple Field (additionalInformationLength)
    writeSimpleField(
        "additionalInformationLength",
        additionalInformationLength,
        writeUnsignedShort(writeBuffer, 8));

    // Array Field (additionalInformation)
    writeComplexTypeArrayField("additionalInformation", additionalInformation, writeBuffer);

    // Simple Field (dataFrame)
    writeSimpleField("dataFrame", dataFrame, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (crc) (Can be skipped, if the value is null)
    writeOptionalField("crc", crc, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("LBusmonInd");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    LBusmonInd _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (additionalInformationLength)
    lengthInBits += 8;

    // Array field
    if (additionalInformation != null) {
      for (Message element : additionalInformation) {
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (dataFrame)
    lengthInBits += dataFrame.getLengthInBits();

    // Optional Field (crc)
    if (crc != null) {
      lengthInBits += 8;
    }

    return lengthInBits;
  }

  public static CEMIBuilder staticParseCEMIBuilder(ReadBuffer readBuffer, Integer size)
      throws ParseException {
    readBuffer.pullContext("LBusmonInd");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short additionalInformationLength =
        readSimpleField("additionalInformationLength", readUnsignedShort(readBuffer, 8));

    List<CEMIAdditionalInformation> additionalInformation =
        readLengthArrayField(
            "additionalInformation",
            new DataReaderComplexDefault<>(
                () -> CEMIAdditionalInformation.staticParse(readBuffer), readBuffer),
            additionalInformationLength);

    LDataFrame dataFrame =
        readSimpleField(
            "dataFrame",
            new DataReaderComplexDefault<>(() -> LDataFrame.staticParse(readBuffer), readBuffer));

    Short crc =
        readOptionalField("crc", readUnsignedShort(readBuffer, 8), dataFrame.getNotAckFrame());

    readBuffer.closeContext("LBusmonInd");
    // Create the instance
    return new LBusmonIndBuilderImpl(
        additionalInformationLength, additionalInformation, dataFrame, crc);
  }

  public static class LBusmonIndBuilderImpl implements CEMI.CEMIBuilder {
    private final short additionalInformationLength;
    private final List<CEMIAdditionalInformation> additionalInformation;
    private final LDataFrame dataFrame;
    private final Short crc;

    public LBusmonIndBuilderImpl(
        short additionalInformationLength,
        List<CEMIAdditionalInformation> additionalInformation,
        LDataFrame dataFrame,
        Short crc) {
      this.additionalInformationLength = additionalInformationLength;
      this.additionalInformation = additionalInformation;
      this.dataFrame = dataFrame;
      this.crc = crc;
    }

    public LBusmonInd build() {
      LBusmonInd lBusmonInd =
          new LBusmonInd(additionalInformationLength, additionalInformation, dataFrame, crc);
      return lBusmonInd;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof LBusmonInd)) {
      return false;
    }
    LBusmonInd that = (LBusmonInd) o;
    return (getAdditionalInformationLength() == that.getAdditionalInformationLength())
        && (getAdditionalInformation() == that.getAdditionalInformation())
        && (getDataFrame() == that.getDataFrame())
        && (getCrc() == that.getCrc())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getAdditionalInformationLength(),
        getAdditionalInformation(),
        getDataFrame(),
        getCrc());
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