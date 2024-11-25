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

public class EndpointUrlListDataType extends ExtensionObjectDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "11945";
  }

  // Properties.
  protected final int noOfEndpointUrlList;
  protected final List<PascalString> endpointUrlList;

  public EndpointUrlListDataType(int noOfEndpointUrlList, List<PascalString> endpointUrlList) {
    super();
    this.noOfEndpointUrlList = noOfEndpointUrlList;
    this.endpointUrlList = endpointUrlList;
  }

  public int getNoOfEndpointUrlList() {
    return noOfEndpointUrlList;
  }

  public List<PascalString> getEndpointUrlList() {
    return endpointUrlList;
  }

  @Override
  protected void serializeExtensionObjectDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("EndpointUrlListDataType");

    // Simple Field (noOfEndpointUrlList)
    writeSimpleField("noOfEndpointUrlList", noOfEndpointUrlList, writeSignedInt(writeBuffer, 32));

    // Array Field (endpointUrlList)
    writeComplexTypeArrayField("endpointUrlList", endpointUrlList, writeBuffer);

    writeBuffer.popContext("EndpointUrlListDataType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    EndpointUrlListDataType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (noOfEndpointUrlList)
    lengthInBits += 32;

    // Array field
    if (endpointUrlList != null) {
      int i = 0;
      for (PascalString element : endpointUrlList) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= endpointUrlList.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ExtensionObjectDefinitionBuilder staticParseExtensionObjectDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("EndpointUrlListDataType");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int noOfEndpointUrlList = readSimpleField("noOfEndpointUrlList", readSignedInt(readBuffer, 32));

    List<PascalString> endpointUrlList =
        readCountArrayField(
            "endpointUrlList",
            new DataReaderComplexDefault<>(() -> PascalString.staticParse(readBuffer), readBuffer),
            noOfEndpointUrlList);

    readBuffer.closeContext("EndpointUrlListDataType");
    // Create the instance
    return new EndpointUrlListDataTypeBuilderImpl(noOfEndpointUrlList, endpointUrlList);
  }

  public static class EndpointUrlListDataTypeBuilderImpl
      implements ExtensionObjectDefinition.ExtensionObjectDefinitionBuilder {
    private final int noOfEndpointUrlList;
    private final List<PascalString> endpointUrlList;

    public EndpointUrlListDataTypeBuilderImpl(
        int noOfEndpointUrlList, List<PascalString> endpointUrlList) {
      this.noOfEndpointUrlList = noOfEndpointUrlList;
      this.endpointUrlList = endpointUrlList;
    }

    public EndpointUrlListDataType build() {
      EndpointUrlListDataType endpointUrlListDataType =
          new EndpointUrlListDataType(noOfEndpointUrlList, endpointUrlList);
      return endpointUrlListDataType;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof EndpointUrlListDataType)) {
      return false;
    }
    EndpointUrlListDataType that = (EndpointUrlListDataType) o;
    return (getNoOfEndpointUrlList() == that.getNoOfEndpointUrlList())
        && (getEndpointUrlList() == that.getEndpointUrlList())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getNoOfEndpointUrlList(), getEndpointUrlList());
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
