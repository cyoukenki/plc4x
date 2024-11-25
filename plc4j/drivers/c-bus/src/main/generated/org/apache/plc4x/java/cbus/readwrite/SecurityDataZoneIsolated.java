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
package org.apache.plc4x.java.cbus.readwrite;

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

public class SecurityDataZoneIsolated extends SecurityData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final short zoneNumber;

  public SecurityDataZoneIsolated(
      SecurityCommandTypeContainer commandTypeContainer, byte argument, short zoneNumber) {
    super(commandTypeContainer, argument);
    this.zoneNumber = zoneNumber;
  }

  public short getZoneNumber() {
    return zoneNumber;
  }

  @Override
  protected void serializeSecurityDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SecurityDataZoneIsolated");

    // Simple Field (zoneNumber)
    writeSimpleField("zoneNumber", zoneNumber, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("SecurityDataZoneIsolated");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SecurityDataZoneIsolated _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (zoneNumber)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static SecurityDataBuilder staticParseSecurityDataBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("SecurityDataZoneIsolated");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short zoneNumber = readSimpleField("zoneNumber", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("SecurityDataZoneIsolated");
    // Create the instance
    return new SecurityDataZoneIsolatedBuilderImpl(zoneNumber);
  }

  public static class SecurityDataZoneIsolatedBuilderImpl
      implements SecurityData.SecurityDataBuilder {
    private final short zoneNumber;

    public SecurityDataZoneIsolatedBuilderImpl(short zoneNumber) {
      this.zoneNumber = zoneNumber;
    }

    public SecurityDataZoneIsolated build(
        SecurityCommandTypeContainer commandTypeContainer, byte argument) {
      SecurityDataZoneIsolated securityDataZoneIsolated =
          new SecurityDataZoneIsolated(commandTypeContainer, argument, zoneNumber);
      return securityDataZoneIsolated;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SecurityDataZoneIsolated)) {
      return false;
    }
    SecurityDataZoneIsolated that = (SecurityDataZoneIsolated) o;
    return (getZoneNumber() == that.getZoneNumber()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getZoneNumber());
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