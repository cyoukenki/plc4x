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
package org.apache.plc4x.java.cbus.readwrite;

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

public class MonitoredSALLongFormSmartMode extends MonitoredSAL implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final int terminatingByte;
  protected final UnitAddress unitAddress;
  protected final BridgeAddress bridgeAddress;
  protected final ApplicationIdContainer application;
  protected final Byte reservedByte;
  protected final ReplyNetwork replyNetwork;
  protected final SALData salData;

  // Arguments.
  protected final CBusOptions cBusOptions;
  // Reserved Fields
  private Byte reservedField0;

  public MonitoredSALLongFormSmartMode(
      byte salType,
      int terminatingByte,
      UnitAddress unitAddress,
      BridgeAddress bridgeAddress,
      ApplicationIdContainer application,
      Byte reservedByte,
      ReplyNetwork replyNetwork,
      SALData salData,
      CBusOptions cBusOptions) {
    super(salType, cBusOptions);
    this.terminatingByte = terminatingByte;
    this.unitAddress = unitAddress;
    this.bridgeAddress = bridgeAddress;
    this.application = application;
    this.reservedByte = reservedByte;
    this.replyNetwork = replyNetwork;
    this.salData = salData;
    this.cBusOptions = cBusOptions;
  }

  public int getTerminatingByte() {
    return terminatingByte;
  }

  public UnitAddress getUnitAddress() {
    return unitAddress;
  }

  public BridgeAddress getBridgeAddress() {
    return bridgeAddress;
  }

  public ApplicationIdContainer getApplication() {
    return application;
  }

  public Byte getReservedByte() {
    return reservedByte;
  }

  public ReplyNetwork getReplyNetwork() {
    return replyNetwork;
  }

  public SALData getSalData() {
    return salData;
  }

  public boolean getIsUnitAddress() {
    return (boolean) ((((getTerminatingByte()) & (0xff))) == (0x00));
  }

  @Override
  protected void serializeMonitoredSALChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("MonitoredSALLongFormSmartMode");

    // Reserved Field (reserved)
    writeReservedField(
        "reserved",
        reservedField0 != null ? reservedField0 : (byte) 0x05,
        writeByte(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isUnitAddress = getIsUnitAddress();
    writeBuffer.writeVirtual("isUnitAddress", isUnitAddress);

    // Optional Field (unitAddress) (Can be skipped, if the value is null)
    writeOptionalField(
        "unitAddress",
        unitAddress,
        new DataWriterComplexDefault<>(writeBuffer),
        getIsUnitAddress());

    // Optional Field (bridgeAddress) (Can be skipped, if the value is null)
    writeOptionalField(
        "bridgeAddress",
        bridgeAddress,
        new DataWriterComplexDefault<>(writeBuffer),
        !(getIsUnitAddress()));

    // Simple Field (application)
    writeSimpleEnumField(
        "application",
        "ApplicationIdContainer",
        application,
        new DataWriterEnumDefault<>(
            ApplicationIdContainer::getValue,
            ApplicationIdContainer::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Optional Field (reservedByte) (Can be skipped, if the value is null)
    writeOptionalField("reservedByte", reservedByte, writeByte(writeBuffer, 8), getIsUnitAddress());

    // Optional Field (replyNetwork) (Can be skipped, if the value is null)
    writeOptionalField(
        "replyNetwork",
        replyNetwork,
        new DataWriterComplexDefault<>(writeBuffer),
        !(getIsUnitAddress()));

    // Optional Field (salData) (Can be skipped, if the value is null)
    writeOptionalField("salData", salData, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("MonitoredSALLongFormSmartMode");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MonitoredSALLongFormSmartMode _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // Optional Field (unitAddress)
    if (unitAddress != null) {
      lengthInBits += unitAddress.getLengthInBits();
    }

    // Optional Field (bridgeAddress)
    if (bridgeAddress != null) {
      lengthInBits += bridgeAddress.getLengthInBits();
    }

    // Simple field (application)
    lengthInBits += 8;

    // Optional Field (reservedByte)
    if (reservedByte != null) {
      lengthInBits += 8;
    }

    // Optional Field (replyNetwork)
    if (replyNetwork != null) {
      lengthInBits += replyNetwork.getLengthInBits();
    }

    // Optional Field (salData)
    if (salData != null) {
      lengthInBits += salData.getLengthInBits();
    }

    return lengthInBits;
  }

  public static MonitoredSALBuilder staticParseMonitoredSALBuilder(
      ReadBuffer readBuffer, CBusOptions cBusOptions) throws ParseException {
    readBuffer.pullContext("MonitoredSALLongFormSmartMode");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 = readReservedField("reserved", readByte(readBuffer, 8), (byte) 0x05);

    int terminatingByte = readPeekField("terminatingByte", readUnsignedInt(readBuffer, 24));
    boolean isUnitAddress =
        readVirtualField("isUnitAddress", boolean.class, (((terminatingByte) & (0xff))) == (0x00));

    UnitAddress unitAddress =
        readOptionalField(
            "unitAddress",
            new DataReaderComplexDefault<>(() -> UnitAddress.staticParse(readBuffer), readBuffer),
            isUnitAddress);

    BridgeAddress bridgeAddress =
        readOptionalField(
            "bridgeAddress",
            new DataReaderComplexDefault<>(() -> BridgeAddress.staticParse(readBuffer), readBuffer),
            !(isUnitAddress));

    ApplicationIdContainer application =
        readEnumField(
            "application",
            "ApplicationIdContainer",
            new DataReaderEnumDefault<>(
                ApplicationIdContainer::enumForValue, readUnsignedShort(readBuffer, 8)));

    Byte reservedByte = readOptionalField("reservedByte", readByte(readBuffer, 8), isUnitAddress);
    // Validation
    if (!(((isUnitAddress) && ((reservedByte) == (0x00))) || (!(isUnitAddress)))) {
      throw new ParseValidationException("invalid unit address");
    }

    ReplyNetwork replyNetwork =
        readOptionalField(
            "replyNetwork",
            new DataReaderComplexDefault<>(() -> ReplyNetwork.staticParse(readBuffer), readBuffer),
            !(isUnitAddress));

    SALData salData =
        readOptionalField(
            "salData",
            new DataReaderComplexDefault<>(
                () ->
                    SALData.staticParse(
                        readBuffer, (ApplicationId) (application.getApplicationId())),
                readBuffer));

    readBuffer.closeContext("MonitoredSALLongFormSmartMode");
    // Create the instance
    return new MonitoredSALLongFormSmartModeBuilderImpl(
        terminatingByte,
        unitAddress,
        bridgeAddress,
        application,
        reservedByte,
        replyNetwork,
        salData,
        cBusOptions,
        reservedField0);
  }

  public static class MonitoredSALLongFormSmartModeBuilderImpl
      implements MonitoredSAL.MonitoredSALBuilder {
    private final int terminatingByte;
    private final UnitAddress unitAddress;
    private final BridgeAddress bridgeAddress;
    private final ApplicationIdContainer application;
    private final Byte reservedByte;
    private final ReplyNetwork replyNetwork;
    private final SALData salData;
    private final CBusOptions cBusOptions;
    private final Byte reservedField0;

    public MonitoredSALLongFormSmartModeBuilderImpl(
        int terminatingByte,
        UnitAddress unitAddress,
        BridgeAddress bridgeAddress,
        ApplicationIdContainer application,
        Byte reservedByte,
        ReplyNetwork replyNetwork,
        SALData salData,
        CBusOptions cBusOptions,
        Byte reservedField0) {
      this.terminatingByte = terminatingByte;
      this.unitAddress = unitAddress;
      this.bridgeAddress = bridgeAddress;
      this.application = application;
      this.reservedByte = reservedByte;
      this.replyNetwork = replyNetwork;
      this.salData = salData;
      this.cBusOptions = cBusOptions;
      this.reservedField0 = reservedField0;
    }

    public MonitoredSALLongFormSmartMode build(byte salType, CBusOptions cBusOptions) {
      MonitoredSALLongFormSmartMode monitoredSALLongFormSmartMode =
          new MonitoredSALLongFormSmartMode(
              salType,
              terminatingByte,
              unitAddress,
              bridgeAddress,
              application,
              reservedByte,
              replyNetwork,
              salData,
              cBusOptions);
      monitoredSALLongFormSmartMode.reservedField0 = reservedField0;
      return monitoredSALLongFormSmartMode;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MonitoredSALLongFormSmartMode)) {
      return false;
    }
    MonitoredSALLongFormSmartMode that = (MonitoredSALLongFormSmartMode) o;
    return (getTerminatingByte() == that.getTerminatingByte())
        && (getUnitAddress() == that.getUnitAddress())
        && (getBridgeAddress() == that.getBridgeAddress())
        && (getApplication() == that.getApplication())
        && (getReservedByte() == that.getReservedByte())
        && (getReplyNetwork() == that.getReplyNetwork())
        && (getSalData() == that.getSalData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getTerminatingByte(),
        getUnitAddress(),
        getBridgeAddress(),
        getApplication(),
        getReservedByte(),
        getReplyNetwork(),
        getSalData());
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