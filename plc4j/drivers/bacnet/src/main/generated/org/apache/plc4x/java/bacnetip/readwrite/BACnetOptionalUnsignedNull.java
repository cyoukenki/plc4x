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

public class BACnetOptionalUnsignedNull extends BACnetOptionalUnsigned implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagNull nullValue;

  public BACnetOptionalUnsignedNull(
      BACnetTagHeader peekedTagHeader, BACnetApplicationTagNull nullValue) {
    super(peekedTagHeader);
    this.nullValue = nullValue;
  }

  public BACnetApplicationTagNull getNullValue() {
    return nullValue;
  }

  @Override
  protected void serializeBACnetOptionalUnsignedChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetOptionalUnsignedNull");

    // Simple Field (nullValue)
    writeSimpleField("nullValue", nullValue, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetOptionalUnsignedNull");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetOptionalUnsignedNull _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nullValue)
    lengthInBits += nullValue.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetOptionalUnsignedBuilder staticParseBACnetOptionalUnsignedBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetOptionalUnsignedNull");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagNull nullValue =
        readSimpleField(
            "nullValue",
            new DataReaderComplexDefault<>(
                () -> (BACnetApplicationTagNull) BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetOptionalUnsignedNull");
    // Create the instance
    return new BACnetOptionalUnsignedNullBuilderImpl(nullValue);
  }

  public static class BACnetOptionalUnsignedNullBuilderImpl
      implements BACnetOptionalUnsigned.BACnetOptionalUnsignedBuilder {
    private final BACnetApplicationTagNull nullValue;

    public BACnetOptionalUnsignedNullBuilderImpl(BACnetApplicationTagNull nullValue) {
      this.nullValue = nullValue;
    }

    public BACnetOptionalUnsignedNull build(BACnetTagHeader peekedTagHeader) {
      BACnetOptionalUnsignedNull bACnetOptionalUnsignedNull =
          new BACnetOptionalUnsignedNull(peekedTagHeader, nullValue);
      return bACnetOptionalUnsignedNull;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetOptionalUnsignedNull)) {
      return false;
    }
    BACnetOptionalUnsignedNull that = (BACnetOptionalUnsignedNull) o;
    return (getNullValue() == that.getNullValue()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNullValue());
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
