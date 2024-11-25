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

public class BACnetActionCommand implements Message {

  // Properties.
  protected final BACnetContextTagObjectIdentifier deviceIdentifier;
  protected final BACnetContextTagObjectIdentifier objectIdentifier;
  protected final BACnetPropertyIdentifierTagged propertyIdentifier;
  protected final BACnetContextTagUnsignedInteger arrayIndex;
  protected final BACnetConstructedData propertyValue;
  protected final BACnetContextTagUnsignedInteger priority;
  protected final BACnetContextTagBoolean postDelay;
  protected final BACnetContextTagBoolean quitOnFailure;
  protected final BACnetContextTagBoolean writeSuccessful;

  public BACnetActionCommand(
      BACnetContextTagObjectIdentifier deviceIdentifier,
      BACnetContextTagObjectIdentifier objectIdentifier,
      BACnetPropertyIdentifierTagged propertyIdentifier,
      BACnetContextTagUnsignedInteger arrayIndex,
      BACnetConstructedData propertyValue,
      BACnetContextTagUnsignedInteger priority,
      BACnetContextTagBoolean postDelay,
      BACnetContextTagBoolean quitOnFailure,
      BACnetContextTagBoolean writeSuccessful) {
    super();
    this.deviceIdentifier = deviceIdentifier;
    this.objectIdentifier = objectIdentifier;
    this.propertyIdentifier = propertyIdentifier;
    this.arrayIndex = arrayIndex;
    this.propertyValue = propertyValue;
    this.priority = priority;
    this.postDelay = postDelay;
    this.quitOnFailure = quitOnFailure;
    this.writeSuccessful = writeSuccessful;
  }

  public BACnetContextTagObjectIdentifier getDeviceIdentifier() {
    return deviceIdentifier;
  }

  public BACnetContextTagObjectIdentifier getObjectIdentifier() {
    return objectIdentifier;
  }

  public BACnetPropertyIdentifierTagged getPropertyIdentifier() {
    return propertyIdentifier;
  }

  public BACnetContextTagUnsignedInteger getArrayIndex() {
    return arrayIndex;
  }

  public BACnetConstructedData getPropertyValue() {
    return propertyValue;
  }

  public BACnetContextTagUnsignedInteger getPriority() {
    return priority;
  }

  public BACnetContextTagBoolean getPostDelay() {
    return postDelay;
  }

  public BACnetContextTagBoolean getQuitOnFailure() {
    return quitOnFailure;
  }

  public BACnetContextTagBoolean getWriteSuccessful() {
    return writeSuccessful;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetActionCommand");

    // Optional Field (deviceIdentifier) (Can be skipped, if the value is null)
    writeOptionalField(
        "deviceIdentifier", deviceIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (objectIdentifier)
    writeSimpleField(
        "objectIdentifier", objectIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (propertyIdentifier)
    writeSimpleField(
        "propertyIdentifier", propertyIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (arrayIndex) (Can be skipped, if the value is null)
    writeOptionalField("arrayIndex", arrayIndex, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (propertyValue) (Can be skipped, if the value is null)
    writeOptionalField("propertyValue", propertyValue, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (priority) (Can be skipped, if the value is null)
    writeOptionalField("priority", priority, new DataWriterComplexDefault<>(writeBuffer));

    // Optional Field (postDelay) (Can be skipped, if the value is null)
    writeOptionalField("postDelay", postDelay, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (quitOnFailure)
    writeSimpleField("quitOnFailure", quitOnFailure, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (writeSuccessful)
    writeSimpleField(
        "writeSuccessful", writeSuccessful, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetActionCommand");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    BACnetActionCommand _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Optional Field (deviceIdentifier)
    if (deviceIdentifier != null) {
      lengthInBits += deviceIdentifier.getLengthInBits();
    }

    // Simple field (objectIdentifier)
    lengthInBits += objectIdentifier.getLengthInBits();

    // Simple field (propertyIdentifier)
    lengthInBits += propertyIdentifier.getLengthInBits();

    // Optional Field (arrayIndex)
    if (arrayIndex != null) {
      lengthInBits += arrayIndex.getLengthInBits();
    }

    // Optional Field (propertyValue)
    if (propertyValue != null) {
      lengthInBits += propertyValue.getLengthInBits();
    }

    // Optional Field (priority)
    if (priority != null) {
      lengthInBits += priority.getLengthInBits();
    }

    // Optional Field (postDelay)
    if (postDelay != null) {
      lengthInBits += postDelay.getLengthInBits();
    }

    // Simple field (quitOnFailure)
    lengthInBits += quitOnFailure.getLengthInBits();

    // Simple field (writeSuccessful)
    lengthInBits += writeSuccessful.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetActionCommand staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static BACnetActionCommand staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetActionCommand");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagObjectIdentifier deviceIdentifier =
        readOptionalField(
            "deviceIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetContextTagObjectIdentifier objectIdentifier =
        readSimpleField(
            "objectIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagObjectIdentifier)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.BACNET_OBJECT_IDENTIFIER)),
                readBuffer));

    BACnetPropertyIdentifierTagged propertyIdentifier =
        readSimpleField(
            "propertyIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetPropertyIdentifierTagged.staticParse(
                        readBuffer, (short) (2), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetContextTagUnsignedInteger arrayIndex =
        readOptionalField(
            "arrayIndex",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (3),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetConstructedData propertyValue =
        readOptionalField(
            "propertyValue",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetConstructedData.staticParse(
                        readBuffer,
                        (short) (4),
                        (BACnetObjectType) (objectIdentifier.getObjectType()),
                        (BACnetPropertyIdentifier) (propertyIdentifier.getValue()),
                        (BACnetTagPayloadUnsignedInteger)
                            (((((arrayIndex) != (null)) ? arrayIndex.getPayload() : null)))),
                readBuffer));

    BACnetContextTagUnsignedInteger priority =
        readOptionalField(
            "priority",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (5),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagBoolean postDelay =
        readOptionalField(
            "postDelay",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagBoolean)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (6), (BACnetDataType) (BACnetDataType.BOOLEAN)),
                readBuffer));

    BACnetContextTagBoolean quitOnFailure =
        readSimpleField(
            "quitOnFailure",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagBoolean)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (7), (BACnetDataType) (BACnetDataType.BOOLEAN)),
                readBuffer));

    BACnetContextTagBoolean writeSuccessful =
        readSimpleField(
            "writeSuccessful",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagBoolean)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (8), (BACnetDataType) (BACnetDataType.BOOLEAN)),
                readBuffer));

    readBuffer.closeContext("BACnetActionCommand");
    // Create the instance
    BACnetActionCommand _bACnetActionCommand;
    _bACnetActionCommand =
        new BACnetActionCommand(
            deviceIdentifier,
            objectIdentifier,
            propertyIdentifier,
            arrayIndex,
            propertyValue,
            priority,
            postDelay,
            quitOnFailure,
            writeSuccessful);
    return _bACnetActionCommand;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetActionCommand)) {
      return false;
    }
    BACnetActionCommand that = (BACnetActionCommand) o;
    return (getDeviceIdentifier() == that.getDeviceIdentifier())
        && (getObjectIdentifier() == that.getObjectIdentifier())
        && (getPropertyIdentifier() == that.getPropertyIdentifier())
        && (getArrayIndex() == that.getArrayIndex())
        && (getPropertyValue() == that.getPropertyValue())
        && (getPriority() == that.getPriority())
        && (getPostDelay() == that.getPostDelay())
        && (getQuitOnFailure() == that.getQuitOnFailure())
        && (getWriteSuccessful() == that.getWriteSuccessful())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        getDeviceIdentifier(),
        getObjectIdentifier(),
        getPropertyIdentifier(),
        getArrayIndex(),
        getPropertyValue(),
        getPriority(),
        getPostDelay(),
        getQuitOnFailure(),
        getWriteSuccessful());
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