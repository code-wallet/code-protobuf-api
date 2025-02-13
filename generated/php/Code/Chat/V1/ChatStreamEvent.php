<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chat/v1/chat_service.proto

namespace Code\Chat\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.chat.v1.ChatStreamEvent</code>
 */
class ChatStreamEvent extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .code.chat.v1.ChatMessage messages = 1 [(.validate.rules) = {</code>
     */
    private $messages;
    /**
     * Generated from protobuf field <code>repeated .code.chat.v1.Pointer pointers = 2 [(.validate.rules) = {</code>
     */
    private $pointers;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type array<\Code\Chat\V1\ChatMessage>|\Google\Protobuf\Internal\RepeatedField $messages
     *     @type array<\Code\Chat\V1\Pointer>|\Google\Protobuf\Internal\RepeatedField $pointers
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Chat\V1\ChatService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .code.chat.v1.ChatMessage messages = 1 [(.validate.rules) = {</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getMessages()
    {
        return $this->messages;
    }

    /**
     * Generated from protobuf field <code>repeated .code.chat.v1.ChatMessage messages = 1 [(.validate.rules) = {</code>
     * @param array<\Code\Chat\V1\ChatMessage>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setMessages($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Code\Chat\V1\ChatMessage::class);
        $this->messages = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>repeated .code.chat.v1.Pointer pointers = 2 [(.validate.rules) = {</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getPointers()
    {
        return $this->pointers;
    }

    /**
     * Generated from protobuf field <code>repeated .code.chat.v1.Pointer pointers = 2 [(.validate.rules) = {</code>
     * @param array<\Code\Chat\V1\Pointer>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setPointers($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Code\Chat\V1\Pointer::class);
        $this->pointers = $arr;

        return $this;
    }

}

