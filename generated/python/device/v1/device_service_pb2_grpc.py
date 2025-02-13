# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from device.v1 import device_service_pb2 as device_dot_v1_dot_device__service__pb2

GRPC_GENERATED_VERSION = '1.70.0'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in device/v1/device_service_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class DeviceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RegisterLoggedInAccounts = channel.unary_unary(
                '/code.device.v1.Device/RegisterLoggedInAccounts',
                request_serializer=device_dot_v1_dot_device__service__pb2.RegisterLoggedInAccountsRequest.SerializeToString,
                response_deserializer=device_dot_v1_dot_device__service__pb2.RegisterLoggedInAccountsResponse.FromString,
                _registered_method=True)
        self.GetLoggedInAccounts = channel.unary_unary(
                '/code.device.v1.Device/GetLoggedInAccounts',
                request_serializer=device_dot_v1_dot_device__service__pb2.GetLoggedInAccountsRequest.SerializeToString,
                response_deserializer=device_dot_v1_dot_device__service__pb2.GetLoggedInAccountsResponse.FromString,
                _registered_method=True)


class DeviceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def RegisterLoggedInAccounts(self, request, context):
        """RegisterLoggedInAccounts registers a set of owner accounts logged for
        an app install. Currently, a single login is enforced per app install.
        After using GetLoggedInAccounts to detect stale logins, clients can use
        this RPC to update the set of accounts with valid login sessions.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetLoggedInAccounts(self, request, context):
        """GetLoggedInAccounts gets the set of logged in accounts for an app install.
        Clients can use this RPC to detect stale logins for boot out of the app.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DeviceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RegisterLoggedInAccounts': grpc.unary_unary_rpc_method_handler(
                    servicer.RegisterLoggedInAccounts,
                    request_deserializer=device_dot_v1_dot_device__service__pb2.RegisterLoggedInAccountsRequest.FromString,
                    response_serializer=device_dot_v1_dot_device__service__pb2.RegisterLoggedInAccountsResponse.SerializeToString,
            ),
            'GetLoggedInAccounts': grpc.unary_unary_rpc_method_handler(
                    servicer.GetLoggedInAccounts,
                    request_deserializer=device_dot_v1_dot_device__service__pb2.GetLoggedInAccountsRequest.FromString,
                    response_serializer=device_dot_v1_dot_device__service__pb2.GetLoggedInAccountsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'code.device.v1.Device', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('code.device.v1.Device', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Device(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def RegisterLoggedInAccounts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/code.device.v1.Device/RegisterLoggedInAccounts',
            device_dot_v1_dot_device__service__pb2.RegisterLoggedInAccountsRequest.SerializeToString,
            device_dot_v1_dot_device__service__pb2.RegisterLoggedInAccountsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetLoggedInAccounts(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/code.device.v1.Device/GetLoggedInAccounts',
            device_dot_v1_dot_device__service__pb2.GetLoggedInAccountsRequest.SerializeToString,
            device_dot_v1_dot_device__service__pb2.GetLoggedInAccountsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
