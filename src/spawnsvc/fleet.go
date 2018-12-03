package spawnsvc

import (
	"context"
	"fmt"
	"log"
	"strconv"

	fleet "github.com/proepkes/speeddate/src/spawnsvc/gen/fleet"

	"agones.dev/agones/pkg/apis/stable/v1alpha1"
	"agones.dev/agones/pkg/client/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// fleet service example implementation.
// The example methods log the requests and return zero values.
type fleetSvc struct {
	logger              *log.Logger
	speeddateNamespace  string
	gameserverNamespace string
	k8sClient           kubernetes.Interface
	client              *versioned.Clientset
}

// NewFleet returns the fleet service implementation.
func NewFleet(logger *log.Logger, speeddateNamespace string, gameserverNamespace string, clusterConfig *rest.Config) fleet.Service {

	k8sClient, err := kubernetes.NewForConfig(clusterConfig)
	if err != nil {
		logger.Fatal("Could not create the k8s clientset")
	}

	// Access to the Agones resources through the Agones Clientset
	// Note that we reuse the same config as we used for the Kubernetes Clientset
	client, err := versioned.NewForConfig(clusterConfig)
	if err != nil {
		logger.Fatal("Could not create the agones api clientset")
	}

	return &fleetSvc{logger, speeddateNamespace, gameserverNamespace, k8sClient, client}
}

// Add a new gameserver.
func (s *fleetSvc) Add(ctx context.Context) (res string, err error) {
	s.logger.Print("fleet.add")

	cm, err := s.getGameserverConfig()
	if err != nil {
		return "", err
	}

	cp, err := strconv.ParseInt(cm.Data["ContainerPort"], 10, 32)
	pp := v1alpha1.PortPolicy(cm.Data["PortPolicy"])

	// Create a GameServer
	gs := &v1alpha1.GameServer{ObjectMeta: metav1.ObjectMeta{GenerateName: cm.Data["GameserverNamePrefix"], Namespace: cm.Data["GameserverNamespace"]},
		Spec: v1alpha1.GameServerSpec{
			Ports: []v1alpha1.GameServerPort{{ContainerPort: int32(cp), PortPolicy: pp}},
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{Name: cm.Data["ContainerName"], Image: cm.Data["ContainerImage"]}},
				},
			},
		},
	}
	newGS, err := s.client.StableV1alpha1().GameServers(s.gameserverNamespace).Create(gs)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New game servers' name is: %s", newGS.ObjectMeta.Name)

	return
}

// Removes all gameserver pods.
func (s *fleetSvc) Clear(ctx context.Context) (res string, err error) {
	s.logger.Print("fleet.clear")
	s.client.StableV1alpha1().GameServers(s.gameserverNamespace).DeleteCollection(&metav1.DeleteOptions{}, metav1.ListOptions{})
	return
}

// Configure gameserver-properties.
func (s *fleetSvc) Configure(ctx context.Context) (res string, err error) {
	s.logger.Print("fleet.configure")
	cm, err := s.getGameserverConfig()
	if err != nil {
		return "", err
	}
	fmt.Println(cm)
	return
}

func (s *fleetSvc) getGameserverConfig() (*corev1.ConfigMap, error) {
	configMaps, err := s.k8sClient.CoreV1().ConfigMaps(s.speeddateNamespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, cm := range configMaps.Items {
		if cm.Name == "gameserver-config" {
			return &cm, nil
		}
	}

	return nil, fmt.Errorf("Could not find gameserver-config in namespace %s", s.speeddateNamespace)
}