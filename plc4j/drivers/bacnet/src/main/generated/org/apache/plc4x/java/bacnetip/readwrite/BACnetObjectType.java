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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum BACnetObjectType {
  ACCESS_CREDENTIAL((short) 32),
  ACCESS_DOOR((short) 30),
  ACCESS_POINT((short) 33),
  ACCESS_RIGHTS((short) 34),
  ACCESS_USER((short) 35),
  ACCESS_ZONE((short) 36),
  ACCUMULATOR((short) 23),
  ALERT_ENROLLMENT((short) 52),
  ANALOG_INPUT((short) 0),
  ANALOG_OUTPUT((short) 1),
  ANALOG_VALUE((short) 2),
  AVERAGING((short) 18),
  BINARY_INPUT((short) 3),
  BINARY_LIGHTING_OUTPUT((short) 55),
  BINARY_OUTPUT((short) 4),
  BINARY_VALUE((short) 5),
  BITSTRING_VALUE((short) 39),
  CALENDAR((short) 6),
  CHANNEL((short) 53),
  CHARACTERSTRING_VALUE((short) 40),
  COMMAND((short) 7),
  CREDENTIAL_DATA_INPUT((short) 37),
  DATEPATTERN_VALUE((short) 41),
  DATE_VALUE((short) 42),
  DATETIMEPATTERN_VALUE((short) 43),
  DATETIME_VALUE((short) 44),
  DEVICE((short) 8),
  ELEVATOR_GROUP((short) 57),
  ESCALATOR((short) 58),
  EVENT_ENROLLMENT((short) 9),
  EVENT_LOG((short) 25),
  FILE((short) 10),
  GLOBAL_GROUP((short) 26),
  GROUP((short) 11),
  INTEGER_VALUE((short) 45),
  LARGE_ANALOG_VALUE((short) 46),
  LIFE_SAFETY_POINT((short) 21),
  LIFE_SAFETY_ZONE((short) 22),
  LIFT((short) 59),
  LIGHTING_OUTPUT((short) 54),
  LOAD_CONTROL((short) 28),
  LOOP((short) 12),
  MULTI_STATE_INPUT((short) 13),
  MULTI_STATE_OUTPUT((short) 14),
  MULTI_STATE_VALUE((short) 19),
  NETWORK_PORT((short) 56),
  NETWORK_SECURITY((short) 38),
  NOTIFICATION_CLASS((short) 15),
  NOTIFICATION_FORWARDER((short) 51),
  OCTETSTRING_VALUE((short) 47),
  POSITIVE_INTEGER_VALUE((short) 48),
  PROGRAM((short) 16),
  PULSE_CONVERTER((short) 24),
  SCHEDULE((short) 17),
  STRUCTURED_VIEW((short) 29),
  TIMEPATTERN_VALUE((short) 49),
  TIME_VALUE((short) 50),
  TIMER((short) 31),
  TREND_LOG((short) 20),
  TREND_LOG_MULTIPLE((short) 27),
  VENDOR_PROPRIETARY_VALUE((short) 0x3FF);
  private static final Map<Short, BACnetObjectType> map;

  static {
    map = new HashMap<>();
    for (BACnetObjectType value : BACnetObjectType.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;

  BACnetObjectType(short value) {
    this.value = value;
  }

  public short getValue() {
    return value;
  }

  public static BACnetObjectType enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
