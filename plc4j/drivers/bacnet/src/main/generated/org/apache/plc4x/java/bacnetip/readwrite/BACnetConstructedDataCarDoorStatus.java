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

import java.math.BigInteger;
import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class BACnetConstructedDataCarDoorStatus extends BACnetConstructedData implements Message {

  // Accessors for discriminator values.
  public BACnetObjectType getObjectTypeArgument() {
    return null;
  }

  public BACnetPropertyIdentifier getPropertyIdentifierArgument() {
    return BACnetPropertyIdentifier.CAR_DOOR_STATUS;
  }

  // Properties.
  protected final BACnetApplicationTagUnsignedInteger numberOfDataElements;
  protected final List<BACnetDoorStatusTagged> carDoorStatus;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

  public BACnetConstructedDataCarDoorStatus(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetApplicationTagUnsignedInteger numberOfDataElements,
      List<BACnetDoorStatusTagged> carDoorStatus,
      Short tagNumber,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument);
    this.numberOfDataElements = numberOfDataElements;
    this.carDoorStatus = carDoorStatus;
    this.tagNumber = tagNumber;
    this.arrayIndexArgument = arrayIndexArgument;
  }

  public BACnetApplicationTagUnsignedInteger getNumberOfDataElements() {
    return numberOfDataElements;
  }

  public List<BACnetDoorStatusTagged> getCarDoorStatus() {
    return carDoorStatus;
  }

  public BigInteger getZero() {
    Object o = 0L;
    if (o instanceof BigInteger) return (BigInteger) o;
    return BigInteger.valueOf(((Number) o).longValue());
  }

  @Override
  protected void serializeBACnetConstructedDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConstructedDataCarDoorStatus");

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BigInteger zero = getZero();
    writeBuffer.writeVirtual("zero", zero);

    // Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
    writeOptionalField(
        "numberOfDataElements",
        numberOfDataElements,
        new DataWriterComplexDefault<>(writeBuffer),
        ((arrayIndexArgument) != (null)) && ((arrayIndexArgument.getActualValue()) == (getZero())));

    // Array Field (carDoorStatus)
    writeComplexTypeArrayField("carDoorStatus", carDoorStatus, writeBuffer);

    writeBuffer.popContext("BACnetConstructedDataCarDoorStatus");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConstructedDataCarDoorStatus _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // A virtual field doesn't have any in- or output.

    // Optional Field (numberOfDataElements)
    if (numberOfDataElements != null) {
      lengthInBits += numberOfDataElements.getLengthInBits();
    }

    // Array field
    if (carDoorStatus != null) {
      for (Message element : carDoorStatus) {
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static BACnetConstructedDataBuilder staticParseBACnetConstructedDataBuilder(
      ReadBuffer readBuffer,
      Short tagNumber,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetConstructedDataCarDoorStatus");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    BigInteger zero = readVirtualField("zero", BigInteger.class, 0L);

    BACnetApplicationTagUnsignedInteger numberOfDataElements =
        readOptionalField(
            "numberOfDataElements",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer),
            ((arrayIndexArgument) != (null)) && ((arrayIndexArgument.getActualValue()) == (zero)));

    List<BACnetDoorStatusTagged> carDoorStatus =
        readTerminatedArrayField(
            "carDoorStatus",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetDoorStatusTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.APPLICATION_TAGS)),
                readBuffer),
            () ->
                ((boolean)
                    (org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper
                        .isBACnetConstructedDataClosingTag(readBuffer, false, tagNumber))));

    readBuffer.closeContext("BACnetConstructedDataCarDoorStatus");
    // Create the instance
    return new BACnetConstructedDataCarDoorStatusBuilderImpl(
        numberOfDataElements, carDoorStatus, tagNumber, arrayIndexArgument);
  }

  public static class BACnetConstructedDataCarDoorStatusBuilderImpl
      implements BACnetConstructedData.BACnetConstructedDataBuilder {
    private final BACnetApplicationTagUnsignedInteger numberOfDataElements;
    private final List<BACnetDoorStatusTagged> carDoorStatus;
    private final Short tagNumber;
    private final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

    public BACnetConstructedDataCarDoorStatusBuilderImpl(
        BACnetApplicationTagUnsignedInteger numberOfDataElements,
        List<BACnetDoorStatusTagged> carDoorStatus,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      this.numberOfDataElements = numberOfDataElements;
      this.carDoorStatus = carDoorStatus;
      this.tagNumber = tagNumber;
      this.arrayIndexArgument = arrayIndexArgument;
    }

    public BACnetConstructedDataCarDoorStatus build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      BACnetConstructedDataCarDoorStatus bACnetConstructedDataCarDoorStatus =
          new BACnetConstructedDataCarDoorStatus(
              openingTag,
              peekedTagHeader,
              closingTag,
              numberOfDataElements,
              carDoorStatus,
              tagNumber,
              arrayIndexArgument);
      return bACnetConstructedDataCarDoorStatus;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConstructedDataCarDoorStatus)) {
      return false;
    }
    BACnetConstructedDataCarDoorStatus that = (BACnetConstructedDataCarDoorStatus) o;
    return (getNumberOfDataElements() == that.getNumberOfDataElements())
        && (getCarDoorStatus() == that.getCarDoorStatus())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNumberOfDataElements(), getCarDoorStatus());
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