<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2\UpgradeableIntent;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.transaction.v2.UpgradeableIntent.UpgradeablePrivateAction</code>
 */
class UpgradeablePrivateAction extends \Google\Protobuf\Internal\Message
{
    /**
     * The transaction blob that was signed by the client. Clients *MUST* use
     * the source and destination account in the timelock::TransferWithAuthority
     * instruction to validate all fields provided by server by locally computing
     * the expected addresses.
     *
     * Generated from protobuf field <code>.code.common.v1.Transaction transaction_blob = 1 [(.validate.rules) = {</code>
     */
    protected $transaction_blob = null;
    /**
     * The client's signature for the transaction. Clients MUST use this to
     * locally validate the transaction blob provided by server.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature client_signature = 2 [(.validate.rules) = {</code>
     */
    protected $client_signature = null;
    /**
     * The action ID of this transaction
     *
     * Generated from protobuf field <code>uint32 action_id = 3;</code>
     */
    protected $action_id = 0;
    /**
     * The source account's type, which hints how to efficiently derive source
     *
     * Generated from protobuf field <code>.code.common.v1.AccountType source_account_type = 4 [(.validate.rules) = {</code>
     */
    protected $source_account_type = 0;
    /**
     * The source account's derivation index, which hints how to efficiently derive source
     *
     * Generated from protobuf field <code>uint64 source_derivation_index = 5;</code>
     */
    protected $source_derivation_index = 0;
    /**
     * The original destination account that was paid by the treasury
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId original_destination = 6 [(.validate.rules) = {</code>
     */
    protected $original_destination = null;
    /**
     * The original quark amount for the action
     *
     * Generated from protobuf field <code>uint64 original_amount = 7 [(.validate.rules) = {</code>
     */
    protected $original_amount = 0;
    /**
     * The treasury used for this the private action
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId treasury = 8 [(.validate.rules) = {</code>
     */
    protected $treasury = null;
    /**
     * The recent root observed at the time of intent creation for this private action
     *
     * Generated from protobuf field <code>.code.common.v1.Hash recent_root = 9 [(.validate.rules) = {</code>
     */
    protected $recent_root = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\Transaction $transaction_blob
     *           The transaction blob that was signed by the client. Clients *MUST* use
     *           the source and destination account in the timelock::TransferWithAuthority
     *           instruction to validate all fields provided by server by locally computing
     *           the expected addresses.
     *     @type \Code\Common\V1\Signature $client_signature
     *           The client's signature for the transaction. Clients MUST use this to
     *           locally validate the transaction blob provided by server.
     *     @type int $action_id
     *           The action ID of this transaction
     *     @type int $source_account_type
     *           The source account's type, which hints how to efficiently derive source
     *     @type int|string $source_derivation_index
     *           The source account's derivation index, which hints how to efficiently derive source
     *     @type \Code\Common\V1\SolanaAccountId $original_destination
     *           The original destination account that was paid by the treasury
     *     @type int|string $original_amount
     *           The original quark amount for the action
     *     @type \Code\Common\V1\SolanaAccountId $treasury
     *           The treasury used for this the private action
     *     @type \Code\Common\V1\Hash $recent_root
     *           The recent root observed at the time of intent creation for this private action
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * The transaction blob that was signed by the client. Clients *MUST* use
     * the source and destination account in the timelock::TransferWithAuthority
     * instruction to validate all fields provided by server by locally computing
     * the expected addresses.
     *
     * Generated from protobuf field <code>.code.common.v1.Transaction transaction_blob = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\Transaction|null
     */
    public function getTransactionBlob()
    {
        return $this->transaction_blob;
    }

    public function hasTransactionBlob()
    {
        return isset($this->transaction_blob);
    }

    public function clearTransactionBlob()
    {
        unset($this->transaction_blob);
    }

    /**
     * The transaction blob that was signed by the client. Clients *MUST* use
     * the source and destination account in the timelock::TransferWithAuthority
     * instruction to validate all fields provided by server by locally computing
     * the expected addresses.
     *
     * Generated from protobuf field <code>.code.common.v1.Transaction transaction_blob = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\Transaction $var
     * @return $this
     */
    public function setTransactionBlob($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\Transaction::class);
        $this->transaction_blob = $var;

        return $this;
    }

    /**
     * The client's signature for the transaction. Clients MUST use this to
     * locally validate the transaction blob provided by server.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature client_signature = 2 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\Signature|null
     */
    public function getClientSignature()
    {
        return $this->client_signature;
    }

    public function hasClientSignature()
    {
        return isset($this->client_signature);
    }

    public function clearClientSignature()
    {
        unset($this->client_signature);
    }

    /**
     * The client's signature for the transaction. Clients MUST use this to
     * locally validate the transaction blob provided by server.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature client_signature = 2 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\Signature $var
     * @return $this
     */
    public function setClientSignature($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\Signature::class);
        $this->client_signature = $var;

        return $this;
    }

    /**
     * The action ID of this transaction
     *
     * Generated from protobuf field <code>uint32 action_id = 3;</code>
     * @return int
     */
    public function getActionId()
    {
        return $this->action_id;
    }

    /**
     * The action ID of this transaction
     *
     * Generated from protobuf field <code>uint32 action_id = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setActionId($var)
    {
        GPBUtil::checkUint32($var);
        $this->action_id = $var;

        return $this;
    }

    /**
     * The source account's type, which hints how to efficiently derive source
     *
     * Generated from protobuf field <code>.code.common.v1.AccountType source_account_type = 4 [(.validate.rules) = {</code>
     * @return int
     */
    public function getSourceAccountType()
    {
        return $this->source_account_type;
    }

    /**
     * The source account's type, which hints how to efficiently derive source
     *
     * Generated from protobuf field <code>.code.common.v1.AccountType source_account_type = 4 [(.validate.rules) = {</code>
     * @param int $var
     * @return $this
     */
    public function setSourceAccountType($var)
    {
        GPBUtil::checkEnum($var, \Code\Common\V1\AccountType::class);
        $this->source_account_type = $var;

        return $this;
    }

    /**
     * The source account's derivation index, which hints how to efficiently derive source
     *
     * Generated from protobuf field <code>uint64 source_derivation_index = 5;</code>
     * @return int|string
     */
    public function getSourceDerivationIndex()
    {
        return $this->source_derivation_index;
    }

    /**
     * The source account's derivation index, which hints how to efficiently derive source
     *
     * Generated from protobuf field <code>uint64 source_derivation_index = 5;</code>
     * @param int|string $var
     * @return $this
     */
    public function setSourceDerivationIndex($var)
    {
        GPBUtil::checkUint64($var);
        $this->source_derivation_index = $var;

        return $this;
    }

    /**
     * The original destination account that was paid by the treasury
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId original_destination = 6 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getOriginalDestination()
    {
        return $this->original_destination;
    }

    public function hasOriginalDestination()
    {
        return isset($this->original_destination);
    }

    public function clearOriginalDestination()
    {
        unset($this->original_destination);
    }

    /**
     * The original destination account that was paid by the treasury
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId original_destination = 6 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setOriginalDestination($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->original_destination = $var;

        return $this;
    }

    /**
     * The original quark amount for the action
     *
     * Generated from protobuf field <code>uint64 original_amount = 7 [(.validate.rules) = {</code>
     * @return int|string
     */
    public function getOriginalAmount()
    {
        return $this->original_amount;
    }

    /**
     * The original quark amount for the action
     *
     * Generated from protobuf field <code>uint64 original_amount = 7 [(.validate.rules) = {</code>
     * @param int|string $var
     * @return $this
     */
    public function setOriginalAmount($var)
    {
        GPBUtil::checkUint64($var);
        $this->original_amount = $var;

        return $this;
    }

    /**
     * The treasury used for this the private action
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId treasury = 8 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getTreasury()
    {
        return $this->treasury;
    }

    public function hasTreasury()
    {
        return isset($this->treasury);
    }

    public function clearTreasury()
    {
        unset($this->treasury);
    }

    /**
     * The treasury used for this the private action
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId treasury = 8 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setTreasury($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->treasury = $var;

        return $this;
    }

    /**
     * The recent root observed at the time of intent creation for this private action
     *
     * Generated from protobuf field <code>.code.common.v1.Hash recent_root = 9 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\Hash|null
     */
    public function getRecentRoot()
    {
        return $this->recent_root;
    }

    public function hasRecentRoot()
    {
        return isset($this->recent_root);
    }

    public function clearRecentRoot()
    {
        unset($this->recent_root);
    }

    /**
     * The recent root observed at the time of intent creation for this private action
     *
     * Generated from protobuf field <code>.code.common.v1.Hash recent_root = 9 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\Hash $var
     * @return $this
     */
    public function setRecentRoot($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\Hash::class);
        $this->recent_root = $var;

        return $this;
    }

}

// Adding a class alias for backwards compatibility with the previous class name.
class_alias(UpgradeablePrivateAction::class, \Code\Transaction\V2\UpgradeableIntent_UpgradeablePrivateAction::class);

