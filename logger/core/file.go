/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package logger is unified facade provided by Dubbo to work with different logger frameworks, eg, Zapper, Logrus.
package core

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/constant"
)

func FileConfig(config *common.URL) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   config.GetParam(constant.LoggerFileNameKey, "dubbo.log"),
		MaxSize:    config.GetParamByIntValue(constant.LoggerFileNaxSizeKey, 1),
		MaxBackups: config.GetParamByIntValue(constant.LoggerFileMaxBackupsKey, 1),
		MaxAge:     config.GetParamByIntValue(constant.LoggerFileMaxAgeKey, 3),
		LocalTime:  config.GetParamBool(constant.LoggerFileLocalTimeKey, true),
		Compress:   config.GetParamBool(constant.LoggerFileCompressKey, true),
	}
}
