package admin

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	mesh_proto "github.com/kumahq/kuma/api/mesh/v1alpha1"
	system_proto "github.com/kumahq/kuma/api/system/v1alpha1"
	config_util "github.com/kumahq/kuma/pkg/config"
	config_cp "github.com/kumahq/kuma/pkg/config/app/kuma-cp"
	"github.com/kumahq/kuma/pkg/config/core/resources/store"
	"github.com/kumahq/kuma/pkg/core"
	core_mesh "github.com/kumahq/kuma/pkg/core/resources/apis/mesh"
	core_system "github.com/kumahq/kuma/pkg/core/resources/apis/system"
	"github.com/kumahq/kuma/pkg/core/resources/manager"
	core_model "github.com/kumahq/kuma/pkg/core/resources/model"
	core_store "github.com/kumahq/kuma/pkg/core/resources/store"
	"github.com/kumahq/kuma/pkg/kds/service"
	util_grpc "github.com/kumahq/kuma/pkg/util/grpc"
	"github.com/kumahq/kuma/pkg/util/k8s"
)

type kdsEnvoyAdminClient struct {
	rpcs       service.EnvoyAdminRPCs
	resManager manager.ReadOnlyResourceManager
}

func NewKDSEnvoyAdminClient(rpcs service.EnvoyAdminRPCs, resManager manager.ReadOnlyResourceManager) EnvoyAdminClient {
	return &kdsEnvoyAdminClient{
		rpcs:       rpcs,
		resManager: resManager,
	}
}

var _ EnvoyAdminClient = &kdsEnvoyAdminClient{}

func (k *kdsEnvoyAdminClient) PostQuit(context.Context, *core_mesh.DataplaneResource) error {
	panic("not implemented")
}

func (k *kdsEnvoyAdminClient) ConfigDump(ctx context.Context, proxy core_model.ResourceWithAddress) ([]byte, error) {
	zone := core_model.ZoneOfResource(proxy)
	nameInZone, err := resNameInZone(ctx, k.resManager, proxy)
	if err != nil {
		return nil, &KDSTransportError{requestType: "XDSConfigRequest", reason: err.Error()}
	}
	reqId := core.NewUUID()
	tenantZoneID := service.TenantZoneClientIDFromCtx(ctx, zone)

	err = k.rpcs.XDSConfigDump.Send(tenantZoneID.String(), &mesh_proto.XDSConfigRequest{
		RequestId:    reqId,
		ResourceType: string(proxy.Descriptor().Name),
		ResourceName: nameInZone,                // send the name which without the added prefix
		ResourceMesh: proxy.GetMeta().GetMesh(), // should be empty for ZoneIngress/ZoneEgress
	})
	if err != nil {
		return nil, &KDSTransportError{requestType: "XDSConfigRequest", reason: err.Error()}
	}

	defer k.rpcs.XDSConfigDump.DeleteWatch(tenantZoneID.String(), reqId)
	ch := make(chan util_grpc.ReverseUnaryMessage)
	if err := k.rpcs.XDSConfigDump.WatchResponse(tenantZoneID.String(), reqId, ch); err != nil {
		return nil, errors.Wrapf(err, "could not watch the response")
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case resp := <-ch:
		configResp, ok := resp.(*mesh_proto.XDSConfigResponse)
		if !ok {
			return nil, errors.New("invalid request type")
		}
		if configResp.GetError() != "" {
			return nil, &KDSTransportError{requestType: "XDSConfigRequest", reason: configResp.GetError()}
		}
		return configResp.GetConfig(), nil
	}
}

func (k *kdsEnvoyAdminClient) Stats(ctx context.Context, proxy core_model.ResourceWithAddress) ([]byte, error) {
	zone := core_model.ZoneOfResource(proxy)
	nameInZone, err := resNameInZone(ctx, k.resManager, proxy)
	if err != nil {
		return nil, &KDSTransportError{requestType: "StatsRequest", reason: err.Error()}
	}
	reqId := core.NewUUID()
	tenantZoneId := service.TenantZoneClientIDFromCtx(ctx, zone)

	err = k.rpcs.Stats.Send(tenantZoneId.String(), &mesh_proto.StatsRequest{
		RequestId:    reqId,
		ResourceType: string(proxy.Descriptor().Name),
		ResourceName: nameInZone,                // send the name which without the added prefix
		ResourceMesh: proxy.GetMeta().GetMesh(), // should be empty for ZoneIngress/ZoneEgress
	})
	if err != nil {
		return nil, &KDSTransportError{requestType: "StatsRequest", reason: err.Error()}
	}

	defer k.rpcs.Stats.DeleteWatch(tenantZoneId.String(), reqId)
	ch := make(chan util_grpc.ReverseUnaryMessage)
	if err := k.rpcs.Stats.WatchResponse(tenantZoneId.String(), reqId, ch); err != nil {
		return nil, errors.Wrapf(err, "could not watch the response")
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case resp := <-ch:
		statsResp, ok := resp.(*mesh_proto.StatsResponse)
		if !ok {
			return nil, errors.New("invalid request type")
		}
		if statsResp.GetError() != "" {
			return nil, &KDSTransportError{requestType: "StatsRequest", reason: statsResp.GetError()}
		}
		return statsResp.GetStats(), nil
	}
}

func (k *kdsEnvoyAdminClient) Clusters(ctx context.Context, proxy core_model.ResourceWithAddress) ([]byte, error) {
	zone := core_model.ZoneOfResource(proxy)
	nameInZone, err := resNameInZone(ctx, k.resManager, proxy)
	if err != nil {
		return nil, &KDSTransportError{requestType: "ClustersRequest", reason: err.Error()}
	}
	reqId := core.NewUUID()
	tenantZoneID := service.TenantZoneClientIDFromCtx(ctx, zone)

	err = k.rpcs.Clusters.Send(tenantZoneID.String(), &mesh_proto.ClustersRequest{
		RequestId:    reqId,
		ResourceType: string(proxy.Descriptor().Name),
		ResourceName: nameInZone,                // send the name which without the added prefix
		ResourceMesh: proxy.GetMeta().GetMesh(), // should be empty for ZoneIngress/ZoneEgress
	})
	if err != nil {
		return nil, &KDSTransportError{requestType: "ClustersRequest", reason: err.Error()}
	}

	defer k.rpcs.Clusters.DeleteWatch(tenantZoneID.String(), reqId)
	ch := make(chan util_grpc.ReverseUnaryMessage)
	if err := k.rpcs.Clusters.WatchResponse(tenantZoneID.String(), reqId, ch); err != nil {
		return nil, errors.Wrapf(err, "could not watch the response")
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case resp := <-ch:
		clustersResp, ok := resp.(*mesh_proto.ClustersResponse)
		if !ok {
			return nil, errors.New("invalid request type")
		}
		if clustersResp.GetError() != "" {
			return nil, &KDSTransportError{requestType: "ClustersRequest", reason: clustersResp.GetError()}
		}
		return clustersResp.GetClusters(), nil
	}
}

func resNameInZone(
	ctx context.Context,
	resManager manager.ReadOnlyResourceManager,
	r core_model.Resource,
) (string, error) {
	name := core_model.GetDisplayName(r)
	zone := core_model.ZoneOfResource(r)
	// we need to check for the legacy name which starts with zoneName
	if strings.HasPrefix(r.GetMeta().GetName(), zone) {
		return name, nil
	}
	storeType, err := getZoneStoreType(ctx, resManager, zone)
	if err != nil {
		return "", err
	}
	// only K8s needs namespace added to the resource name
	if storeType != store.KubernetesStore {
		return name, nil
	}

	if ns := r.GetMeta().GetLabels()[mesh_proto.KubeNamespaceTag]; ns != "" {
		name = k8s.K8sNamespacedNameToCoreName(name, ns)
	}
	return name, nil
}

func getZoneStoreType(
	ctx context.Context,
	resManager manager.ReadOnlyResourceManager,
	zone string,
) (store.StoreType, error) {
	zoneInsightRes := core_system.NewZoneInsightResource()
	if err := resManager.Get(ctx, zoneInsightRes, core_store.GetByKey(zone, core_model.NoMesh)); err != nil {
		return "", err
	}
	subscription := zoneInsightRes.Spec.GetLastSubscription()
	if !subscription.IsOnline() {
		return "", fmt.Errorf("zone is offline")
	}
	kdsSubscription, ok := subscription.(*system_proto.KDSSubscription)
	if !ok {
		return "", fmt.Errorf("cannot map subscription")
	}
	config := kdsSubscription.GetConfig()
	cfg := &config_cp.Config{}
	if err := config_util.FromYAML([]byte(config), cfg); err != nil {
		return "", fmt.Errorf("cannot read control-plane configuration")
	}
	return cfg.Store.Type, nil
}

type KDSTransportError struct {
	requestType string
	reason      string
}

func (e *KDSTransportError) Error() string {
	if e.reason == "" {
		return fmt.Sprintf("could not send %s", e.requestType)
	} else {
		return fmt.Sprintf("could not send %s: %s", e.requestType, e.reason)
	}
}

func (e *KDSTransportError) Is(err error) bool {
	return reflect.TypeOf(e) == reflect.TypeOf(err)
}
