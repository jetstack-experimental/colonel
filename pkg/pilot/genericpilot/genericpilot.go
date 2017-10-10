package genericpilot

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	"github.com/jetstack-experimental/navigator/pkg/client/clientset_generated/clientset"
	listersv1alpha1 "github.com/jetstack-experimental/navigator/pkg/client/listers_generated/navigator/v1alpha1"
	"github.com/jetstack-experimental/navigator/pkg/pilot/genericpilot/process"
)

type GenericPilot struct {
	Options Options

	client clientset.Interface

	pilotLister         listersv1alpha1.PilotLister
	pilotInformerSynced cache.InformerSynced

	queue    workqueue.RateLimitingInterface
	recorder record.EventRecorder

	process process.Interface
}

// only run one worker to prevent threading issues when dealing with processes
const workers = 1

func (g *GenericPilot) enqueuePilot(obj interface{}) {
	key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}
	g.queue.AddRateLimited(key)
}

func (g *GenericPilot) Run() error {
	glog.Infof("Starting Elasticsearch controller")

	if !cache.WaitForCacheSync(g.Options.StopCh, g.pilotInformerSynced) {
		return fmt.Errorf("timed out waiting for caches to sync")
	}

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			wait.Until(g.worker, time.Second, g.Options.StopCh)
		}()
	}

	<-g.Options.StopCh
	g.queue.ShutDown()
	glog.V(4).Infof("Shutting down Elasticsearch controller workers...")
	wg.Wait()
	glog.V(4).Infof("Elasticsearch controller workers stopped.")
	return nil
}