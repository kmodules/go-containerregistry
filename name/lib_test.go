/*
Copyright AppsCode Inc. and Contributors

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

package name

import (
	"reflect"
	"testing"

	"github.com/google/go-containerregistry/pkg/name"
)

func TestParseReference(t *testing.T) {
	type args struct {
		s    string
		opts []name.Option
	}
	tests := []struct {
		name    string
		args    args
		want    *Image
		wantErr bool
	}{
		{
			name: "Test with valid image string",
			args: args{
				s: "docker.io/library/nginx@sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
			},
			want: &Image{
				Original:   "docker.io/library/nginx@sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
				Name:       "index.docker.io/library/nginx@sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
				Registry:   "index.docker.io",
				Repository: "library/nginx",
				Tag:        "",
				Digest:     "sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
			},
			wantErr: false,
		},
		{
			name: "Test with valid image string",
			args: args{
				s: "docker.io/library/mariadb@sha256:8040983db146f729749081c6b216a19d52e0973134e2e34c0b4fd87f48bc15b0",
			},
			want: &Image{
				Original:   "docker.io/library/mariadb@sha256:8040983db146f729749081c6b216a19d52e0973134e2e34c0b4fd87f48bc15b0",
				Name:       "index.docker.io/library/mariadb@sha256:8040983db146f729749081c6b216a19d52e0973134e2e34c0b4fd87f48bc15b0",
				Registry:   "index.docker.io",
				Repository: "library/mariadb",
				Tag:        "",
				Digest:     "sha256:8040983db146f729749081c6b216a19d52e0973134e2e34c0b4fd87f48bc15b0",
			},
			wantErr: false,
		},
		{
			name: "Test with valid image string with tag",
			args: args{
				s: "docker.io/library/mariadb:10.8.2@sha256:8040983db146f729749081c6b216a19d52e0973134e2e34c0b4fd87f48bc15b0",
			},
			want: &Image{
				Original:   "docker.io/library/mariadb:10.8.2@sha256:8040983db146f729749081c6b216a19d52e0973134e2e34c0b4fd87f48bc15b0",
				Name:       "index.docker.io/library/mariadb@sha256:8040983db146f729749081c6b216a19d52e0973134e2e34c0b4fd87f48bc15b0",
				Registry:   "index.docker.io",
				Repository: "library/mariadb",
				Tag:        "10.8.2",
				Digest:     "sha256:8040983db146f729749081c6b216a19d52e0973134e2e34c0b4fd87f48bc15b0",
			},
			wantErr: false,
		},
		{
			name: "Test with valid custom image string",
			args: args{
				s: "docker.io/prom/mysqld-exporter@sha256:a8af600c3ef1c8df179b736b94d04dc5ec209be88407a4c1c1bd0fc6394f56e8",
			},
			want: &Image{
				Original:   "docker.io/prom/mysqld-exporter@sha256:a8af600c3ef1c8df179b736b94d04dc5ec209be88407a4c1c1bd0fc6394f56e8",
				Name:       "index.docker.io/prom/mysqld-exporter@sha256:a8af600c3ef1c8df179b736b94d04dc5ec209be88407a4c1c1bd0fc6394f56e8",
				Registry:   "index.docker.io",
				Repository: "prom/mysqld-exporter",
				Tag:        "",
				Digest:     "sha256:a8af600c3ef1c8df179b736b94d04dc5ec209be88407a4c1c1bd0fc6394f56e8",
			},
			wantErr: false,
		},
		{
			name: "Test with gcr image string",
			args: args{
				s: "gcr.io/google-containers/echoserver:1.10@sha256:cb5c1bddd1b5665e1867a7fa1b5fa843a47ee433bbb75d4293888b71def53229",
			},
			want: &Image{
				Original:   "gcr.io/google-containers/echoserver:1.10@sha256:cb5c1bddd1b5665e1867a7fa1b5fa843a47ee433bbb75d4293888b71def53229",
				Name:       "gcr.io/google-containers/echoserver@sha256:cb5c1bddd1b5665e1867a7fa1b5fa843a47ee433bbb75d4293888b71def53229",
				Registry:   "gcr.io",
				Repository: "google-containers/echoserver",
				Tag:        "1.10",
				Digest:     "sha256:cb5c1bddd1b5665e1867a7fa1b5fa843a47ee433bbb75d4293888b71def53229",
			},
			wantErr: false,
		},
		{
			name: "Test with valid image string without sha256 hash",
			args: args{
				s: "gcr.io/google-containers/echoserver:1.10",
			},
			want: &Image{
				Original:   "gcr.io/google-containers/echoserver:1.10",
				Name:       "gcr.io/google-containers/echoserver:1.10",
				Registry:   "gcr.io",
				Repository: "google-containers/echoserver",
				Tag:        "1.10",
				Digest:     "",
			},
			wantErr: false,
		},
		{
			name: "Test with valid image string without sha256 hash",
			args: args{
				s: "kubedb/ui-server:1.10-v2",
			},
			want: &Image{
				Original:   "kubedb/ui-server:1.10-v2",
				Name:       "index.docker.io/kubedb/ui-server:1.10-v2",
				Registry:   "index.docker.io",
				Repository: "kubedb/ui-server",
				Tag:        "1.10-v2",
				Digest:     "",
			},
			wantErr: false,
		},
		{
			name: "Test with valid image string with sha256 hash",
			args: args{
				s: "docker.io/kubedb/postgres-init@sha256:48aeb8859d8245d82788836918c817d1ff84dfc9e81612f3ee8e9e22b6e3e8a5",
			},
			want: &Image{
				Original:   "docker.io/kubedb/postgres-init@sha256:48aeb8859d8245d82788836918c817d1ff84dfc9e81612f3ee8e9e22b6e3e8a5",
				Name:       "index.docker.io/kubedb/postgres-init@sha256:48aeb8859d8245d82788836918c817d1ff84dfc9e81612f3ee8e9e22b6e3e8a5",
				Registry:   "index.docker.io",
				Repository: "kubedb/postgres-init",
				Tag:        "",
				Digest:     "sha256:48aeb8859d8245d82788836918c817d1ff84dfc9e81612f3ee8e9e22b6e3e8a5",
			},
			wantErr: false,
		},
		{
			name: "Test with official image without sha256 hash",
			args: args{
				s: "nginx:1.0.1",
			},
			want: &Image{
				Original:   "nginx:1.0.1",
				Name:       "index.docker.io/library/nginx:1.0.1",
				Registry:   "index.docker.io",
				Repository: "library/nginx",
				Tag:        "1.0.1",
				Digest:     "",
			},
			wantErr: false,
		},
		{
			name: "Test with official image with sha256 hash",
			args: args{
				s: "nginx:1.0.1@sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
			},
			want: &Image{
				Original:   "nginx:1.0.1@sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
				Name:       "index.docker.io/library/nginx@sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
				Registry:   "index.docker.io",
				Repository: "library/nginx",
				Tag:        "1.0.1",
				Digest:     "sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
			},
			wantErr: false,
		},
		{
			name: "Test with official image with latest tag and sha256 hash",
			args: args{
				s: "nginx:latest@sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
			},
			want: &Image{
				Original:   "nginx:latest@sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
				Name:       "index.docker.io/library/nginx@sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
				Registry:   "index.docker.io",
				Repository: "library/nginx",
				Tag:        "latest",
				Digest:     "sha256:d586384381a0e6834cef73d432b1486f0b86334cb92e54256def62dd403f82ab",
			},
			wantErr: false,
		},
		{
			name: "registry.k8s.io",
			args: args{
				s: "registry.k8s.io/ingress-nginx/controller:v1.9.3",
			},
			want: &Image{
				Original:   "registry.k8s.io/ingress-nginx/controller:v1.9.3",
				Name:       "registry.k8s.io/ingress-nginx/controller:v1.9.3",
				Registry:   "registry.k8s.io",
				Repository: "ingress-nginx/controller",
				Tag:        "v1.9.3",
				Digest:     "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseReference(tt.args.s, tt.args.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseReference() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseReference() got = %v, want %v", got, tt.want)
			}
		})
	}
}
