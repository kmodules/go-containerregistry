/*
Copyright AppsCode Inc. and Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmds

import (
	"fmt"

	"kmodules.xyz/go-containerregistry/authn"

	"github.com/google/go-containerregistry/pkg/authn/k8schain"
	"github.com/spf13/cobra"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
)

func NewRootCmd() *cobra.Command {
	var (
		image   string
		k8sOpts k8schain.Options
	)
	cmd := &cobra.Command{
		Use:               "echo-image-digest [command]",
		Short:             "Print docker image digest",
		Long:              "Print docker image digest",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctrl.SetLogger(klogr.New())
			cfg := ctrl.GetConfigOrDie()
			cfg.QPS = 100
			cfg.Burst = 100
			kc, err := kubernetes.NewForConfig(cfg)
			if err != nil {
				return err
			}
			result, err := authn.ImageWithDigest(kc, image, []k8schain.Options{k8sOpts})
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}
	cmd.Flags().StringVar(&image, "image", image, "Image name")
	cmd.Flags().StringVar(&k8sOpts.Namespace, "namespace", k8sOpts.Namespace, "Pod namespace")
	cmd.Flags().StringVar(&k8sOpts.ServiceAccountName, "service-account-name", k8sOpts.ServiceAccountName, "Pod service account name")
	cmd.Flags().StringSliceVar(&k8sOpts.ImagePullSecrets, "image-pull-secrets", k8sOpts.ImagePullSecrets, "Name of image pull secret")
	authn.AddInsecureRegistriesFlag(cmd.Flags())

	return cmd
}
