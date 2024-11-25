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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetConfirmedServiceRequestReinitializeDevice extends BACnetConfirmedServiceRequest
    implements Message {

  // Accessors for discriminator values.
  public BACnetConfirmedServiceChoice getServiceChoice() {
    return BACnetConfirmedServiceChoice.REINITIALIZE_DEVICE;
  }

  // Properties.
  protected final BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
      reinitializedStateOfDevice;
  protected final BACnetContextTagCharacterString password;

  // Arguments.
  protected final Long serviceRequestLength;

  public BACnetConfirmedServiceRequestReinitializeDevice(
      BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
          reinitializedStateOfDevice,
      BACnetContextTagCharacterString password,
      Long serviceRequestLength) {
    super(serviceRequestLength);
    this.reinitializedStateOfDevice = reinitializedStateOfDevice;
    this.password = password;
    this.serviceRequestLength = serviceRequestLength;
  }

  public BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
      getReinitializedStateOfDevice() {
    return reinitializedStateOfDevice;
  }

  public BACnetContextTagCharacterString getPassword() {
    return password;
  }

  @Override
  protected void serializeBACnetConfirmedServiceRequestChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConfirmedServiceRequestReinitializeDevice");

    // Simple Field (reinitializedStateOfDevice)
    writeSimpleField(
        "reinitializedStateOfDevice",
        reinitializedStateOfDevice,
        new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (password) (Can be skipped, if the value is null)
    writeOptionalField("password", password, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetConfirmedServiceRequestReinitializeDevice");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConfirmedServiceRequestReinitializeDevice _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (reinitializedStateOfDevice)
    lengthInBits += reinitializedStateOfDevice.getLengthInBits();

    // Optional Field (password)
    if (password != null) {
      lengthInBits += password.getLengthInBits();
    }

    return lengthInBits;
  }

  public static BACnetConfirmedServiceRequestBuilder
      staticParseBACnetConfirmedServiceRequestBuilder(
          ReadBuffer readBuffer, Long serviceRequestLength) throws ParseException {
    readBuffer.pullContext("BACnetConfirmedServiceRequestReinitializeDevice");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
        reinitializedStateOfDevice =
            readSimpleField(
                "reinitializedStateOfDevice",
                new DataReaderComplexDefault<>(
                    () ->
                        BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
                            .staticParse(
                                readBuffer,
                                (short) (0),
                                (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                    readBuffer));

    BACnetContextTagCharacterString password =
        readOptionalField(
            "password",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagCharacterString)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.CHARACTER_STRING)),
                readBuffer));

    readBuffer.closeContext("BACnetConfirmedServiceRequestReinitializeDevice");
    // Create the instance
    return new BACnetConfirmedServiceRequestReinitializeDeviceBuilderImpl(
        reinitializedStateOfDevice, password, serviceRequestLength);
  }

  public static class BACnetConfirmedServiceRequestReinitializeDeviceBuilderImpl
      implements BACnetConfirmedServiceRequest.BACnetConfirmedServiceRequestBuilder {
    private final BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
        reinitializedStateOfDevice;
    private final BACnetContextTagCharacterString password;
    private final Long serviceRequestLength;

    public BACnetConfirmedServiceRequestReinitializeDeviceBuilderImpl(
        BACnetConfirmedServiceRequestReinitializeDeviceReinitializedStateOfDeviceTagged
            reinitializedStateOfDevice,
        BACnetContextTagCharacterString password,
        Long serviceRequestLength) {
      this.reinitializedStateOfDevice = reinitializedStateOfDevice;
      this.password = password;
      this.serviceRequestLength = serviceRequestLength;
    }

    public BACnetConfirmedServiceRequestReinitializeDevice build(Long serviceRequestLength) {

      BACnetConfirmedServiceRequestReinitializeDevice
          bACnetConfirmedServiceRequestReinitializeDevice =
              new BACnetConfirmedServiceRequestReinitializeDevice(
                  reinitializedStateOfDevice, password, serviceRequestLength);
      return bACnetConfirmedServiceRequestReinitializeDevice;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConfirmedServiceRequestReinitializeDevice)) {
      return false;
    }
    BACnetConfirmedServiceRequestReinitializeDevice that =
        (BACnetConfirmedServiceRequestReinitializeDevice) o;
    return (getReinitializedStateOfDevice() == that.getReinitializedStateOfDevice())
        && (getPassword() == that.getPassword())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getReinitializedStateOfDevice(), getPassword());
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
