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
package org.apache.plc4x.java.opcua.readwrite;

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

public class CreateSubscriptionRequest extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "787";
  }

  // Properties.
  protected final ExtensionObjectDefinition requestHeader;
  protected final double requestedPublishingInterval;
  protected final long requestedLifetimeCount;
  protected final long requestedMaxKeepAliveCount;
  protected final long maxNotificationsPerPublish;
  protected final boolean publishingEnabled;
  protected final short priority;

  public CreateSubscriptionRequest(
      ExtensionObjectDefinition requestHeader,
      double requestedPublishingInterval,
      long requestedLifetimeCount,
      long requestedMaxKeepAliveCount,
      long maxNotificationsPerPublish,
      boolean publishingEnabled,
      short priority) {
    super();
    this.requestHeader = requestHeader;
    this.requestedPublishingInterval = requestedPublishingInterval;
    this.requestedLifetimeCount = requestedLifetimeCount;
    this.requestedMaxKeepAliveCount = requestedMaxKeepAliveCount;
    this.maxNotificationsPerPublish = maxNotificationsPerPublish;
    this.publishingEnabled = publishingEnabled;
    this.priority = priority;
  }

  public ExtensionObjectDefinition getRequestHeader() {
    return requestHeader;
  }

  public double getRequestedPublishingInterval() {
    return requestedPublishingInterval;
  }

  public long getRequestedLifetimeCount() {
    return requestedLifetimeCount;
  }

  public long getRequestedMaxKeepAliveCount() {
    return requestedMaxKeepAliveCount;
  }

  public long getMaxNotificationsPerPublish() {
    return maxNotificationsPerPublish;
  }

  public boolean getPublishingEnabled() {
    return publishingEnabled;
  }

  public short getPriority() {
    return priority;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CreateSubscriptionRequest");

    // Simple Field (requestHeader)
    writeSimpleField("requestHeader", requestHeader, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (requestedPublishingInterval)
    writeSimpleField(
        "requestedPublishingInterval", requestedPublishingInterval, writeDouble(writeBuffer, 64));

    // Simple Field (requestedLifetimeCount)
    writeSimpleField(
        "requestedLifetimeCount", requestedLifetimeCount, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (requestedMaxKeepAliveCount)
    writeSimpleField(
        "requestedMaxKeepAliveCount",
        requestedMaxKeepAliveCount,
        writeUnsignedLong(writeBuffer, 32));

    // Simple Field (maxNotificationsPerPublish)
    writeSimpleField(
        "maxNotificationsPerPublish",
        maxNotificationsPerPublish,
        writeUnsignedLong(writeBuffer, 32));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (publishingEnabled)
    writeSimpleField("publishingEnabled", publishingEnabled, writeBoolean(writeBuffer));

    // Simple Field (priority)
    writeSimpleField("priority", priority, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("CreateSubscriptionRequest");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CreateSubscriptionRequest _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (requestHeader)
    lengthInBits += requestHeader.getLengthInBits();

    // Simple field (requestedPublishingInterval)
    lengthInBits += 64;

    // Simple field (requestedLifetimeCount)
    lengthInBits += 32;

    // Simple field (requestedMaxKeepAliveCount)
    lengthInBits += 32;

    // Simple field (maxNotificationsPerPublish)
    lengthInBits += 32;

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (publishingEnabled)
    lengthInBits += 1;

    // Simple field (priority)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("CreateSubscriptionRequest");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    ExtensionObjectDefinition requestHeader =
        readSimpleField(
            "requestHeader",
            new DataReaderComplexDefault<>(
                () -> ExtensionObjectDefinition.staticParse(readBuffer, (String) ("391")),
                readBuffer));

    double requestedPublishingInterval =
        readSimpleField("requestedPublishingInterval", readDouble(readBuffer, 64));

    long requestedLifetimeCount =
        readSimpleField("requestedLifetimeCount", readUnsignedLong(readBuffer, 32));

    long requestedMaxKeepAliveCount =
        readSimpleField("requestedMaxKeepAliveCount", readUnsignedLong(readBuffer, 32));

    long maxNotificationsPerPublish =
        readSimpleField("maxNotificationsPerPublish", readUnsignedLong(readBuffer, 32));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean publishingEnabled = readSimpleField("publishingEnabled", readBoolean(readBuffer));

    short priority = readSimpleField("priority", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("CreateSubscriptionRequest");
    // Create the instance
    return new CreateSubscriptionRequestBuilderImpl(
        requestHeader,
        requestedPublishingInterval,
        requestedLifetimeCount,
        requestedMaxKeepAliveCount,
        maxNotificationsPerPublish,
        publishingEnabled,
        priority);
  }

  public static class CreateSubscriptionRequestBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final ExtensionObjectDefinition requestHeader;
    private final double requestedPublishingInterval;
    private final long requestedLifetimeCount;
    private final long requestedMaxKeepAliveCount;
    private final long maxNotificationsPerPublish;
    private final boolean publishingEnabled;
    private final short priority;

    public CreateSubscriptionRequestBuilderImpl(
        ExtensionObjectDefinition requestHeader,
        double requestedPublishingInterval,
        long requestedLifetimeCount,
        long requestedMaxKeepAliveCount,
        long maxNotificationsPerPublish,
        boolean publishingEnabled,
        short priority) {
      this.requestHeader = requestHeader;
      this.requestedPublishingInterval = requestedPublishingInterval;
      this.requestedLifetimeCount = requestedLifetimeCount;
      this.requestedMaxKeepAliveCount = requestedMaxKeepAliveCount;
      this.maxNotificationsPerPublish = maxNotificationsPerPublish;
      this.publishingEnabled = publishingEnabled;
      this.priority = priority;
    }

    public CreateSubscriptionRequest build() {
      CreateSubscriptionRequest createSubscriptionRequest =
          new CreateSubscriptionRequest(
              requestHeader,
              requestedPublishingInterval,
              requestedLifetimeCount,
              requestedMaxKeepAliveCount,
              maxNotificationsPerPublish,
              publishingEnabled,
              priority);
      return createSubscriptionRequest;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CreateSubscriptionRequest)) {
      return false;
    }
    CreateSubscriptionRequest that = (CreateSubscriptionRequest) o;
    return (getRequestHeader() == that.getRequestHeader())
        && (getRequestedPublishingInterval() == that.getRequestedPublishingInterval())
        && (getRequestedLifetimeCount() == that.getRequestedLifetimeCount())
        && (getRequestedMaxKeepAliveCount() == that.getRequestedMaxKeepAliveCount())
        && (getMaxNotificationsPerPublish() == that.getMaxNotificationsPerPublish())
        && (getPublishingEnabled() == that.getPublishingEnabled())
        && (getPriority() == that.getPriority())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getRequestHeader(),
        getRequestedPublishingInterval(),
        getRequestedLifetimeCount(),
        getRequestedMaxKeepAliveCount(),
        getMaxNotificationsPerPublish(),
        getPublishingEnabled(),
        getPriority());
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
