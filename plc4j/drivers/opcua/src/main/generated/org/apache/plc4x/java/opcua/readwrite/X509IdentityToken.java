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

public class X509IdentityToken extends UserIdentityTokenDefinition implements Message {

  // Accessors for discriminator values.
  public String getIdentifier() {
    return (String) "certificate";
  }

  // Properties.
  protected final PascalByteString certificateData;

  public X509IdentityToken(PascalByteString certificateData) {
    super();
    this.certificateData = certificateData;
  }

  public PascalByteString getCertificateData() {
    return certificateData;
  }

  @Override
  protected void serializeUserIdentityTokenDefinitionChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("X509IdentityToken");

    // Simple Field (certificateData)
    writeSimpleField(
        "certificateData", certificateData, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("X509IdentityToken");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    X509IdentityToken _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (certificateData)
    lengthInBits += certificateData.getLengthInBits();

    return lengthInBits;
  }

  public static UserIdentityTokenDefinitionBuilder staticParseUserIdentityTokenDefinitionBuilder(
      ReadBuffer readBuffer, String identifier) throws ParseException {
    readBuffer.pullContext("X509IdentityToken");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    PascalByteString certificateData =
        readSimpleField(
            "certificateData",
            new DataReaderComplexDefault<>(
                () -> PascalByteString.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("X509IdentityToken");
    // Create the instance
    return new X509IdentityTokenBuilderImpl(certificateData);
  }

  public static class X509IdentityTokenBuilderImpl
      implements UserIdentityTokenDefinition.UserIdentityTokenDefinitionBuilder {
    private final PascalByteString certificateData;

    public X509IdentityTokenBuilderImpl(PascalByteString certificateData) {
      this.certificateData = certificateData;
    }

    public X509IdentityToken build() {
      X509IdentityToken x509IdentityToken = new X509IdentityToken(certificateData);
      return x509IdentityToken;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof X509IdentityToken)) {
      return false;
    }
    X509IdentityToken that = (X509IdentityToken) o;
    return (getCertificateData() == that.getCertificateData()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCertificateData());
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