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
package org.apache.plc4x.java.openprotocol.readwrite;

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

public class OpenProtocolMessageApplicationGenericUnsubscribeRev1
    extends OpenProtocolMessageApplicationGenericUnsubscribe implements Message {

  // Accessors for discriminator values.
  public Integer getRevision() {
    return (int) 1;
  }

  // Properties.
  protected final Mid subscriptionMid;
  protected final int extraDataRevision;
  protected final byte[] extraData;

  public OpenProtocolMessageApplicationGenericUnsubscribeRev1(
      Integer midRevision,
      Short noAckFlag,
      Integer targetStationId,
      Integer targetSpindleId,
      Integer sequenceNumber,
      Short numberOfMessageParts,
      Short messagePartNumber,
      Mid subscriptionMid,
      int extraDataRevision,
      byte[] extraData) {
    super(
        midRevision,
        noAckFlag,
        targetStationId,
        targetSpindleId,
        sequenceNumber,
        numberOfMessageParts,
        messagePartNumber);
    this.subscriptionMid = subscriptionMid;
    this.extraDataRevision = extraDataRevision;
    this.extraData = extraData;
  }

  public Mid getSubscriptionMid() {
    return subscriptionMid;
  }

  public int getExtraDataRevision() {
    return extraDataRevision;
  }

  public byte[] getExtraData() {
    return extraData;
  }

  @Override
  protected void serializeOpenProtocolMessageApplicationGenericUnsubscribeChild(
      WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("OpenProtocolMessageApplicationGenericUnsubscribeRev1");

    // Simple Field (subscriptionMid)
    writeSimpleEnumField(
        "subscriptionMid",
        "Mid",
        subscriptionMid,
        new DataWriterEnumDefault<>(Mid::getValue, Mid::name, writeUnsignedLong(writeBuffer, 32)),
        WithOption.WithEncoding("ASCII"));

    // Simple Field (extraDataRevision)
    writeSimpleField(
        "extraDataRevision",
        extraDataRevision,
        writeUnsignedInt(writeBuffer, 24),
        WithOption.WithEncoding("ASCII"));

    // Implicit Field (extraDataLength) (Used for parsing, but its value is not stored as it's
    // implicitly given by the objects content)
    int extraDataLength = (int) (COUNT(getExtraData()));
    writeImplicitField(
        "extraDataLength",
        extraDataLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithEncoding("ASCII"));

    // Array Field (extraData)
    writeByteArrayField("extraData", extraData, writeByteArray(writeBuffer, 8));

    writeBuffer.popContext("OpenProtocolMessageApplicationGenericUnsubscribeRev1");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    OpenProtocolMessageApplicationGenericUnsubscribeRev1 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (subscriptionMid)
    lengthInBits += 32;

    // Simple field (extraDataRevision)
    lengthInBits += 24;

    // Implicit Field (extraDataLength)
    lengthInBits += 16;

    // Array field
    if (extraData != null) {
      lengthInBits += 8 * extraData.length;
    }

    return lengthInBits;
  }

  public static OpenProtocolMessageApplicationGenericUnsubscribeBuilder
      staticParseOpenProtocolMessageApplicationGenericUnsubscribeBuilder(
          ReadBuffer readBuffer, Integer revision) throws ParseException {
    readBuffer.pullContext("OpenProtocolMessageApplicationGenericUnsubscribeRev1");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    Mid subscriptionMid =
        readEnumField(
            "subscriptionMid",
            "Mid",
            new DataReaderEnumDefault<>(Mid::enumForValue, readUnsignedLong(readBuffer, 32)),
            WithOption.WithEncoding("ASCII"));

    int extraDataRevision =
        readSimpleField(
            "extraDataRevision", readUnsignedInt(readBuffer, 24), WithOption.WithEncoding("ASCII"));

    int extraDataLength =
        readImplicitField(
            "extraDataLength", readUnsignedInt(readBuffer, 16), WithOption.WithEncoding("ASCII"));

    byte[] extraData = readBuffer.readByteArray("extraData", Math.toIntExact(extraDataLength));

    readBuffer.closeContext("OpenProtocolMessageApplicationGenericUnsubscribeRev1");
    // Create the instance
    return new OpenProtocolMessageApplicationGenericUnsubscribeRev1BuilderImpl(
        subscriptionMid, extraDataRevision, extraData);
  }

  public static class OpenProtocolMessageApplicationGenericUnsubscribeRev1BuilderImpl
      implements OpenProtocolMessageApplicationGenericUnsubscribe
          .OpenProtocolMessageApplicationGenericUnsubscribeBuilder {
    private final Mid subscriptionMid;
    private final int extraDataRevision;
    private final byte[] extraData;

    public OpenProtocolMessageApplicationGenericUnsubscribeRev1BuilderImpl(
        Mid subscriptionMid, int extraDataRevision, byte[] extraData) {
      this.subscriptionMid = subscriptionMid;
      this.extraDataRevision = extraDataRevision;
      this.extraData = extraData;
    }

    public OpenProtocolMessageApplicationGenericUnsubscribeRev1 build(
        Integer midRevision,
        Short noAckFlag,
        Integer targetStationId,
        Integer targetSpindleId,
        Integer sequenceNumber,
        Short numberOfMessageParts,
        Short messagePartNumber) {
      OpenProtocolMessageApplicationGenericUnsubscribeRev1
          openProtocolMessageApplicationGenericUnsubscribeRev1 =
              new OpenProtocolMessageApplicationGenericUnsubscribeRev1(
                  midRevision,
                  noAckFlag,
                  targetStationId,
                  targetSpindleId,
                  sequenceNumber,
                  numberOfMessageParts,
                  messagePartNumber,
                  subscriptionMid,
                  extraDataRevision,
                  extraData);
      return openProtocolMessageApplicationGenericUnsubscribeRev1;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof OpenProtocolMessageApplicationGenericUnsubscribeRev1)) {
      return false;
    }
    OpenProtocolMessageApplicationGenericUnsubscribeRev1 that =
        (OpenProtocolMessageApplicationGenericUnsubscribeRev1) o;
    return (getSubscriptionMid() == that.getSubscriptionMid())
        && (getExtraDataRevision() == that.getExtraDataRevision())
        && (getExtraData() == that.getExtraData())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getSubscriptionMid(), getExtraDataRevision(), getExtraData());
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