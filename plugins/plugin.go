package plugins

import (
	"context"
	"math/rand"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
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
	plugin := &EnergyEfficientPlugin{handle: handle}
	return plugin, nil
}

func (p *EnergyEfficientPlugin) Name() string {
	return PluginName
}

func (p *EnergyEfficientPlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.Infof("Filtering pod %s on node %s", pod.Name, nodeInfo.Node().Name)
	if nodeInfo.Node() == nil {
		return framework.NewStatus(framework.Error, "Node not found")
	}
	return framework.NewStatus(framework.Success)
}

func (p *EnergyEfficientPlugin) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	energyUsage := rand.Int63n(100)
	score := 100 - energyUsage
	klog.Infof("Scoring node %s: energy usage %d, final score %d", nodeName, energyUsage, score)
	return score, framework.NewStatus(framework.Success)
}

func (p *EnergyEfficientPlugin) ScoreExtensions() framework.ScoreExtensions {
	return nil
}

func (p *EnergyEfficientPlugin) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) (*framework.PreFilterResult, *framework.Status) {
	klog.Infof("PreFilter called for pod %s", pod.Name)
	return nil, framework.NewStatus(framework.Success)
}

func (p *EnergyEfficientPlugin) PreFilterExtensions() framework.PreFilterExtensions {
	return nil
}
