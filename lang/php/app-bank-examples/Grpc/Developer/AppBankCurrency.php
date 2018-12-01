<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: developer.proto

namespace Developer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>developer.AppBankCurrency</code>
 */
class AppBankCurrency extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    private $id = '';
    /**
     * Generated from protobuf field <code>string bank_id = 2;</code>
     */
    private $bank_id = '';
    /**
     * Generated from protobuf field <code>string coin_type = 3;</code>
     */
    private $coin_type = '';
    /**
     * Generated from protobuf field <code>string balance = 4;</code>
     */
    private $balance = '';
    /**
     * Generated from protobuf field <code>string main_net = 5;</code>
     */
    private $main_net = '';
    /**
     * Generated from protobuf field <code>double gas_rate = 6;</code>
     */
    private $gas_rate = 0.0;
    /**
     * Generated from protobuf field <code>uint32 gas = 7;</code>
     */
    private $gas = 0;
    /**
     * Generated from protobuf field <code>string address_balance = 8;</code>
     */
    private $address_balance = '';
    /**
     * Generated from protobuf field <code>string address = 9;</code>
     */
    private $address = '';
    /**
     * Generated from protobuf field <code>uint32 decimals = 10;</code>
     */
    private $decimals = 0;
    /**
     * Generated from protobuf field <code>string crypto_type = 11;</code>
     */
    private $crypto_type = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $bank_id
     *     @type string $coin_type
     *     @type string $balance
     *     @type string $main_net
     *     @type float $gas_rate
     *     @type int $gas
     *     @type string $address_balance
     *     @type string $address
     *     @type int $decimals
     *     @type string $crypto_type
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Developer::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string id = 1;</code>
     * @return string
     */
    public function getId()
    {
        return $this->id;
    }

    /**
     * Generated from protobuf field <code>string id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setId($var)
    {
        GPBUtil::checkString($var, True);
        $this->id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string bank_id = 2;</code>
     * @return string
     */
    public function getBankId()
    {
        return $this->bank_id;
    }

    /**
     * Generated from protobuf field <code>string bank_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setBankId($var)
    {
        GPBUtil::checkString($var, True);
        $this->bank_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string coin_type = 3;</code>
     * @return string
     */
    public function getCoinType()
    {
        return $this->coin_type;
    }

    /**
     * Generated from protobuf field <code>string coin_type = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setCoinType($var)
    {
        GPBUtil::checkString($var, True);
        $this->coin_type = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string balance = 4;</code>
     * @return string
     */
    public function getBalance()
    {
        return $this->balance;
    }

    /**
     * Generated from protobuf field <code>string balance = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setBalance($var)
    {
        GPBUtil::checkString($var, True);
        $this->balance = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string main_net = 5;</code>
     * @return string
     */
    public function getMainNet()
    {
        return $this->main_net;
    }

    /**
     * Generated from protobuf field <code>string main_net = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setMainNet($var)
    {
        GPBUtil::checkString($var, True);
        $this->main_net = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>double gas_rate = 6;</code>
     * @return float
     */
    public function getGasRate()
    {
        return $this->gas_rate;
    }

    /**
     * Generated from protobuf field <code>double gas_rate = 6;</code>
     * @param float $var
     * @return $this
     */
    public function setGasRate($var)
    {
        GPBUtil::checkDouble($var);
        $this->gas_rate = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>uint32 gas = 7;</code>
     * @return int
     */
    public function getGas()
    {
        return $this->gas;
    }

    /**
     * Generated from protobuf field <code>uint32 gas = 7;</code>
     * @param int $var
     * @return $this
     */
    public function setGas($var)
    {
        GPBUtil::checkUint32($var);
        $this->gas = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string address_balance = 8;</code>
     * @return string
     */
    public function getAddressBalance()
    {
        return $this->address_balance;
    }

    /**
     * Generated from protobuf field <code>string address_balance = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setAddressBalance($var)
    {
        GPBUtil::checkString($var, True);
        $this->address_balance = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string address = 9;</code>
     * @return string
     */
    public function getAddress()
    {
        return $this->address;
    }

    /**
     * Generated from protobuf field <code>string address = 9;</code>
     * @param string $var
     * @return $this
     */
    public function setAddress($var)
    {
        GPBUtil::checkString($var, True);
        $this->address = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>uint32 decimals = 10;</code>
     * @return int
     */
    public function getDecimals()
    {
        return $this->decimals;
    }

    /**
     * Generated from protobuf field <code>uint32 decimals = 10;</code>
     * @param int $var
     * @return $this
     */
    public function setDecimals($var)
    {
        GPBUtil::checkUint32($var);
        $this->decimals = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string crypto_type = 11;</code>
     * @return string
     */
    public function getCryptoType()
    {
        return $this->crypto_type;
    }

    /**
     * Generated from protobuf field <code>string crypto_type = 11;</code>
     * @param string $var
     * @return $this
     */
    public function setCryptoType($var)
    {
        GPBUtil::checkString($var, True);
        $this->crypto_type = $var;

        return $this;
    }

}

