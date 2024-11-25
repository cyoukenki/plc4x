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

public class S7ParameterWriteVarResponse extends S7Parameter implements Message {

  // Accessors for discriminator values.
  public Short getParameterType() {
    return (short) 0x05;
  }

  public Short getMessageType() {
    return (short) 0x03;
  }

  // Properties.
  protected final short numItems;

  public S7ParameterWriteVarResponse(short numItems) {
    super();
    this.numItems = numItems;
  }

  public short getNumItems() {
    return numItems;
  }

  @Override
  protected void serializeS7ParameterChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("S7ParameterWriteVarResponse");

    // Simple Field (numItems)
    writeSimpleField("numItems", numItems, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("S7ParameterWriteVarResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7ParameterWriteVarResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (numItems)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static S7ParameterBuilder staticParseS7ParameterBuilder(
      ReadBuffer readBuffer, Short messageType) throws ParseException {
    readBuffer.pullContext("S7ParameterWriteVarResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short numItems = readSimpleField("numItems", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("S7ParameterWriteVarResponse");
    // Create the instance
    return new S7ParameterWriteVarResponseBuilderImpl(numItems);
  }

  public static class S7ParameterWriteVarResponseBuilderImpl
      implements S7Parameter.S7ParameterBuilder {
    private final short numItems;

    public S7ParameterWriteVarResponseBuilderImpl(short numItems) {
      this.numItems = numItems;
    }

    public S7ParameterWriteVarResponse build() {
      S7ParameterWriteVarResponse s7ParameterWriteVarResponse =
          new S7ParameterWriteVarResponse(numItems);
      return s7ParameterWriteVarResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7ParameterWriteVarResponse)) {
      return false;
    }
    S7ParameterWriteVarResponse that = (S7ParameterWriteVarResponse) o;
    return (getNumItems() == that.getNumItems()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNumItems());
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
