package builder

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	backupdriverv1api "github.com/vmware-tanzu/velero-plugin-for-vsphere/pkg/apis/backupdriver/v1"
)

// BackupRepositoryClaimBuilder builds Upload objects
type BackupRepositoryClaimBuilder struct {
	object *backupdriverv1api.BackupRepositoryClaim
}

// ForBackupRepositoryClaim is the constructor for a BackupRepositoryClaimBuilder.
func ForBackupRepositoryClaim(ns, name string) *BackupRepositoryClaimBuilder {
	return &BackupRepositoryClaimBuilder{
		object: &backupdriverv1api.BackupRepositoryClaim{
			TypeMeta: metav1.TypeMeta{
				APIVersion: backupdriverv1api.SchemeGroupVersion.String(),
				Kind:       "BackupRepositoryClaim",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: ns,
				Name:      name,
			},
		},
	}
}

// Result returns the built BackupRepositoryClaim.
func (b *BackupRepositoryClaimBuilder) Result() *backupdriverv1api.BackupRepositoryClaim {
	return b.object
}

// ObjectMeta applies functional options to the BackupRepositoryClaim's ObjectMeta.
func (b *BackupRepositoryClaimBuilder) ObjectMeta(opts ...ObjectMetaOpt) *BackupRepositoryClaimBuilder {
	for _, opt := range opts {
		opt(b.object)
	}

	return b
}

func (b *BackupRepositoryClaimBuilder) RepositoryDriver(repositorydriver string) *BackupRepositoryClaimBuilder {
	b.object.RepositoryDriver = repositorydriver
	return b
}

func (b *BackupRepositoryClaimBuilder) BackupRepository(backuprepository string) *BackupRepositoryClaimBuilder {
	b.object.BackupRepository = backuprepository
	return b
}

func (b *BackupRepositoryClaimBuilder) AllowedNamespaces(namespaces []string) *BackupRepositoryClaimBuilder {
	b.object.AllowedNamespaces = namespaces
	return b
}

func (b *BackupRepositoryClaimBuilder) RepositoryParameters(params map[string]string) *BackupRepositoryClaimBuilder {
	b.object.RepositoryParameters = params
	return b
}