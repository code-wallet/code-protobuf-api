// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts"
// @generated from file push/v1/push_service.proto (package code.push.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddTokenRequest, AddTokenResponse, RemoveTokenRequest, RemoveTokenResponse } from "./push_service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service code.push.v1.Push
 */
export const Push = {
  typeName: "code.push.v1.Push",
  methods: {
    /**
     * AddToken stores a push token in a data container. The call is idempotent
     * and adding an existing valid token will not fail. Token types will be
     * validated against the user agent and any mismatches will result in an
     * INVALID_ARGUMENT status error.
     *
     * The token will be unlinked from any and all other accounts that it was
     * previously bound to.
     *
     * @generated from rpc code.push.v1.Push.AddToken
     */
    addToken: {
      name: "AddToken",
      I: AddTokenRequest,
      O: AddTokenResponse,
      kind: MethodKind.Unary,
    },
    /**
     * RemoveToken removes the provided push token from the account.
     *
     * The provided token must be bound to the current account.
     * Otherwise, the RPC will succeed with without removal.
     *
     * @generated from rpc code.push.v1.Push.RemoveToken
     */
    removeToken: {
      name: "RemoveToken",
      I: RemoveTokenRequest,
      O: RemoveTokenResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

