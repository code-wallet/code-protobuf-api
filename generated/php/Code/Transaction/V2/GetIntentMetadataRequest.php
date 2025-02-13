<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.transaction.v2.GetIntentMetadataRequest</code>
 */
class GetIntentMetadataRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * The intent ID to query
     *
     * Generated from protobuf field <code>.code.common.v1.IntentId intent_id = 1 [(.validate.rules) = {</code>
     */
    protected $intent_id = null;
    /**
     * The verified owner account public key when not signing with the rendezvous
     * key. Only owner accounts involved in the intent can access the metadata.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 2;</code>
     */
    protected $owner = null;
    /**
     * The signature is of serialize(GetIntentStatusRequest) without this field set
     * using the private key of the rendezvous or owner account. This provides an
     * authentication mechanism to the RPC.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 3 [(.validate.rules) = {</code>
     */
    protected $signature = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\IntentId $intent_id
     *           The intent ID to query
     *     @type \Code\Common\V1\SolanaAccountId $owner
     *           The verified owner account public key when not signing with the rendezvous
     *           key. Only owner accounts involved in the intent can access the metadata.
     *     @type \Code\Common\V1\Signature $signature
     *           The signature is of serialize(GetIntentStatusRequest) without this field set
     *           using the private key of the rendezvous or owner account. This provides an
     *           authentication mechanism to the RPC.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * The intent ID to query
     *
     * Generated from protobuf field <code>.code.common.v1.IntentId intent_id = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\IntentId|null
     */
    public function getIntentId()
    {
        return $this->intent_id;
    }

    public function hasIntentId()
    {
        return isset($this->intent_id);
    }

    public function clearIntentId()
    {
        unset($this->intent_id);
    }

    /**
     * The intent ID to query
     *
     * Generated from protobuf field <code>.code.common.v1.IntentId intent_id = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\IntentId $var
     * @return $this
     */
    public function setIntentId($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\IntentId::class);
        $this->intent_id = $var;

        return $this;
    }

    /**
     * The verified owner account public key when not signing with the rendezvous
     * key. Only owner accounts involved in the intent can access the metadata.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 2;</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getOwner()
    {
        return $this->owner;
    }

    public function hasOwner()
    {
        return isset($this->owner);
    }

    public function clearOwner()
    {
        unset($this->owner);
    }

    /**
     * The verified owner account public key when not signing with the rendezvous
     * key. Only owner accounts involved in the intent can access the metadata.
     *
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId owner = 2;</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setOwner($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->owner = $var;

        return $this;
    }

    /**
     * The signature is of serialize(GetIntentStatusRequest) without this field set
     * using the private key of the rendezvous or owner account. This provides an
     * authentication mechanism to the RPC.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 3 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\Signature|null
     */
    public function getSignature()
    {
        return $this->signature;
    }

    public function hasSignature()
    {
        return isset($this->signature);
    }

    public function clearSignature()
    {
        unset($this->signature);
    }

    /**
     * The signature is of serialize(GetIntentStatusRequest) without this field set
     * using the private key of the rendezvous or owner account. This provides an
     * authentication mechanism to the RPC.
     *
     * Generated from protobuf field <code>.code.common.v1.Signature signature = 3 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\Signature $var
     * @return $this
     */
    public function setSignature($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\Signature::class);
        $this->signature = $var;

        return $this;
    }

}

