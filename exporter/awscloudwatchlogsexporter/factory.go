// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package awscloudwatchlogsexporter provides a logging exporter for the OpenTelemetry collector.
// This package is subject to change and may break configuration settings and behavior.
package awscloudwatchlogsexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/exporter/exporterhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/awsutil"
)

const (
	typeStr = "awscloudwatchlogs"
	// The stability level of the exporter.
	stability = component.StabilityLevelBeta
)

func NewFactory() component.ExporterFactory {
	return component.NewExporterFactory(
		typeStr,
		createDefaultConfig,
		component.WithLogsExporter(createLogsExporter, stability))
}

func createDefaultConfig() config.Exporter {
	return &Config{
		ExporterSettings:   config.NewExporterSettings(config.NewComponentID(typeStr)),
		RetrySettings:      exporterhelper.NewDefaultRetrySettings(),
		AWSSessionSettings: awsutil.CreateDefaultSessionConfig(),
		QueueSettings: QueueSettings{
			QueueSize: exporterhelper.NewDefaultQueueSettings().QueueSize,
		},
	}
}

func createLogsExporter(_ context.Context, params component.ExporterCreateSettings, config config.Exporter) (component.LogsExporter, error) {
	expConfig, ok := config.(*Config)
	if !ok {
		return nil, errors.New("invalid configuration type; can't cast to awscloudwatchlogsexporter.Config")
	}
	return newCwLogsExporter(expConfig, params)

}
