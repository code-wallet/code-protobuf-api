// @generated by protoc-gen-connect-es v0.13.0 with parameter "target=ts"
// @generated from file transaction/v2/transaction_service.proto (package code.transaction.v2, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AirdropRequest, AirdropResponse, CanWithdrawToAccountRequest, CanWithdrawToAccountResponse, DeclareFiatOnrampPurchaseAttemptRequest, DeclareFiatOnrampPurchaseAttemptResponse, GetIntentMetadataRequest, GetIntentMetadataResponse, GetLimitsRequest, GetLimitsResponse, GetPrioritizedIntentsForPrivacyUpgradeRequest, GetPrioritizedIntentsForPrivacyUpgradeResponse, GetPrivacyUpgradeStatusRequest, GetPrivacyUpgradeStatusResponse, SubmitIntentRequest, SubmitIntentResponse, SwapRequest, SwapResponse } from "./transaction_service_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service code.transaction.v2.Transaction
 */
export const Transaction = {
  typeName: "code.transaction.v2.Transaction",
  methods: {
    /**
     * SubmitIntent is the mechanism for client and server to agree upon a set of
     * client actions to execute on the blockchain using the Code sequencer for
     * fulfillment.
     *
     * Transactions and virtual instructions are never exchanged between client and server.
     * Instead, the required accounts and arguments for instructions known to each actor are
     * exchanged to allow independent and local construction.
     *
     * Client and server are expected to fully validate the intent. Proofs will
     * be provided for any parameter requiring one. Signatures should only be
     * generated after approval.
     *
     * This RPC is not a traditional streaming endpoint. It bundles two unary calls
     * to enable DB-level transaction semantics.
     *
     * The high-level happy path flow for the RPC is as follows:
     *   1.  Client initiates a stream and sends SubmitIntentRequest.SubmitActions
     *   2.  Server validates the intent, its actions and metadata
     *   3a. If there are transactions or virtual instructions requiring the user's signature,
     *       then server returns SubmitIntentResponse.ServerParameters
     *   3b. Otherwise, server returns SubmitIntentResponse.Success and closes the
     *       stream
     *   4.  For each transaction or virtual instruction requiring the user's signature, the client
     *       locally constructs it, performs validation and collects the signature
     *   5.  Client sends SubmitIntentRequest.SubmitSignatures with the signature
     *       list generated from 4
     *   6.  Server validates all signatures are submitted and are the expected values
     *       using locally constructed transactions or virtual instructions.
     *   7.  Server returns SubmitIntentResponse.Success and closes the stream
     * In the error case:
     *   * Server will return SubmitIntentResponse.Error and close the stream
     *   * Client will close the stream
     *
     * @generated from rpc code.transaction.v2.Transaction.SubmitIntent
     */
    submitIntent: {
      name: "SubmitIntent",
      I: SubmitIntentRequest,
      O: SubmitIntentResponse,
      kind: MethodKind.BiDiStreaming,
    },
    /**
     * GetIntentMetadata gets basic metadata on an intent. It can also be used
     * to fetch the status of submitted intents. Metadata exists only for intents
     * that have been successfully submitted.
     *
     * @generated from rpc code.transaction.v2.Transaction.GetIntentMetadata
     */
    getIntentMetadata: {
      name: "GetIntentMetadata",
      I: GetIntentMetadataRequest,
      O: GetIntentMetadataResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetPrivacyUpgradeStatus gets the status of a private transaction and the
     * ability to upgrade it to permanent privacy.
     *
     * @generated from rpc code.transaction.v2.Transaction.GetPrivacyUpgradeStatus
     */
    getPrivacyUpgradeStatus: {
      name: "GetPrivacyUpgradeStatus",
      I: GetPrivacyUpgradeStatusRequest,
      O: GetPrivacyUpgradeStatusResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetPrioritizedIntentsForPrivacyUpgrade allows clients to get private
     * intent actions that can be upgraded in a secure and verifiable manner.
     *
     * @generated from rpc code.transaction.v2.Transaction.GetPrioritizedIntentsForPrivacyUpgrade
     */
    getPrioritizedIntentsForPrivacyUpgrade: {
      name: "GetPrioritizedIntentsForPrivacyUpgrade",
      I: GetPrioritizedIntentsForPrivacyUpgradeRequest,
      O: GetPrioritizedIntentsForPrivacyUpgradeResponse,
      kind: MethodKind.Unary,
    },
    /**
     * GetLimits gets limits for money moving intents for an owner account in an
     * identity-aware manner
     *
     * @generated from rpc code.transaction.v2.Transaction.GetLimits
     */
    getLimits: {
      name: "GetLimits",
      I: GetLimitsRequest,
      O: GetLimitsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * CanWithdrawToAccount provides hints to clients for submitting withdraw intents.
     * The RPC indicates if a withdrawal is possible, and how it should be performed.
     *
     * @generated from rpc code.transaction.v2.Transaction.CanWithdrawToAccount
     */
    canWithdrawToAccount: {
      name: "CanWithdrawToAccount",
      I: CanWithdrawToAccountRequest,
      O: CanWithdrawToAccountResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Airdrop airdrops Kin to the requesting account
     *
     * @generated from rpc code.transaction.v2.Transaction.Airdrop
     */
    airdrop: {
      name: "Airdrop",
      I: AirdropRequest,
      O: AirdropResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Swap performs an on-chain swap. The high-level flow mirrors SubmitIntent
     * closely. However, due to the time-sensitive nature and unreliability of
     * swaps, they do not fit within the broader intent system. This results in
     * a few key differences:
     *  * Transactions are submitted on a best-effort basis outside of the Code
     *    Sequencer within the RPC handler
     *  * Balance changes are applied after the transaction has finalized
     *  * Transactions use recent blockhashes over a nonce
     * SubmitIntent also operates on VM virtual instructions, whereas Swap uses
     * Solana transactions.
     *
     * The transaction will have the following instruction format:
     *   1. ComputeBudget::SetComputeUnitLimit
     *   2. ComputeBudget::SetComputeUnitPrice
     *   3. SwapValidator::PreSwap
     *   4. Dynamic swap instruction
     *   5. SwapValidator::PostSwap
     *
     * Note: Currently limited to swapping USDC to Kin.
     * Note: Kin is deposited into the token account derived from the VM deposit PDA of the owner account.
     *
     * @generated from rpc code.transaction.v2.Transaction.Swap
     */
    swap: {
      name: "Swap",
      I: SwapRequest,
      O: SwapResponse,
      kind: MethodKind.BiDiStreaming,
    },
    /**
     * DeclareFiatOnrampPurchaseAttempt is called whenever a user attempts to use a fiat
     * onramp to purchase crypto for use in Code.
     *
     * @generated from rpc code.transaction.v2.Transaction.DeclareFiatOnrampPurchaseAttempt
     */
    declareFiatOnrampPurchaseAttempt: {
      name: "DeclareFiatOnrampPurchaseAttempt",
      I: DeclareFiatOnrampPurchaseAttemptRequest,
      O: DeclareFiatOnrampPurchaseAttemptResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

