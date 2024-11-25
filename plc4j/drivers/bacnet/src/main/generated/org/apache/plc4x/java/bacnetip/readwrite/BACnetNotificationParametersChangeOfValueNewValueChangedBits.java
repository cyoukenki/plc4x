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

public class BACnetNotificationParametersChangeOfValueNewValueChangedBits
    extends BACnetNotificationParametersChangeOfValueNewValue implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagBitString changedBits;

  // Arguments.
  protected final Short tagNumber;

  public BACnetNotificationParametersChangeOfValueNewValueChangedBits(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetContextTagBitString changedBits,
      Short tagNumber) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber);
    this.changedBits = changedBits;
    this.tagNumber = tagNumber;
  }

  public BACnetContextTagBitString getChangedBits() {
    return changedBits;
  }

  @Override
  protected void serializeBACnetNotificationParametersChangeOfValueNewValueChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits");

    // Simple Field (changedBits)
    writeSimpleField("changedBits", changedBits, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetNotificationParametersChangeOfValueNewValueChangedBits _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (changedBits)
    lengthInBits += changedBits.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetNotificationParametersChangeOfValueNewValueBuilder
      staticParseBACnetNotificationParametersChangeOfValueNewValueBuilder(
          ReadBuffer readBuffer, Short peekedTagNumber, Short tagNumber) throws ParseException {
    readBuffer.pullContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagBitString changedBits =
        readSimpleField(
            "changedBits",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagBitString)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (0), (BACnetDataType) (BACnetDataType.BIT_STRING)),
                readBuffer));

    readBuffer.closeContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits");
    // Create the instance
    return new BACnetNotificationParametersChangeOfValueNewValueChangedBitsBuilderImpl(
        changedBits, tagNumber);
  }

  public static class BACnetNotificationParametersChangeOfValueNewValueChangedBitsBuilderImpl
      implements BACnetNotificationParametersChangeOfValueNewValue
          .BACnetNotificationParametersChangeOfValueNewValueBuilder {
    private final BACnetContextTagBitString changedBits;
    private final Short tagNumber;

    public BACnetNotificationParametersChangeOfValueNewValueChangedBitsBuilderImpl(
        BACnetContextTagBitString changedBits, Short tagNumber) {
      this.changedBits = changedBits;
      this.tagNumber = tagNumber;
    }

    public BACnetNotificationParametersChangeOfValueNewValueChangedBits build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber) {
      BACnetNotificationParametersChangeOfValueNewValueChangedBits
          bACnetNotificationParametersChangeOfValueNewValueChangedBits =
              new BACnetNotificationParametersChangeOfValueNewValueChangedBits(
                  openingTag, peekedTagHeader, closingTag, changedBits, tagNumber);
      return bACnetNotificationParametersChangeOfValueNewValueChangedBits;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetNotificationParametersChangeOfValueNewValueChangedBits)) {
      return false;
    }
    BACnetNotificationParametersChangeOfValueNewValueChangedBits that =
        (BACnetNotificationParametersChangeOfValueNewValueChangedBits) o;
    return (getChangedBits() == that.getChangedBits()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getChangedBits());
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