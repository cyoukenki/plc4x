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

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariableSubscribed {
  SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_InputArguments(
      (int) 23798L),
  SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddSubscribedDataSet_OutputArguments(
      (int) 23799L),
  SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveSubscribedDataSet_InputArguments(
      (int) 23801L),
  SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_InputArguments(
      (int) 23803L),
  SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_AddDataSetFolder_OutputArguments(
      (int) 23804L),
  SubscribedDataSetFolderType_SubscribedDataSetFolderName_Placeholder_RemoveDataSetFolder_InputArguments(
      (int) 23806L),
  SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_DataSetMetaData(
      (int) 23809L),
  SubscribedDataSetFolderType_StandaloneSubscribedDataSetName_Placeholder_IsConnected((int) 23810L),
  SubscribedDataSetFolderType_AddSubscribedDataSet_InputArguments((int) 23812L),
  SubscribedDataSetFolderType_AddSubscribedDataSet_OutputArguments((int) 23813L),
  SubscribedDataSetFolderType_RemoveSubscribedDataSet_InputArguments((int) 23815L),
  SubscribedDataSetFolderType_AddDataSetFolder_InputArguments((int) 23817L),
  SubscribedDataSetFolderType_AddDataSetFolder_OutputArguments((int) 23818L),
  SubscribedDataSetFolderType_RemoveDataSetFolder_InputArguments((int) 23820L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableSubscribed> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableSubscribed value :
        OpcuaNodeIdServicesVariableSubscribed.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableSubscribed(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableSubscribed enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
