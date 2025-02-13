<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: invite/v2/invite_service.proto

namespace Code\Invite\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.invite.v2.InviteCode</code>
 */
class InviteCode extends \Google\Protobuf\Internal\Message
{
    /**
     * Regex for invite codes
     *
     * Generated from protobuf field <code>string value = 1 [(.validate.rules) = {</code>
     */
    protected $value = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $value
     *           Regex for invite codes
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Invite\V2\InviteService::initOnce();
        parent::__construct($data);
    }

    /**
     * Regex for invite codes
     *
     * Generated from protobuf field <code>string value = 1 [(.validate.rules) = {</code>
     * @return string
     */
    public function getValue()
    {
        return $this->value;
    }

    /**
     * Regex for invite codes
     *
     * Generated from protobuf field <code>string value = 1 [(.validate.rules) = {</code>
     * @param string $var
     * @return $this
     */
    public function setValue($var)
    {
        GPBUtil::checkString($var, True);
        $this->value = $var;

        return $this;
    }

}

