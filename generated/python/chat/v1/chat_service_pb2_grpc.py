# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from chat.v1 import chat_service_pb2 as chat_dot_v1_dot_chat__service__pb2

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
        + f' but the generated code in chat/v1/chat_service_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class ChatStub(object):
    """Deprecated: Use the v2 service
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetChats = channel.unary_unary(
                '/code.chat.v1.Chat/GetChats',
                request_serializer=chat_dot_v1_dot_chat__service__pb2.GetChatsRequest.SerializeToString,
                response_deserializer=chat_dot_v1_dot_chat__service__pb2.GetChatsResponse.FromString,
                _registered_method=True)
        self.GetMessages = channel.unary_unary(
                '/code.chat.v1.Chat/GetMessages',
                request_serializer=chat_dot_v1_dot_chat__service__pb2.GetMessagesRequest.SerializeToString,
                response_deserializer=chat_dot_v1_dot_chat__service__pb2.GetMessagesResponse.FromString,
                _registered_method=True)
        self.AdvancePointer = channel.unary_unary(
                '/code.chat.v1.Chat/AdvancePointer',
                request_serializer=chat_dot_v1_dot_chat__service__pb2.AdvancePointerRequest.SerializeToString,
                response_deserializer=chat_dot_v1_dot_chat__service__pb2.AdvancePointerResponse.FromString,
                _registered_method=True)
        self.SetMuteState = channel.unary_unary(
                '/code.chat.v1.Chat/SetMuteState',
                request_serializer=chat_dot_v1_dot_chat__service__pb2.SetMuteStateRequest.SerializeToString,
                response_deserializer=chat_dot_v1_dot_chat__service__pb2.SetMuteStateResponse.FromString,
                _registered_method=True)
        self.SetSubscriptionState = channel.unary_unary(
                '/code.chat.v1.Chat/SetSubscriptionState',
                request_serializer=chat_dot_v1_dot_chat__service__pb2.SetSubscriptionStateRequest.SerializeToString,
                response_deserializer=chat_dot_v1_dot_chat__service__pb2.SetSubscriptionStateResponse.FromString,
                _registered_method=True)
        self.StreamChatEvents = channel.stream_stream(
                '/code.chat.v1.Chat/StreamChatEvents',
                request_serializer=chat_dot_v1_dot_chat__service__pb2.StreamChatEventsRequest.SerializeToString,
                response_deserializer=chat_dot_v1_dot_chat__service__pb2.StreamChatEventsResponse.FromString,
                _registered_method=True)
        self.SendMessage = channel.unary_unary(
                '/code.chat.v1.Chat/SendMessage',
                request_serializer=chat_dot_v1_dot_chat__service__pb2.SendMessageRequest.SerializeToString,
                response_deserializer=chat_dot_v1_dot_chat__service__pb2.SendMessageResponse.FromString,
                _registered_method=True)


class ChatServicer(object):
    """Deprecated: Use the v2 service
    """

    def GetChats(self, request, context):
        """GetChats gets the set of chats for an owner account
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMessages(self, request, context):
        """GetMessages gets the set of messages for a chat
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AdvancePointer(self, request, context):
        """AdvancePointer advances a pointer in chat history
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetMuteState(self, request, context):
        """SetMuteState configures the mute state of a chat
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetSubscriptionState(self, request, context):
        """SetSubscriptionState configures the susbscription state of a chat
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def StreamChatEvents(self, request_iterator, context):
        """
        Experimental PoC two-way chat APIs below


        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendMessage(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ChatServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetChats': grpc.unary_unary_rpc_method_handler(
                    servicer.GetChats,
                    request_deserializer=chat_dot_v1_dot_chat__service__pb2.GetChatsRequest.FromString,
                    response_serializer=chat_dot_v1_dot_chat__service__pb2.GetChatsResponse.SerializeToString,
            ),
            'GetMessages': grpc.unary_unary_rpc_method_handler(
                    servicer.GetMessages,
                    request_deserializer=chat_dot_v1_dot_chat__service__pb2.GetMessagesRequest.FromString,
                    response_serializer=chat_dot_v1_dot_chat__service__pb2.GetMessagesResponse.SerializeToString,
            ),
            'AdvancePointer': grpc.unary_unary_rpc_method_handler(
                    servicer.AdvancePointer,
                    request_deserializer=chat_dot_v1_dot_chat__service__pb2.AdvancePointerRequest.FromString,
                    response_serializer=chat_dot_v1_dot_chat__service__pb2.AdvancePointerResponse.SerializeToString,
            ),
            'SetMuteState': grpc.unary_unary_rpc_method_handler(
                    servicer.SetMuteState,
                    request_deserializer=chat_dot_v1_dot_chat__service__pb2.SetMuteStateRequest.FromString,
                    response_serializer=chat_dot_v1_dot_chat__service__pb2.SetMuteStateResponse.SerializeToString,
            ),
            'SetSubscriptionState': grpc.unary_unary_rpc_method_handler(
                    servicer.SetSubscriptionState,
                    request_deserializer=chat_dot_v1_dot_chat__service__pb2.SetSubscriptionStateRequest.FromString,
                    response_serializer=chat_dot_v1_dot_chat__service__pb2.SetSubscriptionStateResponse.SerializeToString,
            ),
            'StreamChatEvents': grpc.stream_stream_rpc_method_handler(
                    servicer.StreamChatEvents,
                    request_deserializer=chat_dot_v1_dot_chat__service__pb2.StreamChatEventsRequest.FromString,
                    response_serializer=chat_dot_v1_dot_chat__service__pb2.StreamChatEventsResponse.SerializeToString,
            ),
            'SendMessage': grpc.unary_unary_rpc_method_handler(
                    servicer.SendMessage,
                    request_deserializer=chat_dot_v1_dot_chat__service__pb2.SendMessageRequest.FromString,
                    response_serializer=chat_dot_v1_dot_chat__service__pb2.SendMessageResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'code.chat.v1.Chat', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('code.chat.v1.Chat', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Chat(object):
    """Deprecated: Use the v2 service
    """

    @staticmethod
    def GetChats(request,
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
            '/code.chat.v1.Chat/GetChats',
            chat_dot_v1_dot_chat__service__pb2.GetChatsRequest.SerializeToString,
            chat_dot_v1_dot_chat__service__pb2.GetChatsResponse.FromString,
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
    def GetMessages(request,
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
            '/code.chat.v1.Chat/GetMessages',
            chat_dot_v1_dot_chat__service__pb2.GetMessagesRequest.SerializeToString,
            chat_dot_v1_dot_chat__service__pb2.GetMessagesResponse.FromString,
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
    def AdvancePointer(request,
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
            '/code.chat.v1.Chat/AdvancePointer',
            chat_dot_v1_dot_chat__service__pb2.AdvancePointerRequest.SerializeToString,
            chat_dot_v1_dot_chat__service__pb2.AdvancePointerResponse.FromString,
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
    def SetMuteState(request,
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
            '/code.chat.v1.Chat/SetMuteState',
            chat_dot_v1_dot_chat__service__pb2.SetMuteStateRequest.SerializeToString,
            chat_dot_v1_dot_chat__service__pb2.SetMuteStateResponse.FromString,
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
    def SetSubscriptionState(request,
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
            '/code.chat.v1.Chat/SetSubscriptionState',
            chat_dot_v1_dot_chat__service__pb2.SetSubscriptionStateRequest.SerializeToString,
            chat_dot_v1_dot_chat__service__pb2.SetSubscriptionStateResponse.FromString,
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
    def StreamChatEvents(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(
            request_iterator,
            target,
            '/code.chat.v1.Chat/StreamChatEvents',
            chat_dot_v1_dot_chat__service__pb2.StreamChatEventsRequest.SerializeToString,
            chat_dot_v1_dot_chat__service__pb2.StreamChatEventsResponse.FromString,
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
    def SendMessage(request,
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
            '/code.chat.v1.Chat/SendMessage',
            chat_dot_v1_dot_chat__service__pb2.SendMessageRequest.SerializeToString,
            chat_dot_v1_dot_chat__service__pb2.SendMessageResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
