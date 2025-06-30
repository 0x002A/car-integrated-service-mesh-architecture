// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2025 MBition GmbH

package web

import "embed"

//go:embed all:template
var Templates embed.FS

//go:embed all:static
var Assets embed.FS
