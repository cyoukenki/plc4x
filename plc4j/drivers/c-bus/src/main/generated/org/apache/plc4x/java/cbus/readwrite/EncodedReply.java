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
package org.apache.plc4x.java.cbus.readwrite;

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

public abstract class EncodedReply implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final byte peekedByte;

  // Arguments.
  protected final CBusOptions cBusOptions;
  protected final RequestContext requestContext;

  public EncodedReply(byte peekedByte, CBusOptions cBusOptions, RequestContext requestContext) {
    super();
    this.peekedByte = peekedByte;
    this.cBusOptions = cBusOptions;
    this.requestContext = requestContext;
  }

  public byte getPeekedByte() {
    return peekedByte;
  }

  public boolean getIsMonitoredSAL() {
    return (boolean)
        ((((((((getPeekedByte()) & (0x3F))) == (0x05)) || ((getPeekedByte()) == (0x00)))
                || ((((getPeekedByte()) & (0xF8))) == (0x00))))
            && (!(requestContext.getSendIdentifyRequestBefore())));
  }

  protected abstract void serializeEncodedReplyChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("EncodedReply");

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    boolean isMonitoredSAL = getIsMonitoredSAL();
    writeBuffer.writeVirtual("isMonitoredSAL", isMonitoredSAL);

    // Switch field (Serialize the sub-type)
    serializeEncodedReplyChild(writeBuffer);

    writeBuffer.popContext("EncodedReply");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    EncodedReply _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // A virtual field doesn't have any in- or output.

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static EncodedReply staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    if ((args == null) || (args.length != 2)) {
      throw new PlcRuntimeException(
          "Wrong number of arguments, expected 2, but got " + args.length);
    }
    CBusOptions cBusOptions;
    if (args[0] instanceof CBusOptions) {
      cBusOptions = (CBusOptions) args[0];
    } else {
      throw new PlcRuntimeException(
          "Argument 0 expected to be of type CBusOptions or a string which is parseable but was "
              + args[0].getClass().getName());
    }
    RequestContext requestContext;
    if (args[1] instanceof RequestContext) {
      requestContext = (RequestContext) args[1];
    } else {
      throw new PlcRuntimeException(
          "Argument 1 expected to be of type RequestContext or a string which is parseable but was "
              + args[1].getClass().getName());
    }
    return staticParse(readBuffer, cBusOptions, requestContext);
  }

  public static EncodedReply staticParse(
      ReadBuffer readBuffer, CBusOptions cBusOptions, RequestContext requestContext)
      throws ParseException {
    readBuffer.pullContext("EncodedReply");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte peekedByte = readPeekField("peekedByte", readByte(readBuffer, 8));
    boolean isMonitoredSAL =
        readVirtualField(
            "isMonitoredSAL",
            boolean.class,
            (((((((peekedByte) & (0x3F))) == (0x05)) || ((peekedByte) == (0x00)))
                    || ((((peekedByte) & (0xF8))) == (0x00))))
                && (!(requestContext.getSendIdentifyRequestBefore())));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    EncodedReplyBuilder builder = null;
    if (EvaluationHelper.equals(isMonitoredSAL, (boolean) true)) {
      builder =
          MonitoredSALReply.staticParseEncodedReplyBuilder(readBuffer, cBusOptions, requestContext);
    } else if (true) {
      builder =
          EncodedReplyCALReply.staticParseEncodedReplyBuilder(
              readBuffer, cBusOptions, requestContext);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "isMonitoredSAL="
              + isMonitoredSAL
              + "]");
    }

    readBuffer.closeContext("EncodedReply");
    // Create the instance
    EncodedReply _encodedReply = builder.build(peekedByte, cBusOptions, requestContext);
    return _encodedReply;
  }

  public interface EncodedReplyBuilder {
    EncodedReply build(byte peekedByte, CBusOptions cBusOptions, RequestContext requestContext);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof EncodedReply)) {
      return false;
    }
    EncodedReply that = (EncodedReply) o;
    return (getPeekedByte() == that.getPeekedByte()) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getPeekedByte());
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
