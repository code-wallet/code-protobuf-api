<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: messaging/v1/messaging_service.proto

namespace Code\Messaging\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.messaging.v1.SendMessageRequest</code>
 */
class SendMessageRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * The message to send. Types of messages clients can send are restricted.
     *
     * Generated from protobuf field <code>.code.messaging.v1.Message message = 1 [(.validate.rules) = {</code>
     */
    protected $message = null;
    /**
     * The rendezvous key that the message should be routed to.
     *
     * Generated from protobuf field <code>.code.messaging.v1.RendezvousKey rendezvous_key = 2 [(.validate.rules) = {</code>
     */
    protected $rendezvous_key = null;
    /**
     * The signature is of serialize(Message) using the PrivateKey of the keypair.
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
     *     @type \Code\Messaging\V1\Message $message
     *           The message to send. Types of messages clients can send are restricted.
     *     @type \Code\Messaging\V1\RendezvousKey $rendezvous_key
     *           The rendezvous key that the message should be routed to.
     *     @type \Code\Common\V1\Signature $signature
     *           The signature is of serialize(Message) using the PrivateKey of the keypair.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Messaging\V1\MessagingService::initOnce();
        parent::__construct($data);
    }

    /**
     * The message to send. Types of messages clients can send are restricted.
     *
     * Generated from protobuf field <code>.code.messaging.v1.Message message = 1 [(.validate.rules) = {</code>
     * @return \Code\Messaging\V1\Message|null
     */
    public function getMessage()
    {
        return $this->message;
    }

    public function hasMessage()
    {
        return isset($this->message);
    }

    public function clearMessage()
    {
        unset($this->message);
    }

    /**
     * The message to send. Types of messages clients can send are restricted.
     *
     * Generated from protobuf field <code>.code.messaging.v1.Message message = 1 [(.validate.rules) = {</code>
     * @param \Code\Messaging\V1\Message $var
     * @return $this
     */
    public function setMessage($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\Message::class);
        $this->message = $var;

        return $this;
    }

    /**
     * The rendezvous key that the message should be routed to.
     *
     * Generated from protobuf field <code>.code.messaging.v1.RendezvousKey rendezvous_key = 2 [(.validate.rules) = {</code>
     * @return \Code\Messaging\V1\RendezvousKey|null
     */
    public function getRendezvousKey()
    {
        return $this->rendezvous_key;
    }

    public function hasRendezvousKey()
    {
        return isset($this->rendezvous_key);
    }

    public function clearRendezvousKey()
    {
        unset($this->rendezvous_key);
    }

    /**
     * The rendezvous key that the message should be routed to.
     *
     * Generated from protobuf field <code>.code.messaging.v1.RendezvousKey rendezvous_key = 2 [(.validate.rules) = {</code>
     * @param \Code\Messaging\V1\RendezvousKey $var
     * @return $this
     */
    public function setRendezvousKey($var)
    {
        GPBUtil::checkMessage($var, \Code\Messaging\V1\RendezvousKey::class);
        $this->rendezvous_key = $var;

        return $this;
    }

    /**
     * The signature is of serialize(Message) using the PrivateKey of the keypair.
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
     * The signature is of serialize(Message) using the PrivateKey of the keypair.
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

