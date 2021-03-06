// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package loggingexporter

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/consumer/pdatautil"
	"go.opentelemetry.io/collector/internal/data/testdata"
)

func TestLoggingTraceExporterNoErrors(t *testing.T) {
	lte, err := NewTraceExporter(&configmodels.ExporterSettings{}, "debug", zap.NewNop())
	require.NotNil(t, lte)
	assert.NoError(t, err)

	assert.NoError(t, lte.ConsumeTraces(context.Background(), testdata.GenerateTraceDataEmpty()))
	assert.NoError(t, lte.ConsumeTraces(context.Background(), testdata.GenerateTraceDataOneEmptyOneNilResourceSpans()))
	assert.NoError(t, lte.ConsumeTraces(context.Background(), testdata.GenerateTraceDataOneEmptyOneNilInstrumentationLibrary()))
	assert.NoError(t, lte.ConsumeTraces(context.Background(), testdata.GenerateTraceDataOneSpanOneNil()))
	assert.NoError(t, lte.ConsumeTraces(context.Background(), testdata.GenerateTraceDataTwoSpansSameResourceOneDifferent()))

	assert.NoError(t, lte.Shutdown(context.Background()))
}

func TestLoggingMetricsExporterNoErrors(t *testing.T) {
	lme, err := NewMetricsExporter(&configmodels.ExporterSettings{}, "debug", zap.NewNop())
	require.NotNil(t, lme)
	assert.NoError(t, err)

	assert.NoError(t, lme.ConsumeMetrics(context.Background(), pdatautil.MetricsFromInternalMetrics(testdata.GenerateMetricDataEmpty())))
	assert.NoError(t, lme.ConsumeMetrics(context.Background(), pdatautil.MetricsFromInternalMetrics(testdata.GenerateMetricDataOneEmptyOneNilResourceMetrics())))
	assert.NoError(t, lme.ConsumeMetrics(context.Background(), pdatautil.MetricsFromInternalMetrics(testdata.GenerateMetricDataOneEmptyOneNilInstrumentationLibrary())))
	assert.NoError(t, lme.ConsumeMetrics(context.Background(), pdatautil.MetricsFromInternalMetrics(testdata.GenerateMetricDataOneMetricOneNil())))
	assert.NoError(t, lme.ConsumeMetrics(context.Background(), pdatautil.MetricsFromInternalMetrics(testdata.GenerateMetricDataWithCountersHistogramAndSummary())))

	assert.NoError(t, lme.Shutdown(context.Background()))
}
