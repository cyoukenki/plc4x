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

public class BACnetEventParameterSignedOutOfRange extends BACnetEventParameter implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetContextTagUnsignedInteger timeDelay;
  protected final BACnetContextTagSignedInteger lowLimit;
  protected final BACnetContextTagSignedInteger highLimit;
  protected final BACnetContextTagUnsignedInteger deadband;
  protected final BACnetClosingTag closingTag;

  public BACnetEventParameterSignedOutOfRange(
      BACnetTagHeader peekedTagHeader,
      BACnetOpeningTag openingTag,
      BACnetContextTagUnsignedInteger timeDelay,
      BACnetContextTagSignedInteger lowLimit,
      BACnetContextTagSignedInteger highLimit,
      BACnetContextTagUnsignedInteger deadband,
      BACnetClosingTag closingTag) {
    super(peekedTagHeader);
    this.openingTag = openingTag;
    this.timeDelay = timeDelay;
    this.lowLimit = lowLimit;
    this.highLimit = highLimit;
    this.deadband = deadband;
    this.closingTag = closingTag;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetContextTagUnsignedInteger getTimeDelay() {
    return timeDelay;
  }

  public BACnetContextTagSignedInteger getLowLimit() {
    return lowLimit;
  }

  public BACnetContextTagSignedInteger getHighLimit() {
    return highLimit;
  }

  public BACnetContextTagUnsignedInteger getDeadband() {
    return deadband;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  @Override
  protected void serializeBACnetEventParameterChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetEventParameterSignedOutOfRange");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (timeDelay)
    writeSimpleField("timeDelay", timeDelay, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (lowLimit)
    writeSimpleField("lowLimit", lowLimit, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (highLimit)
    writeSimpleField("highLimit", highLimit, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (deadband)
    writeSimpleField("deadband", deadband, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetEventParameterSignedOutOfRange");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetEventParameterSignedOutOfRange _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Simple field (timeDelay)
    lengthInBits += timeDelay.getLengthInBits();

    // Simple field (lowLimit)
    lengthInBits += lowLimit.getLengthInBits();

    // Simple field (highLimit)
    lengthInBits += highLimit.getLengthInBits();

    // Simple field (deadband)
    lengthInBits += deadband.getLengthInBits();

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetEventParameterBuilder staticParseBACnetEventParameterBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetEventParameterSignedOutOfRange");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (15)), readBuffer));

    BACnetContextTagUnsignedInteger timeDelay =
        readSimpleField(
            "timeDelay",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagSignedInteger lowLimit =
        readSimpleField(
            "lowLimit",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagSignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (1),
                            (BACnetDataType) (BACnetDataType.SIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagSignedInteger highLimit =
        readSimpleField(
            "highLimit",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagSignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (2),
                            (BACnetDataType) (BACnetDataType.SIGNED_INTEGER)),
                readBuffer));

    BACnetContextTagUnsignedInteger deadband =
        readSimpleField(
            "deadband",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (3),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (15)), readBuffer));

    readBuffer.closeContext("BACnetEventParameterSignedOutOfRange");
    // Create the instance
    return new BACnetEventParameterSignedOutOfRangeBuilderImpl(
        openingTag, timeDelay, lowLimit, highLimit, deadband, closingTag);
  }

  public static class BACnetEventParameterSignedOutOfRangeBuilderImpl
      implements BACnetEventParameter.BACnetEventParameterBuilder {
    private final BACnetOpeningTag openingTag;
    private final BACnetContextTagUnsignedInteger timeDelay;
    private final BACnetContextTagSignedInteger lowLimit;
    private final BACnetContextTagSignedInteger highLimit;
    private final BACnetContextTagUnsignedInteger deadband;
    private final BACnetClosingTag closingTag;

    public BACnetEventParameterSignedOutOfRangeBuilderImpl(
        BACnetOpeningTag openingTag,
        BACnetContextTagUnsignedInteger timeDelay,
        BACnetContextTagSignedInteger lowLimit,
        BACnetContextTagSignedInteger highLimit,
        BACnetContextTagUnsignedInteger deadband,
        BACnetClosingTag closingTag) {
      this.openingTag = openingTag;
      this.timeDelay = timeDelay;
      this.lowLimit = lowLimit;
      this.highLimit = highLimit;
      this.deadband = deadband;
      this.closingTag = closingTag;
    }

    public BACnetEventParameterSignedOutOfRange build(BACnetTagHeader peekedTagHeader) {
      BACnetEventParameterSignedOutOfRange bACnetEventParameterSignedOutOfRange =
          new BACnetEventParameterSignedOutOfRange(
              peekedTagHeader, openingTag, timeDelay, lowLimit, highLimit, deadband, closingTag);
      return bACnetEventParameterSignedOutOfRange;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetEventParameterSignedOutOfRange)) {
      return false;
    }
    BACnetEventParameterSignedOutOfRange that = (BACnetEventParameterSignedOutOfRange) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getTimeDelay() == that.getTimeDelay())
        && (getLowLimit() == that.getLowLimit())
        && (getHighLimit() == that.getHighLimit())
        && (getDeadband() == that.getDeadband())
        && (getClosingTag() == that.getClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getOpeningTag(),
        getTimeDelay(),
        getLowLimit(),
        getHighLimit(),
        getDeadband(),
        getClosingTag());
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