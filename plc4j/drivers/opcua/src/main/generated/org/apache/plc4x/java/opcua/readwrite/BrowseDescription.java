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

public class BrowseDescription extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "516";
  }

  // Properties.
  protected final NodeId nodeId;
  protected final BrowseDirection browseDirection;
  protected final NodeId referenceTypeId;
  protected final boolean includeSubtypes;
  protected final long nodeClassMask;
  protected final long resultMask;

  public BrowseDescription(
      NodeId nodeId,
      BrowseDirection browseDirection,
      NodeId referenceTypeId,
      boolean includeSubtypes,
      long nodeClassMask,
      long resultMask) {
    super();
    this.nodeId = nodeId;
    this.browseDirection = browseDirection;
    this.referenceTypeId = referenceTypeId;
    this.includeSubtypes = includeSubtypes;
    this.nodeClassMask = nodeClassMask;
    this.resultMask = resultMask;
  }

  public NodeId getNodeId() {
    return nodeId;
  }

  public BrowseDirection getBrowseDirection() {
    return browseDirection;
  }

  public NodeId getReferenceTypeId() {
    return referenceTypeId;
  }

  public boolean getIncludeSubtypes() {
    return includeSubtypes;
  }

  public long getNodeClassMask() {
    return nodeClassMask;
  }

  public long getResultMask() {
    return resultMask;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("BrowseDescription");

    // Simple Field (nodeId)
    writeSimpleField("nodeId", nodeId, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (browseDirection)
    writeSimpleEnumField(
        "browseDirection",
        "BrowseDirection",
        browseDirection,
        new DataWriterEnumDefault<>(
            BrowseDirection::getValue, BrowseDirection::name, writeUnsignedLong(writeBuffer, 32)));

    // Simple Field (referenceTypeId)
    writeSimpleField(
        "referenceTypeId", referenceTypeId, new DataWriterComplexDefault<>(writeBuffer));

    // Reserved Field (reserved)
    writeReservedField("reserved", (byte) 0x00, writeUnsignedByte(writeBuffer, 7));

    // Simple Field (includeSubtypes)
    writeSimpleField("includeSubtypes", includeSubtypes, writeBoolean(writeBuffer));

    // Simple Field (nodeClassMask)
    writeSimpleField("nodeClassMask", nodeClassMask, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (resultMask)
    writeSimpleField("resultMask", resultMask, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("BrowseDescription");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BrowseDescription _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (nodeId)
    lengthInBits += nodeId.getLengthInBits();

    // Simple field (browseDirection)
    lengthInBits += 32;

    // Simple field (referenceTypeId)
    lengthInBits += referenceTypeId.getLengthInBits();

    // Reserved Field (reserved)
    lengthInBits += 7;

    // Simple field (includeSubtypes)
    lengthInBits += 1;

    // Simple field (nodeClassMask)
    lengthInBits += 32;

    // Simple field (resultMask)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("BrowseDescription");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    NodeId nodeId =
        readSimpleField(
            "nodeId",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    BrowseDirection browseDirection =
        readEnumField(
            "browseDirection",
            "BrowseDirection",
            new DataReaderEnumDefault<>(
                BrowseDirection::enumForValue, readUnsignedLong(readBuffer, 32)));

    NodeId referenceTypeId =
        readSimpleField(
            "referenceTypeId",
            new DataReaderComplexDefault<>(() -> NodeId.staticParse(readBuffer), readBuffer));

    Byte reservedField0 =
        readReservedField("reserved", readUnsignedByte(readBuffer, 7), (byte) 0x00);

    boolean includeSubtypes = readSimpleField("includeSubtypes", readBoolean(readBuffer));

    long nodeClassMask = readSimpleField("nodeClassMask", readUnsignedLong(readBuffer, 32));

    long resultMask = readSimpleField("resultMask", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("BrowseDescription");
    // Create the instance
    return new BrowseDescriptionBuilderImpl(
        nodeId, browseDirection, referenceTypeId, includeSubtypes, nodeClassMask, resultMask);
  }

  public static class BrowseDescriptionBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final NodeId nodeId;
    private final BrowseDirection browseDirection;
    private final NodeId referenceTypeId;
    private final boolean includeSubtypes;
    private final long nodeClassMask;
    private final long resultMask;

    public BrowseDescriptionBuilderImpl(
        NodeId nodeId,
        BrowseDirection browseDirection,
        NodeId referenceTypeId,
        boolean includeSubtypes,
        long nodeClassMask,
        long resultMask) {
      this.nodeId = nodeId;
      this.browseDirection = browseDirection;
      this.referenceTypeId = referenceTypeId;
      this.includeSubtypes = includeSubtypes;
      this.nodeClassMask = nodeClassMask;
      this.resultMask = resultMask;
    }

    public BrowseDescription build() {
      BrowseDescription browseDescription =
          new BrowseDescription(
              nodeId, browseDirection, referenceTypeId, includeSubtypes, nodeClassMask, resultMask);
      return browseDescription;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BrowseDescription)) {
      return false;
    }
    BrowseDescription that = (BrowseDescription) o;
    return (getNodeId() == that.getNodeId())
        && (getBrowseDirection() == that.getBrowseDirection())
        && (getReferenceTypeId() == that.getReferenceTypeId())
        && (getIncludeSubtypes() == that.getIncludeSubtypes())
        && (getNodeClassMask() == that.getNodeClassMask())
        && (getResultMask() == that.getResultMask())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getNodeId(),
        getBrowseDirection(),
        getReferenceTypeId(),
        getIncludeSubtypes(),
        getNodeClassMask(),
        getResultMask());
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
