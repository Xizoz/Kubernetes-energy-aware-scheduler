package main

import (
	"context"
	"math/rand"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"

	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const PluginName = "EnergyEfficientScheduler"

type EnergyEfficientPlugin struct {
	handle framework.Handle
}

var _ framework.ScorePlugin = &EnergyEfficientPlugin{}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	klog.Info("Plugin runs!")
	return &EnergyEfficientPlugin{handle: h}, nil
}

func (p *EnergyEfficientPlugin) Name() string {
	return PluginName
}
func (p *EnergyEfficientPlugin) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	// Simulate energy efficiency score (lower energy usage is better)
	energyUsage := rand.Int63n(100) // Fake energy metric (0-99)
	score := 100 - energyUsage      // Higher score is better
	klog.Infof("Scoring node %s with energy usage %d, final score: %d", nodeName, energyUsage, score)
	return score, framework.NewStatus(framework.Success)
}

func (p *EnergyEfficientPlugin) ScoreExtensions() framework.ScoreExtensions {
	return nil
}

func loadFakePods() {
	klog.Info("Loading fake pods for simulation2...")
	fakePods := []string{"pod-1", "pod-2", "pod-3"}
	for _, pod := range fakePods {
		time.Sleep(1 * time.Second)
		klog.Infof("Fake Pod Loaded: %s", pod)
	}
}

func main() {
	klog.Info("Starting Energy Efficient Scheduler Plugin...")
	go loadFakePods()
	defer klog.Flush()
	command := app.NewSchedulerCommand(
		app.WithPlugin(PluginName, New),
	)

	klog.Info("Scheduler command created, executing...")

	// Run the scheduler
	if err := command.Execute(); err != nil {
		klog.ErrorS(err, "Failed to run scheduler")
	}

	select {} // Keep running
}
