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

public class SessionlessInvokeResponseType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "21001";
  }

  // Properties.
  protected final int noOfNamespaceUris;
  protected final List<PascalString> namespaceUris;
  protected final int noOfServerUris;
  protected final List<PascalString> serverUris;
  protected final long serviceId;

  public SessionlessInvokeResponseType(
      int noOfNamespaceUris,
      List<PascalString> namespaceUris,
      int noOfServerUris,
      List<PascalString> serverUris,
      long serviceId) {
    super();
    this.noOfNamespaceUris = noOfNamespaceUris;
    this.namespaceUris = namespaceUris;
    this.noOfServerUris = noOfServerUris;
    this.serverUris = serverUris;
    this.serviceId = serviceId;
  }

  public int getNoOfNamespaceUris() {
    return noOfNamespaceUris;
  }

  public List<PascalString> getNamespaceUris() {
    return namespaceUris;
  }

  public int getNoOfServerUris() {
    return noOfServerUris;
  }

  public List<PascalString> getServerUris() {
    return serverUris;
  }

  public long getServiceId() {
    return serviceId;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("SessionlessInvokeResponseType");

    // Simple Field (noOfNamespaceUris)
    writeSimpleField("noOfNamespaceUris", noOfNamespaceUris, writeSignedInt(writeBuffer, 32));

    // Array Field (namespaceUris)
    writeComplexTypeArrayField("namespaceUris", namespaceUris, writeBuffer);

    // Simple Field (noOfServerUris)
    writeSimpleField("noOfServerUris", noOfServerUris, writeSignedInt(writeBuffer, 32));

    // Array Field (serverUris)
    writeComplexTypeArrayField("serverUris", serverUris, writeBuffer);

    // Simple Field (serviceId)
    writeSimpleField("serviceId", serviceId, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("SessionlessInvokeResponseType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    SessionlessInvokeResponseType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (noOfNamespaceUris)
    lengthInBits += 32;

    // Array field
    if (namespaceUris != null) {
      int i = 0;
      for (PascalString element : namespaceUris) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= namespaceUris.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (noOfServerUris)
    lengthInBits += 32;

    // Array field
    if (serverUris != null) {
      int i = 0;
      for (PascalString element : serverUris) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= serverUris.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    // Simple field (serviceId)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("SessionlessInvokeResponseType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int noOfNamespaceUris = readSimpleField("noOfNamespaceUris", readSignedInt(readBuffer, 32));

    List<PascalString> namespaceUris =
        readCountArrayField(
            "namespaceUris",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfNamespaceUris);

    int noOfServerUris = readSimpleField("noOfServerUris", readSignedInt(readBuffer, 32));

    List<PascalString> serverUris =
        readCountArrayField(
            "serverUris",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfServerUris);

    long serviceId = readSimpleField("serviceId", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("SessionlessInvokeResponseType");
    // Create the instance
    return new SessionlessInvokeResponseTypeBuilderImpl(
        noOfNamespaceUris, namespaceUris, noOfServerUris, serverUris, serviceId);
  }

  public static class SessionlessInvokeResponseTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final int noOfNamespaceUris;
    private final List<PascalString> namespaceUris;
    private final int noOfServerUris;
    private final List<PascalString> serverUris;
    private final long serviceId;

    public SessionlessInvokeResponseTypeBuilderImpl(
        int noOfNamespaceUris,
        List<PascalString> namespaceUris,
        int noOfServerUris,
        List<PascalString> serverUris,
        long serviceId) {
      this.noOfNamespaceUris = noOfNamespaceUris;
      this.namespaceUris = namespaceUris;
      this.noOfServerUris = noOfServerUris;
      this.serverUris = serverUris;
      this.serviceId = serviceId;
    }

    public SessionlessInvokeResponseType build() {
      SessionlessInvokeResponseType sessionlessInvokeResponseType =
          new SessionlessInvokeResponseType(
              noOfNamespaceUris, namespaceUris, noOfServerUris, serverUris, serviceId);
      return sessionlessInvokeResponseType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof SessionlessInvokeResponseType)) {
      return false;
    }
    SessionlessInvokeResponseType that = (SessionlessInvokeResponseType) o;
    return (getNoOfNamespaceUris() == that.getNoOfNamespaceUris())
        && (getNamespaceUris() == that.getNamespaceUris())
        && (getNoOfServerUris() == that.getNoOfServerUris())
        && (getServerUris() == that.getServerUris())
        && (getServiceId() == that.getServiceId())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(),
        getNoOfNamespaceUris(),
        getNamespaceUris(),
        getNoOfServerUris(),
        getServerUris(),
        getServiceId());
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
