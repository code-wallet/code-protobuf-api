<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chat/v2/chat_service.proto

namespace Code\Chat\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Identity revealed content that is inserted into chat whenever a chat member
 * reveals their identity
 *
 * Generated from protobuf message <code>code.chat.v2.IdentityRevealedContent</code>
 */
class IdentityRevealedContent extends \Google\Protobuf\Internal\Message
{
    /**
     * The chat member who revealed their identity
     *
     * Generated from protobuf field <code>.code.chat.v2.ChatMemberId member_id = 1 [(.validate.rules) = {</code>
     */
    protected $member_id = null;
    /**
     * The identity that was revealed
     *
     * Generated from protobuf field <code>.code.chat.v2.ChatMemberIdentity identity = 2 [(.validate.rules) = {</code>
     */
    protected $identity = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Chat\V2\ChatMemberId $member_id
     *           The chat member who revealed their identity
     *     @type \Code\Chat\V2\ChatMemberIdentity $identity
     *           The identity that was revealed
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Chat\V2\ChatService::initOnce();
        parent::__construct($data);
    }

    /**
     * The chat member who revealed their identity
     *
     * Generated from protobuf field <code>.code.chat.v2.ChatMemberId member_id = 1 [(.validate.rules) = {</code>
     * @return \Code\Chat\V2\ChatMemberId|null
     */
    public function getMemberId()
    {
        return $this->member_id;
    }

    public function hasMemberId()
    {
        return isset($this->member_id);
    }

    public function clearMemberId()
    {
        unset($this->member_id);
    }

    /**
     * The chat member who revealed their identity
     *
     * Generated from protobuf field <code>.code.chat.v2.ChatMemberId member_id = 1 [(.validate.rules) = {</code>
     * @param \Code\Chat\V2\ChatMemberId $var
     * @return $this
     */
    public function setMemberId($var)
    {
        GPBUtil::checkMessage($var, \Code\Chat\V2\ChatMemberId::class);
        $this->member_id = $var;

        return $this;
    }

    /**
     * The identity that was revealed
     *
     * Generated from protobuf field <code>.code.chat.v2.ChatMemberIdentity identity = 2 [(.validate.rules) = {</code>
     * @return \Code\Chat\V2\ChatMemberIdentity|null
     */
    public function getIdentity()
    {
        return $this->identity;
    }

    public function hasIdentity()
    {
        return isset($this->identity);
    }

    public function clearIdentity()
    {
        unset($this->identity);
    }

    /**
     * The identity that was revealed
     *
     * Generated from protobuf field <code>.code.chat.v2.ChatMemberIdentity identity = 2 [(.validate.rules) = {</code>
     * @param \Code\Chat\V2\ChatMemberIdentity $var
     * @return $this
     */
    public function setIdentity($var)
    {
        GPBUtil::checkMessage($var, \Code\Chat\V2\ChatMemberIdentity::class);
        $this->identity = $var;

        return $this;
    }

}

