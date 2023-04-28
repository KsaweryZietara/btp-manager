package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sync"
)

// InstanceBindingControllerManager runs and stops the ServiceInstance controller
type InstanceBindingControllerManager struct {
	client.Client
	*rest.Config
	Scheme *runtime.Scheme

	sisbControllerMu sync.Mutex
	sisbReconciler   *ServiceInstanceReconciler
	mgr              ctrl.Manager
	stopper          func()
	enabled          bool
	cfg              *rest.Config
	ctx              context.Context
}

func NewInstanceBindingControllerManager(ctx context.Context, client client.Client, scheme *runtime.Scheme, cfg *rest.Config) *InstanceBindingControllerManager {
	return &InstanceBindingControllerManager{
		Client: client,
		Scheme: scheme,
		cfg:    cfg,
		ctx:    ctx,
	}
}

func (r *InstanceBindingControllerManager) EnableSISBController() {
	r.sisbControllerMu.Lock()
	defer r.sisbControllerMu.Unlock()
	logger := log.Log

	if r.enabled {
		return
	}
	mgr, err := ctrl.NewManager(r.cfg, ctrl.Options{
		Scheme:                 r.Scheme,
		MetricsBindAddress:     "0",
		HealthProbeBindAddress: "0",
	})
	if err != nil {
		logger.Error(err, "unable to create controller manager")
		return
	}
	r.mgr = mgr

	r.sisbReconciler = NewServiceInstanceReconciler(r.Client, r.Scheme)
	err = r.sisbReconciler.SetupWithManager(r.mgr)
	if err != nil {
		logger.Error(err, "unable to create SI SB controller")
		return
	}
	r.enabled = true

	_, cancel := context.WithCancel(r.ctx)
	r.stopper = cancel
	go func() {
		err = r.mgr.Start(r.ctx)
		logger.Error(err, "unable to start SI SB controller")
	}()

}

func (r *InstanceBindingControllerManager) DisableSISBController() {
	r.sisbControllerMu.Lock()
	defer r.sisbControllerMu.Unlock()
	if !r.enabled {
		return
	}

	if r.stopper != nil {
		r.stopper()
	}
	r.enabled = false

}