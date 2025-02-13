<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/v1/model.proto

namespace Code\Common\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * InstructionAccount is an account public key used within the context of
 * an instruction.
 *
 * Generated from protobuf message <code>code.common.v1.InstructionAccount</code>
 */
class InstructionAccount extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId account = 1 [(.validate.rules) = {</code>
     */
    protected $account = null;
    /**
     * Generated from protobuf field <code>bool is_signer = 2;</code>
     */
    protected $is_signer = false;
    /**
     * Generated from protobuf field <code>bool is_writable = 3;</code>
     */
    protected $is_writable = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\SolanaAccountId $account
     *     @type bool $is_signer
     *     @type bool $is_writable
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Common\V1\Model::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId account = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\SolanaAccountId|null
     */
    public function getAccount()
    {
        return $this->account;
    }

    public function hasAccount()
    {
        return isset($this->account);
    }

    public function clearAccount()
    {
        unset($this->account);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.SolanaAccountId account = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\SolanaAccountId $var
     * @return $this
     */
    public function setAccount($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\SolanaAccountId::class);
        $this->account = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool is_signer = 2;</code>
     * @return bool
     */
    public function getIsSigner()
    {
        return $this->is_signer;
    }

    /**
     * Generated from protobuf field <code>bool is_signer = 2;</code>
     * @param bool $var
     * @return $this
     */
    public function setIsSigner($var)
    {
        GPBUtil::checkBool($var);
        $this->is_signer = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool is_writable = 3;</code>
     * @return bool
     */
    public function getIsWritable()
    {
        return $this->is_writable;
    }

    /**
     * Generated from protobuf field <code>bool is_writable = 3;</code>
     * @param bool $var
     * @return $this
     */
    public function setIsWritable($var)
    {
        GPBUtil::checkBool($var);
        $this->is_writable = $var;

        return $this;
    }

}

