package scheduler

import (
    "github.com/feel-easy/gpushare-scheduler-extender/pkg/cache"
    "k8s.io/client-go/kubernetes"
)

func NewGPUsharePredicate(clientset *kubernetes.Clientset, c *cache.SchedulerCache) *Predicate {
	return &Predicate{Name: "gpusharingfilter", cache: c}
}
