<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: messaging/v1/messaging_service.proto

namespace Code\Messaging\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.messaging.v1.PollMessagesResponse</code>
 */
class PollMessagesResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .code.messaging.v1.Message messages = 1 [(.validate.rules) = {</code>
     */
    private $messages;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type array<\Code\Messaging\V1\Message>|\Google\Protobuf\Internal\RepeatedField $messages
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Messaging\V1\MessagingService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .code.messaging.v1.Message messages = 1 [(.validate.rules) = {</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getMessages()
    {
        return $this->messages;
    }

    /**
     * Generated from protobuf field <code>repeated .code.messaging.v1.Message messages = 1 [(.validate.rules) = {</code>
     * @param array<\Code\Messaging\V1\Message>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setMessages($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Code\Messaging\V1\Message::class);
        $this->messages = $arr;

        return $this;
    }

}

