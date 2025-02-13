<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: invite/v2/invite_service.proto

namespace Code\Invite\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.invite.v2.GetInvitationStatusResponse</code>
 */
class GetInvitationStatusResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.code.invite.v2.GetInvitationStatusResponse.Result result = 1;</code>
     */
    protected $result = 0;
    /**
     * The user's invitation status
     *
     * Generated from protobuf field <code>.code.invite.v2.InvitationStatus status = 2;</code>
     */
    protected $status = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $result
     *     @type int $status
     *           The user's invitation status
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Invite\V2\InviteService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.invite.v2.GetInvitationStatusResponse.Result result = 1;</code>
     * @return int
     */
    public function getResult()
    {
        return $this->result;
    }

    /**
     * Generated from protobuf field <code>.code.invite.v2.GetInvitationStatusResponse.Result result = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setResult($var)
    {
        GPBUtil::checkEnum($var, \Code\Invite\V2\GetInvitationStatusResponse\Result::class);
        $this->result = $var;

        return $this;
    }

    /**
     * The user's invitation status
     *
     * Generated from protobuf field <code>.code.invite.v2.InvitationStatus status = 2;</code>
     * @return int
     */
    public function getStatus()
    {
        return $this->status;
    }

    /**
     * The user's invitation status
     *
     * Generated from protobuf field <code>.code.invite.v2.InvitationStatus status = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setStatus($var)
    {
        GPBUtil::checkEnum($var, \Code\Invite\V2\InvitationStatus::class);
        $this->status = $var;

        return $this;
    }

}

