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

public class BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
    extends BACnetEventParameterChangeOfValueCivCriteria implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetContextTagReal referencedPropertyIncrement;

  // Arguments.
  protected final Short tagNumber;

  public BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetContextTagReal referencedPropertyIncrement,
      Short tagNumber) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber);
    this.referencedPropertyIncrement = referencedPropertyIncrement;
    this.tagNumber = tagNumber;
  }

  public BACnetContextTagReal getReferencedPropertyIncrement() {
    return referencedPropertyIncrement;
  }

  @Override
  protected void serializeBACnetEventParameterChangeOfValueCivCriteriaChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext(
        "BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement");

    // Simple Field (referencedPropertyIncrement)
    writeSimpleField(
        "referencedPropertyIncrement",
        referencedPropertyIncrement,
        new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext(
        "BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (referencedPropertyIncrement)
    lengthInBits += referencedPropertyIncrement.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetEventParameterChangeOfValueCivCriteriaBuilder
      staticParseBACnetEventParameterChangeOfValueCivCriteriaBuilder(
          ReadBuffer readBuffer, Short tagNumber) throws ParseException {
    readBuffer.pullContext(
        "BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetContextTagReal referencedPropertyIncrement =
        readSimpleField(
            "referencedPropertyIncrement",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagReal)
                        BACnetContextTag.staticParse(
                            readBuffer, (short) (1), (BACnetDataType) (BACnetDataType.REAL)),
                readBuffer));

    readBuffer.closeContext(
        "BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement");
    // Create the instance
    return new BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilderImpl(
        referencedPropertyIncrement, tagNumber);
  }

  public static
  class BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilderImpl
      implements BACnetEventParameterChangeOfValueCivCriteria
          .BACnetEventParameterChangeOfValueCivCriteriaBuilder {
    private final BACnetContextTagReal referencedPropertyIncrement;
    private final Short tagNumber;

    public BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrementBuilderImpl(
        BACnetContextTagReal referencedPropertyIncrement, Short tagNumber) {
      this.referencedPropertyIncrement = referencedPropertyIncrement;
      this.tagNumber = tagNumber;
    }

    public BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber) {
      BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement
          bACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement =
              new BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement(
                  openingTag, peekedTagHeader, closingTag, referencedPropertyIncrement, tagNumber);
      return bACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement)) {
      return false;
    }
    BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement that =
        (BACnetEventParameterChangeOfValueCivCriteriaReferencedPropertyIncrement) o;
    return (getReferencedPropertyIncrement() == that.getReferencedPropertyIncrement())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getReferencedPropertyIncrement());
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
