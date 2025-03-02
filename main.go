package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/angao/scheduler-framework-sample/pkg/plugins/sample"
	"k8s.io/component-base/logs"
	"k8s.io/klog"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	logs.InitLogs()
	defer logs.FlushLogs()

	klog.Info("Starting Custom Kubernetes Scheduler")

	cmd := app.NewSchedulerCommand(
		app.WithPlugin(sample.Name, sample.New),
	)

	klog.Info("Scheduler command initialized, executing...")

	if err := cmd.Execute(); err != nil {
		klog.Errorf("Scheduler failed: %v", err)
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	klog.Info("Scheduler execution completed")
}
