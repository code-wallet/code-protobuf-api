<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/v1/model.proto

namespace Code\Common\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.common.v1.ClientPong</code>
 */
class ClientPong extends \Google\Protobuf\Internal\Message
{
    /**
     * Timestamp the Pong was sent on the stream, for server to get a sense
     * of potential network latency
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp timestamp = 1 [(.validate.rules) = {</code>
     */
    protected $timestamp = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Google\Protobuf\Timestamp $timestamp
     *           Timestamp the Pong was sent on the stream, for server to get a sense
     *           of potential network latency
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Common\V1\Model::initOnce();
        parent::__construct($data);
    }

    /**
     * Timestamp the Pong was sent on the stream, for server to get a sense
     * of potential network latency
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp timestamp = 1 [(.validate.rules) = {</code>
     * @return \Google\Protobuf\Timestamp|null
     */
    public function getTimestamp()
    {
        return $this->timestamp;
    }

    public function hasTimestamp()
    {
        return isset($this->timestamp);
    }

    public function clearTimestamp()
    {
        unset($this->timestamp);
    }

    /**
     * Timestamp the Pong was sent on the stream, for server to get a sense
     * of potential network latency
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp timestamp = 1 [(.validate.rules) = {</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setTimestamp($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->timestamp = $var;

        return $this;
    }

}

