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
package org.apache.plc4x.java.profinet.readwrite;

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

public class PnIoCm_Submodule_InputAndOutputData extends PnIoCm_Submodule implements Message {

  // Accessors for discriminator values.
  public PnIoCm_SubmoduleType getSubmoduleType() {
    return PnIoCm_SubmoduleType.INPUT_AND_OUTPUT_DATA;
  }

  // Constant values.
  public static final Integer INPUTDATADESCRIPTION = 0x0001;
  public static final Short IOINPUTLENGTHIOCS = 0x01;
  public static final Short IOINPUTLENGTHIOPS = 0x01;
  public static final Integer OUTPUTDATADESCRIPTION = 0x0002;
  public static final Short IOOUTPUTLENGTHIOCS = 0x01;
  public static final Short IOOUTPUTLENGTHIOPS = 0x01;

  // Properties.
  protected final int inputSubmoduleDataLength;
  protected final int outputSubmoduleDataLength;

  public PnIoCm_Submodule_InputAndOutputData(
      int slotNumber,
      long submoduleIdentNumber,
      boolean discardIoxs,
      boolean reduceOutputModuleDataLength,
      boolean reduceInputModuleDataLength,
      boolean sharedInput,
      int inputSubmoduleDataLength,
      int outputSubmoduleDataLength) {
    super(
        slotNumber,
        submoduleIdentNumber,
        discardIoxs,
        reduceOutputModuleDataLength,
        reduceInputModuleDataLength,
        sharedInput);
    this.inputSubmoduleDataLength = inputSubmoduleDataLength;
    this.outputSubmoduleDataLength = outputSubmoduleDataLength;
  }

  public int getInputSubmoduleDataLength() {
    return inputSubmoduleDataLength;
  }

  public int getOutputSubmoduleDataLength() {
    return outputSubmoduleDataLength;
  }

  public int getInputDataDescription() {
    return INPUTDATADESCRIPTION;
  }

  public short getIoInputLengthIoCs() {
    return IOINPUTLENGTHIOCS;
  }

  public short getIoInputLengthIoPs() {
    return IOINPUTLENGTHIOPS;
  }

  public int getOutputDataDescription() {
    return OUTPUTDATADESCRIPTION;
  }

  public short getIoOutputLengthIoCs() {
    return IOOUTPUTLENGTHIOCS;
  }

  public short getIoOutputLengthIoPs() {
    return IOOUTPUTLENGTHIOPS;
  }

  @Override
  protected void serializePnIoCm_SubmoduleChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("PnIoCm_Submodule_InputAndOutputData");

    // Const Field (inputDataDescription)
    writeConstField(
        "inputDataDescription",
        INPUTDATADESCRIPTION,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (inputSubmoduleDataLength)
    writeSimpleField(
        "inputSubmoduleDataLength",
        inputSubmoduleDataLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (ioInputLengthIoCs)
    writeConstField(
        "ioInputLengthIoCs",
        IOINPUTLENGTHIOCS,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (ioInputLengthIoPs)
    writeConstField(
        "ioInputLengthIoPs",
        IOINPUTLENGTHIOPS,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (outputDataDescription)
    writeConstField(
        "outputDataDescription",
        OUTPUTDATADESCRIPTION,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Simple Field (outputSubmoduleDataLength)
    writeSimpleField(
        "outputSubmoduleDataLength",
        outputSubmoduleDataLength,
        writeUnsignedInt(writeBuffer, 16),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (ioOutputLengthIoCs)
    writeConstField(
        "ioOutputLengthIoCs",
        IOOUTPUTLENGTHIOCS,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    // Const Field (ioOutputLengthIoPs)
    writeConstField(
        "ioOutputLengthIoPs",
        IOOUTPUTLENGTHIOPS,
        writeUnsignedShort(writeBuffer, 8),
        WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    writeBuffer.popContext("PnIoCm_Submodule_InputAndOutputData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    PnIoCm_Submodule_InputAndOutputData _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Const Field (inputDataDescription)
    lengthInBits += 16;

    // Simple field (inputSubmoduleDataLength)
    lengthInBits += 16;

    // Const Field (ioInputLengthIoCs)
    lengthInBits += 8;

    // Const Field (ioInputLengthIoPs)
    lengthInBits += 8;

    // Const Field (outputDataDescription)
    lengthInBits += 16;

    // Simple field (outputSubmoduleDataLength)
    lengthInBits += 16;

    // Const Field (ioOutputLengthIoCs)
    lengthInBits += 8;

    // Const Field (ioOutputLengthIoPs)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static PnIoCm_SubmoduleBuilder staticParsePnIoCm_SubmoduleBuilder(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("PnIoCm_Submodule_InputAndOutputData");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    int inputDataDescription =
        readConstField(
            "inputDataDescription",
            readUnsignedInt(readBuffer, 16),
            PnIoCm_Submodule_InputAndOutputData.INPUTDATADESCRIPTION,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int inputSubmoduleDataLength =
        readSimpleField(
            "inputSubmoduleDataLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short ioInputLengthIoCs =
        readConstField(
            "ioInputLengthIoCs",
            readUnsignedShort(readBuffer, 8),
            PnIoCm_Submodule_InputAndOutputData.IOINPUTLENGTHIOCS,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short ioInputLengthIoPs =
        readConstField(
            "ioInputLengthIoPs",
            readUnsignedShort(readBuffer, 8),
            PnIoCm_Submodule_InputAndOutputData.IOINPUTLENGTHIOPS,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int outputDataDescription =
        readConstField(
            "outputDataDescription",
            readUnsignedInt(readBuffer, 16),
            PnIoCm_Submodule_InputAndOutputData.OUTPUTDATADESCRIPTION,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    int outputSubmoduleDataLength =
        readSimpleField(
            "outputSubmoduleDataLength",
            readUnsignedInt(readBuffer, 16),
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short ioOutputLengthIoCs =
        readConstField(
            "ioOutputLengthIoCs",
            readUnsignedShort(readBuffer, 8),
            PnIoCm_Submodule_InputAndOutputData.IOOUTPUTLENGTHIOCS,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    short ioOutputLengthIoPs =
        readConstField(
            "ioOutputLengthIoPs",
            readUnsignedShort(readBuffer, 8),
            PnIoCm_Submodule_InputAndOutputData.IOOUTPUTLENGTHIOPS,
            WithOption.WithByteOrder(ByteOrder.BIG_ENDIAN));

    readBuffer.closeContext("PnIoCm_Submodule_InputAndOutputData");
    // Create the instance
    return new PnIoCm_Submodule_InputAndOutputDataBuilderImpl(
        inputSubmoduleDataLength, outputSubmoduleDataLength);
  }

  public static class PnIoCm_Submodule_InputAndOutputDataBuilderImpl
      implements PnIoCm_Submodule.PnIoCm_SubmoduleBuilder {
    private final int inputSubmoduleDataLength;
    private final int outputSubmoduleDataLength;

    public PnIoCm_Submodule_InputAndOutputDataBuilderImpl(
        int inputSubmoduleDataLength, int outputSubmoduleDataLength) {
      this.inputSubmoduleDataLength = inputSubmoduleDataLength;
      this.outputSubmoduleDataLength = outputSubmoduleDataLength;
    }

    public PnIoCm_Submodule_InputAndOutputData build(
        int slotNumber,
        long submoduleIdentNumber,
        boolean discardIoxs,
        boolean reduceOutputModuleDataLength,
        boolean reduceInputModuleDataLength,
        boolean sharedInput) {
      PnIoCm_Submodule_InputAndOutputData pnIoCm_Submodule_InputAndOutputData =
          new PnIoCm_Submodule_InputAndOutputData(
              slotNumber,
              submoduleIdentNumber,
              discardIoxs,
              reduceOutputModuleDataLength,
              reduceInputModuleDataLength,
              sharedInput,
              inputSubmoduleDataLength,
              outputSubmoduleDataLength);
      return pnIoCm_Submodule_InputAndOutputData;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof PnIoCm_Submodule_InputAndOutputData)) {
      return false;
    }
    PnIoCm_Submodule_InputAndOutputData that = (PnIoCm_Submodule_InputAndOutputData) o;
    return (getInputSubmoduleDataLength() == that.getInputSubmoduleDataLength())
        && (getOutputSubmoduleDataLength() == that.getOutputSubmoduleDataLength())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getInputSubmoduleDataLength(), getOutputSubmoduleDataLength());
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