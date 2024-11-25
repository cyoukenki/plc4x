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
package org.apache.plc4x.java.s7.readwrite;

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

public class S7ParameterUserData extends S7Parameter implements Message {

  // Accessors for discriminator values.
  public Short getParameterType() {
    return (short) 0x00;
  }

  public Short getMessageType() {
    return (short) 0x07;
  }

  // Properties.
  protected final List<S7ParameterUserDataItem> items;

  public S7ParameterUserData(List<S7ParameterUserDataItem> items) {
    super();
    this.items = items;
  }

  public List<S7ParameterUserDataItem> getItems() {
    return items;
  }

  @Override
  protected void serializeS7ParameterChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("S7ParameterUserData");

    // Implicit Field (numItems) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    short numItems = (short) (COUNT(getItems()));
    writeImplicitField("numItems", numItems, writeUnsignedShort(writeBuffer, 8));

    // Array Field (items)
    writeComplexTypeArrayField("items", items, writeBuffer);

    writeBuffer.popContext("S7ParameterUserData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7ParameterUserData _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Implicit Field (numItems)
    lengthInBits += 8;

    // Array field
    if (items != null) {
      int i = 0;
      for (S7ParameterUserDataItem element : items) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= items.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static S7ParameterBuilder staticParseS7ParameterBuilder(
      ReadBuffer readBuffer, Short messageType) throws ParseException {
    readBuffer.pullContext("S7ParameterUserData");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short numItems = readImplicitField("numItems", readUnsignedShort(readBuffer, 8));

    List<S7ParameterUserDataItem> items =
        readCountArrayField(
            "items",
            new DataReaderComplexDefault<>(
                () -> S7ParameterUserDataItem.staticParse(readBuffer), readBuffer),
            numItems);

    readBuffer.closeContext("S7ParameterUserData");
    // Create the instance
    return new S7ParameterUserDataBuilderImpl(items);
  }

  public static class S7ParameterUserDataBuilderImpl implements S7Parameter.S7ParameterBuilder {
    private final List<S7ParameterUserDataItem> items;

    public S7ParameterUserDataBuilderImpl(List<S7ParameterUserDataItem> items) {
      this.items = items;
    }

    public S7ParameterUserData build() {
      S7ParameterUserData s7ParameterUserData = new S7ParameterUserData(items);
      return s7ParameterUserData;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7ParameterUserData)) {
      return false;
    }
    S7ParameterUserData that = (S7ParameterUserData) o;
    return (getItems() == that.getItems()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getItems());
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