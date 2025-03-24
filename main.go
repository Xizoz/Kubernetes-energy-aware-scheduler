package main

import (
	"context"
	"math/rand"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	app "k8s.io/kubernetes/cmd/kube-scheduler/app"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const PluginName = "EnergyEfficientScheduler"

type EnergyEfficientPlugin struct {
	handle framework.Handle
}

var _ framework.FilterPlugin = &EnergyEfficientPlugin{}
var _ framework.ScorePlugin = &EnergyEfficientPlugin{}
var _ framework.PreFilterPlugin = &EnergyEfficientPlugin{}

func New(ctx context.Context, obj runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	klog.Info("Initializing Energy Efficient Plugin...")
	klog.Infof("Context: %+v", ctx)
	klog.Infof("Object: %+v", obj)
	klog.Infof("Plugin handle: %+v", handle)

	plugin := &EnergyEfficientPlugin{handle: handle}
	klog.Infof("Plugin created: %+v", plugin)
	return plugin, nil
}

func (p *EnergyEfficientPlugin) Name() string {
	return PluginName
}

func (p *EnergyEfficientPlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.Infof("EnergyEfficientPlugin: Filtering pod %s on node %s", pod.Name, nodeInfo.Node().Name)
	klog.Flush()
	node := nodeInfo.Node()
	if node == nil {
		return framework.NewStatus(framework.Error, "Node not found")
	}
	return framework.NewStatus(framework.Success)

}

func (p *EnergyEfficientPlugin) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.Infof("EnergyEfficientPlugin: Scoring pod %s on node %s", pod.Name, nodeName)
	klog.Flush()

	energyUsage := rand.Int63n(100) // Simulate energy efficiency (lower is better)
	score := 100 - energyUsage      // Higher score is better
	klog.Infof("Scoring node %s with energy usage %d, final score: %d", nodeName, energyUsage, score)
	klog.Flush()

	return score, framework.NewStatus(framework.Success)
}

func (p *EnergyEfficientPlugin) ScoreExtensions() framework.ScoreExtensions {
	klog.Info("EnergyEfficientPlugin: ScoreExtensions called")
	klog.Flush()

	return nil
}

func (p *EnergyEfficientPlugin) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) (*framework.PreFilterResult, *framework.Status) {
	klog.Infof("EnergyEfficientPlugin: PreFilter called for pod %s", pod.Name)
	return nil, framework.NewStatus(framework.Success)
}

func (p *EnergyEfficientPlugin) PreFilterExtensions() framework.PreFilterExtensions {
	klog.Info("EnergyEfficientPlugin: PreFilterExtensions called")
	return nil
}
func main() {
	klog.Info("Starting Energy Efficient Scheduler Plugin...")
	defer klog.Flush()

	command := app.NewSchedulerCommand(
		app.WithPlugin(PluginName, New),
	)

	klog.Info("Scheduler command created, executing...")

	if err := command.Execute(); err != nil {
		klog.Fatal(err)
	}

	select {} // Keep running
}
