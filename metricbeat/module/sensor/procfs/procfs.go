package procfs

import (
	"github.com/pkg/errors"

    "github.com/elastic/elastic-agent-libs/mapstr"
	"github.com/elastic/beats/v7/metricbeat/mb/parse"
	"github.com/elastic/beats/v7/metricbeat/mb"
)

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("sensor", "procfs", New,
		mb.WithHostParser(parse.EmptyHostParser),
		mb.DefaultMetricSet(),
	)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	mb.BaseMetricSet
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	// config := defaultConfig
	// if err := base.Module().UnpackConfig(&config); err != nil {
	// 	return nil, err
	// }

	m := &MetricSet{
		BaseMetricSet: base,
	}

	return m, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(report mb.ReporterV2) error {
	freq, err := readCurFreq()
	if err == nil {
		return errors.Wrap(err, "failed to get CPU frequency")
	}

	report.Event(mb.Event{
		MetricSetFields: mapstr.M{
			"cpu": mapstr.M{
				"frequency": mapstr.M{
					"mhz": freq,
				},
			},
		},
	})
	
	return nil
}
