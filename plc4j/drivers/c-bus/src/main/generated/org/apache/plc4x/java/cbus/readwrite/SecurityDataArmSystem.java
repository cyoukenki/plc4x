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

public class SecurityDataArmSystem extends SecurityData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte armMode;

  public SecurityDataArmSystem(
      SecurityCommandTypeContainer commandTypeContainer, byte argument, byte armMode) {
    super(commandTypeContainer, argument);
    this.armMode = armMode;
  }

  public byte getArmMode() {
    return armMode;
  }

  public boolean getIsReserved() {
    return (boolean)
        (((getArmMode()) == (0x00))
            || ((((getArmMode()) >= (0x05)) && ((getArmMode()) <= (0xFE)))));
  }

  public boolean getIsArmToAwayMode() {
    return (boolean) ((getArmMode()) == (0x01));
  }

  public boolean getIsArmToNightMode() {
    return (boolean) ((getArmMode()) == (0x02));
  }

  public boolean getIsArmToDayMode() {
    return (boolean) ((getArmMode()) == (0x03));
  }

  public boolean getIsArmToVacationMode() {
    return (boolean) ((getArmMode()) == (0x04));
  }

  public boolean getIsArmToHighestLevelOfProtection() {
    return (boolean) ((getArmMode()) > (0xFE));
  }

  @Override
  protected void serializeSecurityDataChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SecurityDataArmSystem");

    // Simple Field (armMode)
    writeSimpleField("armMode", armMode, writeByte(writeBuffer, 8));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isReserved = getIsReserved();
    writeBuffer.writeVirtual("isReserved", isReserved);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isArmToAwayMode = getIsArmToAwayMode();
    writeBuffer.writeVirtual("isArmToAwayMode", isArmToAwayMode);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isArmToNightMode = getIsArmToNightMode();
    writeBuffer.writeVirtual("isArmToNightMode", isArmToNightMode);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isArmToDayMode = getIsArmToDayMode();
    writeBuffer.writeVirtual("isArmToDayMode", isArmToDayMode);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isArmToVacationMode = getIsArmToVacationMode();
    writeBuffer.writeVirtual("isArmToVacationMode", isArmToVacationMode);

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isArmToHighestLevelOfProtection = getIsArmToHighestLevelOfProtection();
    writeBuffer.writeVirtual("isArmToHighestLevelOfProtection", isArmToHighestLevelOfProtection);

    writeBuffer.popContext("SecurityDataArmSystem");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SecurityDataArmSystem _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (armMode)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    // A virtual field doesn't have any in- or output.

    return lengthInBits;
  }

  public static SecurityDataBuilder staticParseSecurityDataBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("SecurityDataArmSystem");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte armMode = readSimpleField("armMode", readByte(readBuffer, 8));
    boolean isReserved =
        readVirtualField(
            "isReserved",
            boolean.class,
            ((armMode) == (0x00)) || ((((armMode) >= (0x05)) && ((armMode) <= (0xFE)))));
    boolean isArmToAwayMode =
        readVirtualField("isArmToAwayMode", boolean.class, (armMode) == (0x01));
    boolean isArmToNightMode =
        readVirtualField("isArmToNightMode", boolean.class, (armMode) == (0x02));
    boolean isArmToDayMode = readVirtualField("isArmToDayMode", boolean.class, (armMode) == (0x03));
    boolean isArmToVacationMode =
        readVirtualField("isArmToVacationMode", boolean.class, (armMode) == (0x04));
    boolean isArmToHighestLevelOfProtection =
        readVirtualField("isArmToHighestLevelOfProtection", boolean.class, (armMode) > (0xFE));

    readBuffer.closeContext("SecurityDataArmSystem");
    // Create the instance
    return new SecurityDataArmSystemBuilderImpl(armMode);
  }

  public static class SecurityDataArmSystemBuilderImpl implements SecurityData.SecurityDataBuilder {
    private final byte armMode;

    public SecurityDataArmSystemBuilderImpl(byte armMode) {
      this.armMode = armMode;
    }

    public SecurityDataArmSystem build(
        SecurityCommandTypeContainer commandTypeContainer, byte argument) {
      SecurityDataArmSystem securityDataArmSystem =
          new SecurityDataArmSystem(commandTypeContainer, argument, armMode);
      return securityDataArmSystem;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SecurityDataArmSystem)) {
      return false;
    }
    SecurityDataArmSystem that = (SecurityDataArmSystem) o;
    return (getArmMode() == that.getArmMode()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getArmMode());
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