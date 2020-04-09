package e2e

import (
	"fmt"
	"github.com/vmware-tanzu/velero/pkg/builder"
	"github.com/vmware-tanzu/velero/pkg/generated/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"testing"
	"time"
)

const DefaultBackupTTL time.Duration = 30 * 24 * time.Hour

func Test_Basic(t *testing.T) {
	path := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		t.Fatal("Got error " + err.Error())
	}
	veleroClient, err := versioned.NewForConfig(config)
	if err != nil {
		t.Fatal("Got error " + err.Error())
	}

	// create
	backupBuilder := builder.ForBackup("velero", "demo-app-test")
	backupBuilder.
		IncludedNamespaces("demo-app").
		VolumeSnapshotLocations("vsl-vsphere").
		SnapshotVolumes(true)
	backup := backupBuilder.Result()
	_, err = veleroClient.VeleroV1().Backups(backup.Namespace).Create(backup)
	fmt.Printf("Backup request %q submitted successfully.\n", backup.Name)

	// delete
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		t.Fatal("Got error " + err.Error())
	}
	clientset.CoreV1().Namespaces().Delete("demo-app", metav1.NewDeleteOptions(0))

	// restore
	restoreBuilder := builder.ForRestore("velero", "demo-app-restore-test")
	restoreBuilder.Backup("demo-app-04-06-2")
	restore := restoreBuilder.Result()
	_, err = veleroClient.VeleroV1().Restores(restore.Namespace).Create(restore)
}