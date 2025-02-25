package main

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const PluginName = "EnergyEfficientScheduler"
const PrometheusURL = "http://prometheus.monitoring:9090/api/v1/query"

type CustomSchedulerPlugin struct {
	handle framework.Handle
}

var _ framework.ScorePlugin = &CustomSchedulerPlugin{}

func (p *CustomSchedulerPlugin) Name() string {
	return PluginName
}

func main() {
	framework.Register(PluginName, func(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
		return &CustomSchedulerPlugin{handle: h}, nil
	})
}
