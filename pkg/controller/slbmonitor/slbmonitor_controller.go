package slbmonitor

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"

	middlewarev1alpha1 "github.com/riete/aliyun-monitor-operator/pkg/apis/middleware/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_slbmonitor")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new SlbMonitor Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileSlbMonitor{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("slbmonitor-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource SlbMonitor
	err = c.Watch(&source.Kind{Type: &middlewarev1alpha1.SlbMonitor{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner SlbMonitor
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &middlewarev1alpha1.SlbMonitor{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileSlbMonitor implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileSlbMonitor{}

// ReconcileSlbMonitor reconciles a SlbMonitor object
type ReconcileSlbMonitor struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a SlbMonitor object and makes changes based on the state read
// and what is in the SlbMonitor.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileSlbMonitor) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling SlbMonitor")

	// Fetch the SlbMonitor instance
	instance := &middlewarev1alpha1.SlbMonitor{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	for _, account := range instance.Spec.SlbAliyunAccount {

		// Define a new Deployment object
		exporter := newDeployment(account)
		name := fmt.Sprintf("%s-%s-slb-exporter", account.Name, account.RegionId)
		// Set SlbMonitor instance as the owner and controller
		if err := controllerutil.SetControllerReference(instance, exporter, r.scheme); err != nil {
			return reconcile.Result{}, err
		}

		// Check if this Deployment already exists
		found := &appsv1.Deployment{}
		err = r.client.Get(context.TODO(), types.NamespacedName{Name: exporter.Name, Namespace: exporter.Namespace}, found)
		if err != nil && errors.IsNotFound(err) {
			reqLogger.Info("Creating a new Deployment", "Deployment.Namespace", exporter.Namespace, "Deployment.Name", name)
			err = r.client.Create(context.TODO(), exporter)
			if err != nil {
				return reconcile.Result{}, err
			}
		} else if err != nil {
			return reconcile.Result{}, err
		} else {
			found.Spec.Template.Spec.Containers[0].Env = []corev1.EnvVar{
				{Name: "ACCESS_KEY_ID", Value: account.AccessKeyId},
				{Name: "ACCESS_KEY_SECRET", Value: account.AccessKeySecret},
				{Name: "REGION_ID", Value: account.RegionId},
				{Name: "INSTANCE_ID", Value: strings.Join(account.InstanceId, ",")},
			}
			err = r.client.Update(context.TODO(), found)
			if err != nil {
				reqLogger.Error(err, "Failed to update Deployment", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
				return reconcile.Result{}, err
			}
			reqLogger.Info("Update Deployment", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
		}

		// Deployment already exists - don't requeue
		reqLogger.Info("Skip reconcile: Deployment already exists", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
	}
	return reconcile.Result{}, nil
}

func newDeployment(account middlewarev1alpha1.SlbAliyunAccount) *appsv1.Deployment {
	var replicas int32 = 1
	cpuRequest, _ := resource.ParseQuantity("10m")
	cpuLimit, _ := resource.ParseQuantity("50m")
	memoryRequest, _ := resource.ParseQuantity("10Mi")
	memoryLimit, _ := resource.ParseQuantity("50Mi")
	name := fmt.Sprintf("%s-%s-slb-exporter", account.Name, account.RegionId)
	labels := map[string]string{"app": name}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: middlewarev1alpha1.SlbExporterNamespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
					Annotations: map[string]string{
						"prometheus.io/path":   "/metrics",
						"prometheus.io/port":   strconv.Itoa(int(middlewarev1alpha1.SlbExporterPort)),
						"prometheus.io/scrape": "true",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Image:           middlewarev1alpha1.SlbExporterImage,
						Name:            name,
						ImagePullPolicy: corev1.PullAlways,
						Env: []corev1.EnvVar{
							{Name: "ACCESS_KEY_ID", Value: account.AccessKeyId},
							{Name: "ACCESS_KEY_SECRET", Value: account.AccessKeySecret},
							{Name: "REGION_ID", Value: account.RegionId},
							{Name: "INSTANCE_ID", Value: strings.Join(account.InstanceId, ",")},
						},
						Ports: []corev1.ContainerPort{{
							ContainerPort: middlewarev1alpha1.SlbExporterPort,
							Name:          "slb-exporter",
							Protocol:      corev1.ProtocolTCP,
						}},
						LivenessProbe: &corev1.Probe{
							FailureThreshold:    3,
							InitialDelaySeconds: 10,
							PeriodSeconds:       10,
							SuccessThreshold:    1,
							TimeoutSeconds:      3,
							Handler: corev1.Handler{
								TCPSocket: &corev1.TCPSocketAction{
									Port: intstr.FromInt(int(middlewarev1alpha1.SlbExporterPort)),
								},
							},
						},
						ReadinessProbe: &corev1.Probe{
							FailureThreshold:    3,
							InitialDelaySeconds: 10,
							PeriodSeconds:       10,
							SuccessThreshold:    1,
							TimeoutSeconds:      3,
							Handler: corev1.Handler{
								TCPSocket: &corev1.TCPSocketAction{
									Port: intstr.FromInt(int(middlewarev1alpha1.SlbExporterPort)),
								},
							},
						},
						Resources: corev1.ResourceRequirements{
							Requests: corev1.ResourceList{
								corev1.ResourceCPU:    cpuRequest,
								corev1.ResourceMemory: memoryRequest,
							},
							Limits: corev1.ResourceList{
								corev1.ResourceCPU:    cpuLimit,
								corev1.ResourceMemory: memoryLimit,
							},
						},
					}},
				},
			},
		},
	}
}
