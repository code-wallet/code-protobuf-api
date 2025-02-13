<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: currency/v1/currency_service.proto

namespace Code\Currency\V1;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>code.currency.v1.GetAllRatesResponse</code>
 */
class GetAllRatesResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.code.currency.v1.GetAllRatesResponse.Result result = 1;</code>
     */
    protected $result = 0;
    /**
     * The time the exchange rates were observed
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp as_of = 2 [(.validate.rules) = {</code>
     */
    protected $as_of = null;
    /**
     * The price of 1 Kin in different currencies, keyed on 3- or 4- letter lowercase currency code.
     *
     * Generated from protobuf field <code>map<string, double> rates = 3 [(.validate.rules) = {</code>
     */
    private $rates;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $result
     *     @type \Google\Protobuf\Timestamp $as_of
     *           The time the exchange rates were observed
     *     @type array|\Google\Protobuf\Internal\MapField $rates
     *           The price of 1 Kin in different currencies, keyed on 3- or 4- letter lowercase currency code.
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Currency\V1\CurrencyService::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.code.currency.v1.GetAllRatesResponse.Result result = 1;</code>
     * @return int
     */
    public function getResult()
    {
        return $this->result;
    }

    /**
     * Generated from protobuf field <code>.code.currency.v1.GetAllRatesResponse.Result result = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setResult($var)
    {
        GPBUtil::checkEnum($var, \Code\Currency\V1\GetAllRatesResponse\Result::class);
        $this->result = $var;

        return $this;
    }

    /**
     * The time the exchange rates were observed
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp as_of = 2 [(.validate.rules) = {</code>
     * @return \Google\Protobuf\Timestamp|null
     */
    public function getAsOf()
    {
        return $this->as_of;
    }

    public function hasAsOf()
    {
        return isset($this->as_of);
    }

    public function clearAsOf()
    {
        unset($this->as_of);
    }

    /**
     * The time the exchange rates were observed
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp as_of = 2 [(.validate.rules) = {</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setAsOf($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->as_of = $var;

        return $this;
    }

    /**
     * The price of 1 Kin in different currencies, keyed on 3- or 4- letter lowercase currency code.
     *
     * Generated from protobuf field <code>map<string, double> rates = 3 [(.validate.rules) = {</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getRates()
    {
        return $this->rates;
    }

    /**
     * The price of 1 Kin in different currencies, keyed on 3- or 4- letter lowercase currency code.
     *
     * Generated from protobuf field <code>map<string, double> rates = 3 [(.validate.rules) = {</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setRates($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::DOUBLE);
        $this->rates = $arr;

        return $this;
    }

}

