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

public class BACnetConstructedDataBACnetIPv6Mode extends BACnetConstructedData implements Message {

  // Accessors for discriminator values.
  public BACnetObjectType getObjectTypeArgument() {
    return null;
  }

  public BACnetPropertyIdentifier getPropertyIdentifierArgument() {
    return BACnetPropertyIdentifier.BACNET_IPV6_MODE;
  }

  // Properties.
  protected final BACnetIPModeTagged bacnetIpv6Mode;

  // Arguments.
  protected final Short tagNumber;
  protected final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

  public BACnetConstructedDataBACnetIPv6Mode(
      BACnetOpeningTag openingTag,
      BACnetTagHeader peekedTagHeader,
      BACnetClosingTag closingTag,
      BACnetIPModeTagged bacnetIpv6Mode,
      Short tagNumber,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
    super(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument);
    this.bacnetIpv6Mode = bacnetIpv6Mode;
    this.tagNumber = tagNumber;
    this.arrayIndexArgument = arrayIndexArgument;
  }

  public BACnetIPModeTagged getBacnetIpv6Mode() {
    return bacnetIpv6Mode;
  }

  public BACnetIPModeTagged getActualValue() {
    return (BACnetIPModeTagged) (getBacnetIpv6Mode());
  }

  @Override
  protected void serializeBACnetConstructedDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetConstructedDataBACnetIPv6Mode");

    // Simple Field (bacnetIpv6Mode)
    writeSimpleField("bacnetIpv6Mode", bacnetIpv6Mode, new DataWriterComplexDefault<>(writeBuffer));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    BACnetIPModeTagged actualValue = getActualValue();
    writeBuffer.writeVirtual("actualValue", actualValue);

    writeBuffer.popContext("BACnetConstructedDataBACnetIPv6Mode");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetConstructedDataBACnetIPv6Mode _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (bacnetIpv6Mode)
    lengthInBits += bacnetIpv6Mode.getLengthInBits();

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static BACnetConstructedDataBuilder staticParseBACnetConstructedDataBuilder(
      ReadBuffer readBuffer,
      Short tagNumber,
      BACnetObjectType objectTypeArgument,
      BACnetPropertyIdentifier propertyIdentifierArgument,
      BACnetTagPayloadUnsignedInteger arrayIndexArgument)
      throws ParseException {
    readBuffer.pullContext("BACnetConstructedDataBACnetIPv6Mode");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetIPModeTagged bacnetIpv6Mode =
        readSimpleField(
            "bacnetIpv6Mode",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetIPModeTagged.staticParse(
                        readBuffer, (short) (0), (TagClass) (TagClass.APPLICATION_TAGS)),
                readBuffer));
    BACnetIPModeTagged actualValue =
        readVirtualField("actualValue", BACnetIPModeTagged.class, bacnetIpv6Mode);

    readBuffer.closeContext("BACnetConstructedDataBACnetIPv6Mode");
    // Create the instance
    return new BACnetConstructedDataBACnetIPv6ModeBuilderImpl(
        bacnetIpv6Mode, tagNumber, arrayIndexArgument);
  }

  public static class BACnetConstructedDataBACnetIPv6ModeBuilderImpl
      implements BACnetConstructedData.BACnetConstructedDataBuilder {
    private final BACnetIPModeTagged bacnetIpv6Mode;
    private final Short tagNumber;
    private final BACnetTagPayloadUnsignedInteger arrayIndexArgument;

    public BACnetConstructedDataBACnetIPv6ModeBuilderImpl(
        BACnetIPModeTagged bacnetIpv6Mode,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      this.bacnetIpv6Mode = bacnetIpv6Mode;
      this.tagNumber = tagNumber;
      this.arrayIndexArgument = arrayIndexArgument;
    }

    public BACnetConstructedDataBACnetIPv6Mode build(
        BACnetOpeningTag openingTag,
        BACnetTagHeader peekedTagHeader,
        BACnetClosingTag closingTag,
        Short tagNumber,
        BACnetTagPayloadUnsignedInteger arrayIndexArgument) {
      BACnetConstructedDataBACnetIPv6Mode bACnetConstructedDataBACnetIPv6Mode =
          new BACnetConstructedDataBACnetIPv6Mode(
              openingTag,
              peekedTagHeader,
              closingTag,
              bacnetIpv6Mode,
              tagNumber,
              arrayIndexArgument);
      return bACnetConstructedDataBACnetIPv6Mode;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetConstructedDataBACnetIPv6Mode)) {
      return false;
    }
    BACnetConstructedDataBACnetIPv6Mode that = (BACnetConstructedDataBACnetIPv6Mode) o;
    return (getBacnetIpv6Mode() == that.getBacnetIpv6Mode()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getBacnetIpv6Mode());
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
