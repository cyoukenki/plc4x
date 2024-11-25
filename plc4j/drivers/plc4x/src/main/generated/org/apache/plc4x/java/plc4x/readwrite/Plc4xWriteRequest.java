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
package org.apache.plc4x.java.plc4x.readwrite;

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

public class Plc4xWriteRequest extends Plc4xMessage implements Message {

  // Accessors for discriminator values.
  public Plc4xRequestType getRequestType() {
    return Plc4xRequestType.WRITE_REQUEST;
  }

  // Properties.
  protected final int connectionId;
  protected final List<Plc4xTagValueRequest> tags;

  public Plc4xWriteRequest(int requestId, int connectionId, List<Plc4xTagValueRequest> tags) {
    super(requestId);
    this.connectionId = connectionId;
    this.tags = tags;
  }

  public int getConnectionId() {
    return connectionId;
  }

  public List<Plc4xTagValueRequest> getTags() {
    return tags;
  }

  @Override
  protected void serializePlc4xMessageChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("Plc4xWriteRequest");

    // Simple Field (connectionId)
    writeSimpleField(
        "connectionId",
        connectionId,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Implicit Field (numTags) (Used for parsing, but its value is not stored as it's implicitly
    // given by the objects content)
    short numTags = (short) (COUNT(getTags()));
    writeImplicitField(
        "numTags",
        numTags,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Array Field (tags)
    writeComplexTypeArrayField(
        "tags", tags, writeBuffer, WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("Plc4xWriteRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    Plc4xWriteRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (connectionId)
    lengthInBits += 16;

    // Implicit Field (numTags)
    lengthInBits += 8;

    // Array field
    if (tags != null) {
      int i = 0;
      for (Plc4xTagValueRequest element : tags) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= tags.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static Plc4xMessageBuilder staticParsePlc4xMessageBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("Plc4xWriteRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int connectionId =
        readSimpleField(
            "connectionId",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short numTags =
        readImplicitField(
            "numTags",
            readUnsignedShort(readBuffer, 8),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    List<Plc4xTagValueRequest> tags =
        readCountArrayField(
            "tags",
            new DataReaderComplexDefault<>(
                () -> Plc4xTagValueRequest.staticParse(readBuffer), readBuffer),
            numTags,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("Plc4xWriteRequest");
    // Create the instance
    return new Plc4xWriteRequestBuilderImpl(connectionId, tags);
  }

  public static class Plc4xWriteRequestBuilderImpl implements Plc4xMessage.Plc4xMessageBuilder {
    private final int connectionId;
    private final List<Plc4xTagValueRequest> tags;

    public Plc4xWriteRequestBuilderImpl(int connectionId, List<Plc4xTagValueRequest> tags) {
      this.connectionId = connectionId;
      this.tags = tags;
    }

    public Plc4xWriteRequest build(int requestId) {
      Plc4xWriteRequest plc4xWriteRequest = new Plc4xWriteRequest(requestId, connectionId, tags);
      return plc4xWriteRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof Plc4xWriteRequest)) {
      return false;
    }
    Plc4xWriteRequest that = (Plc4xWriteRequest) o;
    return (getConnectionId() == that.getConnectionId())
        && (getTags() == that.getTags())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getConnectionId(), getTags());
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