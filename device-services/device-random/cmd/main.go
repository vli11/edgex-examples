// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2021 IOTech Ltd
// Copyright 2023 Intel Corp.
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	device_random "github.com/edgexfoundry/device-random"
	"github.com/edgexfoundry/device-random/driver"
	"github.com/edgexfoundry/device-sdk-go/v3/pkg/startup"
)

const (
	serviceName string = "device-random"
)

func main() {
	d := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, device_random.Version, d)
}
