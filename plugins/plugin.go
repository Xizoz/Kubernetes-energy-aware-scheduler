package plugins

import (
	"context"
	"math"
	"math/rand"
	"strings"

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
	var nodeCpuTotal int64 = nodeInfo.Allocatable.MilliCPU
	var nodeMemoryTotal int64 = nodeInfo.Allocatable.Memory

	var nodeCpuInUse int64 = nodeInfo.Requested.MilliCPU
	var nodeMemoryInUse int64 = nodeInfo.Requested.Memory

	var nodeCpuInUsePercentage = (float64(nodeCpuInUse) / float64(nodeCpuTotal)) * 100
	nodeCpuInUsePercentageRounded := math.Round(nodeCpuInUsePercentage*100) / 100

	nodeMemoryInUsePercentage := (float64(nodeMemoryInUse) / float64(nodeMemoryTotal)) * 100
	nodeMemoryInUsePercentageRounded := math.Round(nodeMemoryInUsePercentage*100) / 100

	if nodeCpuInUsePercentageRounded > 90 || nodeMemoryInUsePercentageRounded > 90 {
		return framework.NewStatus(framework.Error, "Node overhead is too big")
	}

	if nodeInfo.Node() == nil {
		return framework.NewStatus(framework.Error, "Node not found")
	}
	return framework.NewStatus(framework.Success)
}

func (p *EnergyEfficientPlugin) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	energyUsage := rand.Int63n(100)
	score := 100 - energyUsage

	var podStatus v1.PodStatus = pod.Status
	var podPhase string = string(podStatus.Phase)

	if !strings.Contains(podPhase, "Pending") {
		klog.Info("Pod was already scheduled or something went wrong!")
		return 0, framework.NewStatus(framework.Pending)
	}
	var podName string = pod.Name
	var podPriority *int32 = pod.Spec.Priority
	var podSchedulerNameAffilation string = pod.Spec.SchedulerName

	// Prepare function
	// callKubeRLBridge(nodeName, podName, podPriority, podSchedulerNameAffilation )
	// callKubeRLBridge("energy-aware-k8-cluster-worker", "test-pod-7", 1, "energy-scheduler")

	klog.Infof("Scoring node logs:", podName, podPriority, podSchedulerNameAffilation)
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
