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

public class CANOpenSDOResponse extends CANOpenPayload implements Message {

  // Accessors for discriminator values.
  public CANOpenService getService() {
    return CANOpenService.TRANSMIT_SDO;
  }

  // Properties.
  protected final SDOResponseCommand command;
  protected final SDOResponse response;

  public CANOpenSDOResponse(SDOResponseCommand command, SDOResponse response) {
    super();
    this.command = command;
    this.response = response;
  }

  public SDOResponseCommand getCommand() {
    return command;
  }

  public SDOResponse getResponse() {
    return response;
  }

  @Override
  protected void serializeCANOpenPayloadChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("CANOpenSDOResponse");

    // Simple Field (command)
    writeSimpleEnumField(
        "command",
        "SDOResponseCommand",
        command,
        new DataWriterEnumDefault<>(
            SDOResponseCommand::getValue,
            SDOResponseCommand::name,
            writeUnsignedByte(writeBuffer, 3)));

    // Simple Field (response)
    writeSimpleField("response", response, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("CANOpenSDOResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    CANOpenSDOResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (command)
    lengthInBits += 3;

    // Simple field (response)
    lengthInBits += response.getLengthInBits();

    return lengthInBits;
  }

  public static CANOpenPayloadBuilder staticParseCANOpenPayloadBuilder(
      ReadBuffer readBuffer, CANOpenService service) throws ParseException {
    readBuffer.pullContext("CANOpenSDOResponse");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    SDOResponseCommand command =
        readEnumField(
            "command",
            "SDOResponseCommand",
            new DataReaderEnumDefault<>(
                SDOResponseCommand::enumForValue, readUnsignedByte(readBuffer, 3)));

    SDOResponse response =
        readSimpleField(
            "response",
            new DataReaderComplexDefault<>(
                () -> SDOResponse.staticParse(readBuffer, (SDOResponseCommand) (command)),
                readBuffer));

    readBuffer.closeContext("CANOpenSDOResponse");
    // Create the instance
    return new CANOpenSDOResponseBuilderImpl(command, response);
  }

  public static class CANOpenSDOResponseBuilderImpl
      implements CANOpenPayload.CANOpenPayloadBuilder {
    private final SDOResponseCommand command;
    private final SDOResponse response;

    public CANOpenSDOResponseBuilderImpl(SDOResponseCommand command, SDOResponse response) {
      this.command = command;
      this.response = response;
    }

    public CANOpenSDOResponse build() {
      CANOpenSDOResponse cANOpenSDOResponse = new CANOpenSDOResponse(command, response);
      return cANOpenSDOResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof CANOpenSDOResponse)) {
      return false;
    }
    CANOpenSDOResponse that = (CANOpenSDOResponse) o;
    return (getCommand() == that.getCommand())
        && (getResponse() == that.getResponse())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCommand(), getResponse());
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