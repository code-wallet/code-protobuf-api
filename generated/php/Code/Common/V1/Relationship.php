<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/v1/model.proto

namespace Code\Common\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Relationship is a set of identifiers that a user can establish a relationship
 * with.
 *
 * Generated from protobuf message <code>code.common.v1.Relationship</code>
 */
class Relationship extends \Google\Protobuf\Internal\Message
{
    protected $type;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Code\Common\V1\Domain $domain
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Common\V1\Model::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.Domain domain = 1 [(.validate.rules) = {</code>
     * @return \Code\Common\V1\Domain|null
     */
    public function getDomain()
    {
        return $this->readOneof(1);
    }

    public function hasDomain()
    {
        return $this->hasOneof(1);
    }

    /**
     * Generated from protobuf field <code>.code.common.v1.Domain domain = 1 [(.validate.rules) = {</code>
     * @param \Code\Common\V1\Domain $var
     * @return $this
     */
    public function setDomain($var)
    {
        GPBUtil::checkMessage($var, \Code\Common\V1\Domain::class);
        $this->writeOneof(1, $var);

        return $this;
    }

    /**
     * @return string
     */
    public function getType()
    {
        return $this->whichOneof("type");
    }

}

