package resources

import (
	"github.com/loft-sh/vcluster/pkg/mappings/generic"
	"github.com/loft-sh/vcluster/pkg/syncer/synccontext"
	"github.com/loft-sh/vcluster/pkg/util/translate"
	corev1 "k8s.io/api/core/v1"
)

func CreateSecretsMapper(ctx *synccontext.RegisterContext) (synccontext.Mapper, error) {
	return generic.NewMapper(ctx, &corev1.Secret{}, translate.Default.HostName)
}
