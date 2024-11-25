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

public class BACnetProcessIdSelectionValue extends BACnetProcessIdSelection implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetApplicationTagUnsignedInteger processIdentifier;

  public BACnetProcessIdSelectionValue(
      BACnetTagHeader peekedTagHeader, BACnetApplicationTagUnsignedInteger processIdentifier) {
    super(peekedTagHeader);
    this.processIdentifier = processIdentifier;
  }

  public BACnetApplicationTagUnsignedInteger getProcessIdentifier() {
    return processIdentifier;
  }

  @Override
  protected void serializeBACnetProcessIdSelectionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BACnetProcessIdSelectionValue");

    // Simple Field (processIdentifier)
    writeSimpleField(
        "processIdentifier", processIdentifier, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetProcessIdSelectionValue");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetProcessIdSelectionValue _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (processIdentifier)
    lengthInBits += processIdentifier.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetProcessIdSelectionBuilder staticParseBACnetProcessIdSelectionBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetProcessIdSelectionValue");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetApplicationTagUnsignedInteger processIdentifier =
        readSimpleField(
            "processIdentifier",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetApplicationTagUnsignedInteger)
                        BACnetApplicationTag.staticParse(readBuffer),
                readBuffer));

    readBuffer.closeContext("BACnetProcessIdSelectionValue");
    // Create the instance
    return new BACnetProcessIdSelectionValueBuilderImpl(processIdentifier);
  }

  public static class BACnetProcessIdSelectionValueBuilderImpl
      implements BACnetProcessIdSelection.BACnetProcessIdSelectionBuilder {
    private final BACnetApplicationTagUnsignedInteger processIdentifier;

    public BACnetProcessIdSelectionValueBuilderImpl(
        BACnetApplicationTagUnsignedInteger processIdentifier) {
      this.processIdentifier = processIdentifier;
    }

    public BACnetProcessIdSelectionValue build(BACnetTagHeader peekedTagHeader) {
      BACnetProcessIdSelectionValue bACnetProcessIdSelectionValue =
          new BACnetProcessIdSelectionValue(peekedTagHeader, processIdentifier);
      return bACnetProcessIdSelectionValue;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetProcessIdSelectionValue)) {
      return false;
    }
    BACnetProcessIdSelectionValue that = (BACnetProcessIdSelectionValue) o;
    return (getProcessIdentifier() == that.getProcessIdentifier()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getProcessIdentifier());
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