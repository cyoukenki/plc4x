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

public class BACnetPropertyStatesNetworkNumberQuality extends BACnetPropertyStates
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetNetworkNumberQualityTagged networkNumberQuality;

  public BACnetPropertyStatesNetworkNumberQuality(
      BACnetTagHeader peekedTagHeader, BACnetNetworkNumberQualityTagged networkNumberQuality) {
    super(peekedTagHeader);
    this.networkNumberQuality = networkNumberQuality;
  }

  public BACnetNetworkNumberQualityTagged getNetworkNumberQuality() {
    return networkNumberQuality;
  }

  @Override
  protected void serializeBACnetPropertyStatesChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetPropertyStatesNetworkNumberQuality");

    // Simple Field (networkNumberQuality)
    writeSimpleField(
        "networkNumberQuality", networkNumberQuality, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetPropertyStatesNetworkNumberQuality");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetPropertyStatesNetworkNumberQuality _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (networkNumberQuality)
    lengthInBits += networkNumberQuality.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetPropertyStatesBuilder staticParseBACnetPropertyStatesBuilder(
      ReadBuffer readBuffer, Short peekedTagNumber) throws ParseException {
    readBuffer.pullContext("BACnetPropertyStatesNetworkNumberQuality");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetNetworkNumberQualityTagged networkNumberQuality =
        readSimpleField(
            "networkNumberQuality",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetNetworkNumberQualityTagged.staticParse(
                        readBuffer,
                        (short) (peekedTagNumber),
                        (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    readBuffer.closeContext("BACnetPropertyStatesNetworkNumberQuality");
    // Create the instance
    return new BACnetPropertyStatesNetworkNumberQualityBuilderImpl(networkNumberQuality);
  }

  public static class BACnetPropertyStatesNetworkNumberQualityBuilderImpl
      implements BACnetPropertyStates.BACnetPropertyStatesBuilder {
    private final BACnetNetworkNumberQualityTagged networkNumberQuality;

    public BACnetPropertyStatesNetworkNumberQualityBuilderImpl(
        BACnetNetworkNumberQualityTagged networkNumberQuality) {
      this.networkNumberQuality = networkNumberQuality;
    }

    public BACnetPropertyStatesNetworkNumberQuality build(BACnetTagHeader peekedTagHeader) {
      BACnetPropertyStatesNetworkNumberQuality bACnetPropertyStatesNetworkNumberQuality =
          new BACnetPropertyStatesNetworkNumberQuality(peekedTagHeader, networkNumberQuality);
      return bACnetPropertyStatesNetworkNumberQuality;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetPropertyStatesNetworkNumberQuality)) {
      return false;
    }
    BACnetPropertyStatesNetworkNumberQuality that = (BACnetPropertyStatesNetworkNumberQuality) o;
    return (getNetworkNumberQuality() == that.getNetworkNumberQuality())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNetworkNumberQuality());
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