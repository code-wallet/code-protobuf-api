<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: device/v1/device_service.proto

namespace Code\Device\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.device.v1.GetLoggedInAccountsResponse</code>
 */
class GetLoggedInAccountsResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.code.device.v1.GetLoggedInAccountsResponse.Result result = 1;</code>
     */
    protected $result = 0;
    /**
     * Generated from protobuf field <code>repeated .code.common.v1.SolanaAccountId owners = 2 [(.validate.rules) = {</code>
     */
    private $owners;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $result
     *     @type array<\Code\Common\V1\SolanaAccountId>|\Google\Protobuf\Internal\RepeatedField $owners
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Device\V1\DeviceService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.device.v1.GetLoggedInAccountsResponse.Result result = 1;</code>
     * @return int
     */
    public function getResult()
    {
        return $this->result;
    }

    /**
     * Generated from protobuf field <code>.code.device.v1.GetLoggedInAccountsResponse.Result result = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setResult($var)
    {
        GPBUtil::checkEnum($var, \Code\Device\V1\GetLoggedInAccountsResponse\Result::class);
        $this->result = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .code.common.v1.SolanaAccountId owners = 2 [(.validate.rules) = {</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getOwners()
    {
        return $this->owners;
    }

    /**
     * Generated from protobuf field <code>repeated .code.common.v1.SolanaAccountId owners = 2 [(.validate.rules) = {</code>
     * @param array<\Code\Common\V1\SolanaAccountId>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setOwners($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Code\Common\V1\SolanaAccountId::class);
        $this->owners = $arr;

        return $this;
    }

}

