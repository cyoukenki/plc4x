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

public abstract class MediaTransportControlData implements Message {

  // Abstract accessors for discriminator values.

  // Properties.
  protected final MediaTransportControlCommandTypeContainer commandTypeContainer;
  protected final byte mediaLinkGroup;

  public MediaTransportControlData(
      MediaTransportControlCommandTypeContainer commandTypeContainer, byte mediaLinkGroup) {
    super();
    this.commandTypeContainer = commandTypeContainer;
    this.mediaLinkGroup = mediaLinkGroup;
  }

  public MediaTransportControlCommandTypeContainer getCommandTypeContainer() {
    return commandTypeContainer;
  }

  public byte getMediaLinkGroup() {
    return mediaLinkGroup;
  }

  public MediaTransportControlCommandType getCommandType() {
    return (MediaTransportControlCommandType) (getCommandTypeContainer().getCommandType());
  }

  protected abstract void serializeMediaTransportControlDataChild(WriteBuffer writeBuffer)
      throws SerializationException;

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    writeBuffer.pushContext("MediaTransportControlData");

    // Simple Field (commandTypeContainer)
    writeSimpleEnumField(
        "commandTypeContainer",
        "MediaTransportControlCommandTypeContainer",
        commandTypeContainer,
        new DataWriterEnumDefault<>(
            MediaTransportControlCommandTypeContainer::getValue,
            MediaTransportControlCommandTypeContainer::name,
            writeUnsignedShort(writeBuffer, 8)));

    // Virtual field (doesn't actually serialize anything, just makes the value available)
    MediaTransportControlCommandType commandType = getCommandType();
    writeBuffer.writeVirtual("commandType", commandType);

    // Simple Field (mediaLinkGroup)
    writeSimpleField("mediaLinkGroup", mediaLinkGroup, writeByte(writeBuffer, 8));

    // Switch field (Serialize the sub-type)
    serializeMediaTransportControlDataChild(writeBuffer);

    writeBuffer.popContext("MediaTransportControlData");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    MediaTransportControlData _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (commandTypeContainer)
    lengthInBits += 8;

    // A virtual field doesn't have any in- or output.

    // Simple field (mediaLinkGroup)
    lengthInBits += 8;

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits;
  }

  public static MediaTransportControlData staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static MediaTransportControlData staticParse(ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("MediaTransportControlData");
    PositionAware positionAware = readBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    // Validation
    if (!(org.apache.plc4x.java.cbus.readwrite.utils.StaticHelper
        .knowsMediaTransportControlCommandTypeContainer(readBuffer))) {
      throw new ParseAssertException("no command type could be found");
    }

    MediaTransportControlCommandTypeContainer commandTypeContainer =
        readEnumField(
            "commandTypeContainer",
            "MediaTransportControlCommandTypeContainer",
            new DataReaderEnumDefault<>(
                MediaTransportControlCommandTypeContainer::enumForValue,
                readUnsignedShort(readBuffer, 8)));
    MediaTransportControlCommandType commandType =
        readVirtualField(
            "commandType",
            MediaTransportControlCommandType.class,
            commandTypeContainer.getCommandType());

    byte mediaLinkGroup = readSimpleField("mediaLinkGroup", readByte(readBuffer, 8));

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    MediaTransportControlDataBuilder builder = null;
    if (EvaluationHelper.equals(commandType, MediaTransportControlCommandType.STOP)) {
      builder =
          MediaTransportControlDataStop.staticParseMediaTransportControlDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(commandType, MediaTransportControlCommandType.PLAY)) {
      builder =
          MediaTransportControlDataPlay.staticParseMediaTransportControlDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.PAUSE_RESUME)) {
      builder =
          MediaTransportControlDataPauseResume.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.SELECT_CATEGORY)) {
      builder =
          MediaTransportControlDataSetCategory.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.SELECT_SELECTION)) {
      builder =
          MediaTransportControlDataSetSelection.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.SELECT_TRACK)) {
      builder =
          MediaTransportControlDataSetTrack.staticParseMediaTransportControlDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.SHUFFLE_ON_OFF)) {
      builder =
          MediaTransportControlDataShuffleOnOff.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.REPEAT_ON_OFF)) {
      builder =
          MediaTransportControlDataRepeatOnOff.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.NEXT_PREVIOUS_CATEGORY)) {
      builder =
          MediaTransportControlDataNextPreviousCategory.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.NEXT_PREVIOUS_SELECTION)) {
      builder =
          MediaTransportControlDataNextPreviousSelection
              .staticParseMediaTransportControlDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.NEXT_PREVIOUS_TRACK)) {
      builder =
          MediaTransportControlDataNextPreviousTrack.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.FAST_FORWARD)) {
      builder =
          MediaTransportControlDataFastForward.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(commandType, MediaTransportControlCommandType.REWIND)) {
      builder =
          MediaTransportControlDataRewind.staticParseMediaTransportControlDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.SOURCE_POWER_CONTROL)) {
      builder =
          MediaTransportControlDataSourcePowerControl.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.TOTAL_TRACKS)) {
      builder =
          MediaTransportControlDataTotalTracks.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.STATUS_REQUEST)) {
      builder =
          MediaTransportControlDataStatusRequest.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.ENUMERATE_CATEGORIES_SELECTIONS_TRACKS)) {
      builder =
          MediaTransportControlDataEnumerateCategoriesSelectionTracks
              .staticParseMediaTransportControlDataBuilder(readBuffer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.ENUMERATION_SIZE)) {
      builder =
          MediaTransportControlDataEnumerationsSize.staticParseMediaTransportControlDataBuilder(
              readBuffer);
    } else if (EvaluationHelper.equals(commandType, MediaTransportControlCommandType.TRACK_NAME)) {
      builder =
          MediaTransportControlDataTrackName.staticParseMediaTransportControlDataBuilder(
              readBuffer, commandTypeContainer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.SELECTION_NAME)) {
      builder =
          MediaTransportControlDataSelectionName.staticParseMediaTransportControlDataBuilder(
              readBuffer, commandTypeContainer);
    } else if (EvaluationHelper.equals(
        commandType, MediaTransportControlCommandType.CATEGORY_NAME)) {
      builder =
          MediaTransportControlDataCategoryName.staticParseMediaTransportControlDataBuilder(
              readBuffer, commandTypeContainer);
    }
    if (builder == null) {
      throw new ParseException(
          "Unsupported case for discriminated type"
              + " parameters ["
              + "commandType="
              + commandType
              + "]");
    }

    readBuffer.closeContext("MediaTransportControlData");
    // Create the instance
    MediaTransportControlData _mediaTransportControlData =
        builder.build(commandTypeContainer, mediaLinkGroup);
    return _mediaTransportControlData;
  }

  public interface MediaTransportControlDataBuilder {
    MediaTransportControlData build(
        MediaTransportControlCommandTypeContainer commandTypeContainer, byte mediaLinkGroup);
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MediaTransportControlData)) {
      return false;
    }
    MediaTransportControlData that = (MediaTransportControlData) o;
    return (getCommandTypeContainer() == that.getCommandTypeContainer())
        && (getMediaLinkGroup() == that.getMediaLinkGroup())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getCommandTypeContainer(), getMediaLinkGroup());
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
