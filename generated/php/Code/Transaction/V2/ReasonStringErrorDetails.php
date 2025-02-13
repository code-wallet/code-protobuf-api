<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: transaction/v2/transaction_service.proto

namespace Code\Transaction\V2;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.transaction.v2.ReasonStringErrorDetails</code>
 */
class ReasonStringErrorDetails extends \Google\Protobuf\Internal\Message
{
    /**
     * Human readable string indicating the failure.
     *
     * Generated from protobuf field <code>string reason = 1 [(.validate.rules) = {</code>
     */
    protected $reason = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $reason
     *           Human readable string indicating the failure.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Transaction\V2\TransactionService::initOnce();
        parent::__construct($data);
    }

    /**
     * Human readable string indicating the failure.
     *
     * Generated from protobuf field <code>string reason = 1 [(.validate.rules) = {</code>
     * @return string
     */
    public function getReason()
    {
        return $this->reason;
    }

    /**
     * Human readable string indicating the failure.
     *
     * Generated from protobuf field <code>string reason = 1 [(.validate.rules) = {</code>
     * @param string $var
     * @return $this
     */
    public function setReason($var)
    {
        GPBUtil::checkString($var, True);
        $this->reason = $var;

        return $this;
    }

}

