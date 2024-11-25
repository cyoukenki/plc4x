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

public class NLMEstablishConnectionToNetwork extends NLM implements Message {

  // Accessors for discriminator values.
  public Short getMessageType() {
    return (short) 0x08;
  }

  // Properties.
  protected final int destinationNetworkAddress;
  protected final short terminationTime;

  // Arguments.
  protected final Integer apduLength;

  public NLMEstablishConnectionToNetwork(
      int destinationNetworkAddress, short terminationTime, Integer apduLength) {
    super(apduLength);
    this.destinationNetworkAddress = destinationNetworkAddress;
    this.terminationTime = terminationTime;
    this.apduLength = apduLength;
  }

  public int getDestinationNetworkAddress() {
    return destinationNetworkAddress;
  }

  public short getTerminationTime() {
    return terminationTime;
  }

  @Override
  protected void serializeNLMChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("NLMEstablishConnectionToNetwork");

    // Simple Field (destinationNetworkAddress)
    writeSimpleField(
        "destinationNetworkAddress", destinationNetworkAddress, writeUnsignedInt(writeBuffer, 16));

    // Simple Field (terminationTime)
    writeSimpleField("terminationTime", terminationTime, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("NLMEstablishConnectionToNetwork");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    NLMEstablishConnectionToNetwork _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (destinationNetworkAddress)
    lengthInBits += 16;

    // Simple field (terminationTime)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static NLMBuilder staticParseNLMBuilder(ReadBuffer readBuffer, Integer apduLength)
      throws ParseException {
    readBuffer.pullContext("NLMEstablishConnectionToNetwork");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int destinationNetworkAddress =
        readSimpleField("destinationNetworkAddress", readUnsignedInt(readBuffer, 16));

    short terminationTime = readSimpleField("terminationTime", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("NLMEstablishConnectionToNetwork");
    // Create the instance
    return new NLMEstablishConnectionToNetworkBuilderImpl(
        destinationNetworkAddress, terminationTime, apduLength);
  }

  public static class NLMEstablishConnectionToNetworkBuilderImpl implements NLM.NLMBuilder {
    private final int destinationNetworkAddress;
    private final short terminationTime;
    private final Integer apduLength;

    public NLMEstablishConnectionToNetworkBuilderImpl(
        int destinationNetworkAddress, short terminationTime, Integer apduLength) {
      this.destinationNetworkAddress = destinationNetworkAddress;
      this.terminationTime = terminationTime;
      this.apduLength = apduLength;
    }

    public NLMEstablishConnectionToNetwork build(Integer apduLength) {

      NLMEstablishConnectionToNetwork nLMEstablishConnectionToNetwork =
          new NLMEstablishConnectionToNetwork(
              destinationNetworkAddress, terminationTime, apduLength);
      return nLMEstablishConnectionToNetwork;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NLMEstablishConnectionToNetwork)) {
      return false;
    }
    NLMEstablishConnectionToNetwork that = (NLMEstablishConnectionToNetwork) o;
    return (getDestinationNetworkAddress() == that.getDestinationNetworkAddress())
        && (getTerminationTime() == that.getTerminationTime())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getDestinationNetworkAddress(), getTerminationTime());
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
