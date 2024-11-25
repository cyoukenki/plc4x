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
package org.apache.plc4x.java.canopen.readwrite;

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

public class SDOInitiateUploadResponse extends SDOResponse implements Message {

  // Accessors for discriminator values.
  public SDOResponseCommand getCommand() {
    return SDOResponseCommand.INITIATE_UPLOAD;
  }

  // Properties.
  protected final boolean expedited;
  protected final boolean indicated;
  protected final IndexAddress address;
  protected final SDOInitiateUploadResponsePayload payload;

  public SDOInitiateUploadResponse(
      boolean expedited,
      boolean indicated,
      IndexAddress address,
      SDOInitiateUploadResponsePayload payload) {
    super();
    this.expedited = expedited;
    this.indicated = indicated;
    this.address = address;
    this.payload = payload;
  }

  public boolean getExpedited() {
    return expedited;
  }

  public boolean getIndicated() {
    return indicated;
  }

  public IndexAddress getAddress() {
    return address;
  }

  public SDOInitiateUploadResponsePayload getPayload() {
    return payload;
  }

  @Override
  protected void serializeSDOResponseChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SDOInitiateUploadResponse");

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 1));

    // Implicit Field (size) (Used for parsing, but its value is not stored as it's implicitly given
    // by the objects content)
    byte size =
        (byte)
            (org.apache.plc4x.java.canopen.readwrite.utils.StaticHelper.count(
                getExpedited(), getIndicated(), getPayload()));
    writeImplicitField("size", size, writeUnsignedByte(writeBuffer, 2));

    // Simple Field (expedited)
    writeSimpleField("expedited", expedited, writeBoolean(writeBuffer));

    // Simple Field (indicated)
    writeSimpleField("indicated", indicated, writeBoolean(writeBuffer));

    // Simple Field (address)
    writeSimpleField("address", address, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (payload)
    writeSimpleField("payload", payload, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("SDOInitiateUploadResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SDOInitiateUploadResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Reserved Field (reserved)
    lengthInBits += 1;

    // Implicit Field (size)
    lengthInBits += 2;

    // Simple field (expedited)
    lengthInBits += 1;

    // Simple field (indicated)
    lengthInBits += 1;

    // Simple field (address)
    lengthInBits += address.getLengthInBits();

    // Simple field (payload)
    lengthInBits += payload.getLengthInBits();

    return lengthInBits;
  }

  public static SDOResponseBuilder staticParseSDOResponseBuilder(
      ReadBuffer readBuffer, SDOResponseCommand command) throws ParseException {
    readBuffer.pullContext("SDOInitiateUploadResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 1), (byte) 0x00);

    byte size = readImplicitField("size", readUnsignedByte(readBuffer, 2));

    boolean expedited = readSimpleField("expedited", readBoolean(readBuffer));

    boolean indicated = readSimpleField("indicated", readBoolean(readBuffer));

    IndexAddress address =
        readSimpleField(
            "address",
            new DataReaderComplexDefault<>(() -> IndexAddress.staticParse(readBuffer), readBuffer));

    SDOInitiateUploadResponsePayload payload =
        readSimpleField(
            "payload",
            new DataReaderComplexDefault<>(
                () ->
                    SDOInitiateUploadResponsePayload.staticParse(
                        readBuffer, (boolean) (expedited), (boolean) (indicated), (byte) (size)),
                readBuffer));

    readBuffer.closeContext("SDOInitiateUploadResponse");
    // Create the instance
    return new SDOInitiateUploadResponseBuilderImpl(expedited, indicated, address, payload);
  }

  public static class SDOInitiateUploadResponseBuilderImpl
      implements SDOResponse.SDOResponseBuilder {
    private final boolean expedited;
    private final boolean indicated;
    private final IndexAddress address;
    private final SDOInitiateUploadResponsePayload payload;

    public SDOInitiateUploadResponseBuilderImpl(
        boolean expedited,
        boolean indicated,
        IndexAddress address,
        SDOInitiateUploadResponsePayload payload) {
      this.expedited = expedited;
      this.indicated = indicated;
      this.address = address;
      this.payload = payload;
    }

    public SDOInitiateUploadResponse build() {
      SDOInitiateUploadResponse sDOInitiateUploadResponse =
          new SDOInitiateUploadResponse(expedited, indicated, address, payload);
      return sDOInitiateUploadResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SDOInitiateUploadResponse)) {
      return false;
    }
    SDOInitiateUploadResponse that = (SDOInitiateUploadResponse) o;
    return (getExpedited() == that.getExpedited())
        && (getIndicated() == that.getIndicated())
        && (getAddress() == that.getAddress())
        && (getPayload() == that.getPayload())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getExpedited(), getIndicated(), getAddress(), getPayload());
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