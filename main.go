package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/nikimoro-qlik/k8s_playground/pkg/log"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	podNamespace              = apiv1.NamespaceDefault
	podDeletionCostAnnotation = "controller.kubernetes.io/pod-deletion-cost"
)

type metaStruct struct {
	Metadata struct {
		Annotations map[string]string `json:"annotations,omitempty"`
	} `json:"metadata,omitempty"`
}

var (
	podName             = os.Getenv("HOSTNAME")
	currentDeletionCost = 0
)

func main() {
	log.GetLogger().Infow("Starting application", "podName", podName)

	// init context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// intercept quit signals
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quitCh
		cancel()
	}()

	// get k8s client
	config, err := rest.InClusterConfig()
	if err != nil {
		log.GetLogger().Errorw("Not running in cluster", "podName", podName, "err", err)
		os.Exit(1)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.GetLogger().Errorw("error creating kubernetes client", "podName", podName, "err", err)
		os.Exit(1)
	}
	podService := client.CoreV1().Pods(podNamespace)

	// main loop
	ticker := time.NewTicker(time.Second * time.Duration(15))
	defer ticker.Stop()
main_loop:
	for {
		select {
		case <-ticker.C:
			// generate new random deletion cost
			currentDeletionCost = rand.Intn(100)

			// update annotations
			var newMeta metaStruct
			newMeta.Metadata.Annotations = map[string]string{
				podDeletionCostAnnotation: strconv.Itoa(currentDeletionCost),
			}
			newMetaBytes, err := json.Marshal(newMeta)
			if err != nil {
				log.GetLogger().Errorw("Error marshalling annotations", "podName", podName, "err", err)
				break
			}

			// update pod
			_, err = podService.Patch(ctx, podName, types.MergePatchType, newMetaBytes, metav1.PatchOptions{})
			if err != nil {
				log.GetLogger().Errorw("Error updating pod annotations", "podName", podName, "err", err)
				break
			}
			log.GetLogger().Infow("Pod annotation updated", "podName", podName, "deletionCost", currentDeletionCost)

		case <-ctx.Done():
			break main_loop
		}
	}

	log.GetLogger().Infow("Bye bye", "podName", podName, "deletionCost", currentDeletionCost)
}
