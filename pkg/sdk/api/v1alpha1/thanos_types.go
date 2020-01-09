// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"github.com/banzaicloud/logging-operator/pkg/sdk/model/secret"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	thanosImageRepository = "quay.io/thanos/thanos"
	thanosImageTag        = "v0.9.0"
	defaultPullPolicy     = "pIfNotPresent"
)

var DefaultQuery = Query{
	BaseObject: BaseObject{
		Image: ImageSpec{
			Repository: thanosImageRepository,
			Tag:        thanosImageTag,
			PullPolicy: defaultPullPolicy,
		},
	},
	LogLevel:    "info",
	HttpAddress: "0.0.0.0:10902",
	GRPCAddress: "0.0.0.0:10901",
}

// ThanosSpec defines the desired state of Thanos
type ThanosSpec struct {
	Remote          *Remote          `json:"remote,omitempty"`
	ThanosDiscovery *ThanosDiscovery `json:"thanosDiscovery,omitempty"`
	Local           *Local           `json:"local,omitempty"`
	StoreGateway    *StoreGateway    `json:"storeGateway,omitempty"`
	Rule            *Rule            `json:"rule,omitempty"`
	ObjectStore     string           `json:"object_store,omitempty"`
	Query           *Query           `json:"query,omitempty"`
}

type Query struct {
	BaseObject `json:",inline"`
	LogLevel   string `json:"logLevel,omitempty"`
	LogFormat  string `json:"logFormat,omitempty"`
	// Listen host:port for HTTP endpoints.
	HttpAddress string `json:"httpAddress"`
	// Time to wait after an interrupt received for HTTP Server.
	HttpGracePeriod string `json:"http_grace_period"`
	// Listen ip:port address for gRPC endpoints
	GRPCAddress string `json:"grpcAddress"`
	// Time to wait after an interrupt received for GRPC Server.
	GRPCGracePeriod string `json:"grpcGracePeriod"`
	// Prefix for API and UI endpoints. This allows thanos UI to be served on a sub-path. This
	// option is analogous to --web.route-prefix of Promethus.
	WebRoutePrefix string `json:"webRoutePrefix,omitempty"`
	// Static prefix for all HTML links and redirect URLs in the UI query web interface. Actual
	// endpoints are still served on / or the web.route-prefix. This allows thanos UI to be
	// served behind a reverse proxy that strips a URL sub-path.
	WebExternalPrefix string `json:"webExternalPrefix,omitempty"`
	// Name of HTTP request header used for dynamic prefixing of UI links and redirects. This
	// option is ignored if web.external-prefix argument is set. Security risk: enable this
	// option only if a reverse proxy in front of thanos is resetting the header. The
	// --web.prefix-header=X-Forwarded-Prefix option can be useful, for example, if Thanos UI is
	// served via Traefik reverse proxy with PathPrefixStrip option enabled, which sends the
	// stripped prefix value in X-Forwarded-Prefix header. This allows thanos UI to be served on a
	// sub-path.
	WebPrefixHeader string `json:"webPrefixHeader,omitempty"`
	// Maximum time to process query by query node.
	QueryTimeout metav1.Duration `json:"queryTimeout,omitempty"`
	// Maximum number of queries processed concurrently by query node.
	QueryMaxConcurrent int `json:"queryMaxConcurrent,omitempty"`
	// Labels to treat as a replica indicator along which data is deduplicated. Still you will be
	// able to query without deduplication using 'dedup=false' parameter.
	QueryReplicaLabels map[string]string `json:"queryReplicaLabel,omitempty"`
	// Query selector labels that will be exposed in info endpoint (repeated).
	SelectorLabels map[string]string `json:"selectorLabels,omitempty"`
	// Addresses of statically configured store API servers (repeatable). The scheme may be
	// prefixed with 'dns+' or 'dnssrv+' to detect store API servers through respective DNS lookups.
	Stores []string `json:"stores,omitempty"`
	//	Interval between DNS resolutions.
	StoreSDDNSInterval metav1.Duration `json:"storeSDDNSInterval,omitempty"`
	//	Timeout before an unhealthy store is cleaned from the store UI page.
	StoreUnhealthyTimeout metav1.Duration `json:"storeUnhealthyTimeout,omitempty"`
	// Enable automatic adjustment (step / 5) to what source of data should be used in store gateways
	// if no max_source_resolution param is specified.
	QueryAutoDownsampling bool `json:"queryAutoDownsampling,omitempty"`
	// Enable partial response for queries if no partial_response param is specified.
	QueryPartialResponse bool `json:"queryPartialResponse,omitempty"`
	//	Set default evaluation interval for sub queries.
	QueryDefaultEvaluationInterval metav1.Duration `json:"queryDefaultEvaluationInterval,omitempty"`
	//	If a Store doesn't send any data in this specified duration then a Store will be ignored
	//	and partial data will be returned if it's enabled. 0 disables timeout.
	StoreResponseTimeout metav1.Duration `json:"storeResponseTimeout,omitempty"`
}

type TLS struct {
	Managed     ManagedTLS    `json:"managedTLS,omitempty"`
	Certificate secret.Secret `json:"certificate,omitempty"`
}

// TODO how the runtime generated certificate will work
type ManagedTLS struct {
}

type Remote struct {
	URLs []string `json:"urls"`
	TLS  TLS      `json:"tls"`
}

type Local struct {
	URLs []string `json:"urls"`
	TLS  TLS      `json:"tls"`
}

type ThanosDiscovery struct {
	LabelSelector metav1.LabelSelector `json:"labelSelector"`
}

type TimeRange struct {
	// Start of time range limit to serve. Thanos Store will serve only metrics, which happened
	// later than this value. Option can be a constant time in RFC3339 format or time duration
	// relative to current time, such as -1d or 2h45m. Valid duration units are ms, s, m, h, d, w, y.
	MinTime string `json:"minTime,omitempty"`
	// 	End of time range limit to serve. Thanos Store
	//	will serve only blocks, which happened eariler
	//	than this value. Option can be a constant time
	//	in RFC3339 format or time duration relative to
	//	current time, such as -1d or 2h45m. Valid
	//	duration units are ms, s, m, h, d, w, y.
	MaxTime string `json:"maxTime,omitempty"`
}

type StoreGateway struct {
	BaseObject `json:",inline"`
	LogLevel   string `json:"logLevel,omitempty"`
	LogFormat  string `json:"logFormat,omitempty"`
	// Listen host:port for HTTP endpoints.
	HttpAddress string `json:"httpAddress"`
	// Time to wait after an interrupt received for HTTP Server.
	HttpGracePeriod string `json:"http_grace_period"`
	// Listen ip:port address for gRPC endpoints
	GRPCAddress string `json:"grpcAddress"`
	// Time to wait after an interrupt received for GRPC Server.
	GRPCGracePeriod string `json:"grpcGracePeriod"`
	// Maximum size of items held in the in-memory index cache.
	IndexCacheSize string `json:"indexCacheSize"`
	// Maximum size of concurrently allocatable bytes for chunks.
	ChunkPoolSize string `json:"chunkPoolSize,omitempty"`
	// Maximum amount of samples returned via a single Series call. 0 means no limit. NOTE: For
	// efficiency we take 120 as the number of samples in chunk (it cannot be bigger than that), so
	// the actual number of samples might be lower, even though the maximum could be hit.
	StoreGRPCSeriesSampleLimit string `json:"storeGRPCSeriesSampleLimit,omitempty"`
	// Maximum number of concurrent Series calls.
	StoreGRPCSeriesMaxConcurrency int `json:"storeGRPCSeriesMaxConcurrency,omitempty"`
	// Repeat interval for syncing the blocks between local and remote view.
	SyncBlockDuration string `json:"syncBlockDuration,omitempty"`
	// Number of goroutines to use when constructing index-cache.json blocks from object storage.
	BlockSyncConcurrency int `json:"blockSyncConcurrency,omitempty"`
	// TimeRanges is a list of TimeRange to partition Store Gateway
	TimeRanges []TimeRange `json:"timeRanges,omitempty"`
}

type Rule struct {
	BaseObject `json:",inline"`
	LogLevel   string `json:"logLevel,omitempty"`
	LogFormat  string `json:"logFormat,omitempty"`
	// Listen host:port for HTTP endpoints.
	HttpAddress string `json:"httpAddress"`
	// Time to wait after an interrupt received for HTTP Server.
	HttpGracePeriod string `json:"http_grace_period"`
	// Listen ip:port address for gRPC endpoints
	GRPCAddress string `json:"grpcAddress"`
	// Time to wait after an interrupt received for GRPC Server.
	GRPCGracePeriod string `json:"grpcGracePeriod"`
	// 	Labels to be applied to all generated metrics
	//(repeated). Similar to external labels for
	//	Prometheus, used to identify ruler and its
	//	blocks as unique source.
	Labels map[string]string `json:"labels,omitempty"`
	// Rules
	Rules string `json:"rules,omitempty"`
	// Minimum amount of time to wait before resending an alert to Alertmanager.
	ResendDelay string `json:"resendDelay,omitempty"`
	// The default evaluation interval to use.
	EvalInterval string `json:"evalInterval,omitempty"`
	// Block duration for TSDB block.
	TSDBBlockDuration string `json:"tsdbBlockDuration,omitempty"`
	// Block retention time on local disk.
	TSDBRetention string `json:"tsdbRetention,omitempty"`
	// Alertmanager replica URLs to push firing alerts. Ruler claims success if push to at
	// least one alertmanager from discovered succeeds. The scheme should not be empty e.g
	// `http` might be used. The scheme may be prefixed with 'dns+' or 'dnssrv+' to detect
	// Alertmanager IPs through respective DNS lookups. The port defaults to 9093 or the SRV
	// record's value. The URL path is used as a prefix for the regular Alertmanager API path.
	AlertmanagersURLs []string `json:"alertmanagersURLs,omitempty"`
	// Timeout for sending alerts to Alertmanager
	AlertmanagersSendTimeout string `json:"alertmanagersSendTimeout,omitempty"`
	// Interval between DNS resolutions of Alertmanager hosts.
	AlertmanagersSDDNSInterval string `json:"alertmanagersSDDNSInterval,omitempty"`
	// The external Thanos Query URL that would be set in all alerts 'Source' field
	AlertQueryURL string `json:"alertQueryUrl,omitempty"`
	// Labels by name to drop before sending to alertmanager. This allows alert to be
	// deduplicated on replica label (repeated). Similar Prometheus alert relabelling
	AlertLabelDrop map[string]string `json:"alertLabelDrop,omitempty"`
	// Prefix for API and UI endpoints. This allows thanos UI to be served on a sub-path. This
	// option is analogous to --web.route-prefix of Promethus.
	WebRoutePrefix string `json:"webRoutePrefix,omitempty"`
	// Static prefix for all HTML links and redirect URLs in the UI query web interface. Actual
	// endpoints are still served on / or the web.route-prefix. This allows thanos UI to be
	// served behind a reverse proxy that strips a URL sub-path.
	WebExternalPrefix string `json:"webExternalPrefix,omitempty"`
	// Name of HTTP request header used for dynamic prefixing of UI links and redirects. This
	// option is ignored if web.external-prefix argument is set. Security risk: enable this
	// option only if a reverse proxy in front of thanos is resetting the header. The
	// --web.prefix-header=X-Forwarded-Prefix option can be useful, for example, if Thanos UI is
	// served via Traefik reverse proxy with PathPrefixStrip option enabled, which sends the
	// stripped prefix value in X-Forwarded-Prefix header. This allows thanos UI to be served on a
	// sub-path.
	WebPrefixHeader string `json:"webPrefixHeader,omitempty"`
	// Addresses of statically configured query API servers (repeatable). The scheme may be
	// prefixed with 'dns+' or 'dnssrv+' to detect query API servers through respective DNS
	// lookups.
	Queries []string `json:"queries"`
	// Interval between DNS resolutions.
	QuerySDDNSInterval string `json:"querySddnsInterval,omitempty"`
}

// ThanosStatus defines the observed state of Thanos
type ThanosStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// Thanos is the Schema for the thanos API
type Thanos struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ThanosSpec   `json:"spec,omitempty"`
	Status ThanosStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ThanosList contains a list of Thanos
type ThanosList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Thanos `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Thanos{}, &ThanosList{})
}