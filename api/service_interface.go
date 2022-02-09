package client

import (
	"context"
	"net/http"

	client "github.com/Nordix/go-redfish/client"
)

//go:generate mockery -name=RedfishAPI -output ./mocks
type RedfishAPI interface {
	GetManager(context.Context,
		string,
	) client.ApiGetManagerRequest

	GetManagerExecute(client.ApiGetManagerRequest,
	) (*client.Manager,
		*http.Response,
		error,
	)

	GetManagerVirtualMedia(context.Context,
		string,
		string,
	) client.ApiGetManagerVirtualMediaRequest

	GetManagerVirtualMediaExecute(client.ApiGetManagerVirtualMediaRequest,
	) (*client.VirtualMedia,
		*http.Response,
		error,
	)

	GetRoot(context.Context,
	) client.ApiGetRootRequest

	GetRootExecute(client.ApiGetRootRequest,
	) (*client.Root,
		*http.Response,
		error,
	)

	GetSystem(context.Context,
		string,
	) client.ApiGetSystemRequest

	GetSystemExecute(client.ApiGetSystemRequest,
	) (*client.ComputerSystem,
		*http.Response,
		error,
	)

	InsertVirtualMedia(context.Context,
		string,
		string,
	) client.ApiInsertVirtualMediaRequest

	InsertVirtualMediaExecute(client.ApiInsertVirtualMediaRequest,
	) (*client.RedfishError,
		*http.Response,
		error,
	)

	ListManagerVirtualMedia(context.Context,
		string,
	) client.ApiListManagerVirtualMediaRequest

	ListManagerVirtualMediaExecute(client.ApiListManagerVirtualMediaRequest,
	) (*client.Collection,
		*http.Response,
		error,
	)

	ListManagers(context.Context,
	) client.ApiListManagersRequest

	ListManagersExecute(client.ApiListManagersRequest,
	) (*client.Collection,
		*http.Response,
		error,
	)

	ListSystems(context.Context,
	) client.ApiListSystemsRequest

	ListSystemsExecute(client.ApiListSystemsRequest,
	) (*client.Collection,
		*http.Response,
		error,
	)

	ResetSystem(context.Context,
		string,
	) client.ApiResetSystemRequest

	ResetSystemExecute(client.ApiResetSystemRequest,
	) (*client.RedfishError,
		*http.Response,
		error,
	)

	SetSystem(context.Context,
		string,
	) client.ApiSetSystemRequest

	SetSystemExecute(client.ApiSetSystemRequest,
	) (*client.ComputerSystem,
		*http.Response,
		error,
	)
}
