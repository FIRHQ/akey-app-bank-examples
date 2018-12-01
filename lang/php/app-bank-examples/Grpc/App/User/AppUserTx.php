<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user.proto

namespace App\User;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>app.user.AppUserTx</code>
 */
class AppUserTx extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    private $id = '';
    /**
     * Generated from protobuf field <code>string app_id = 2;</code>
     */
    private $app_id = '';
    /**
     * Generated from protobuf field <code>string from_user_id = 3;</code>
     */
    private $from_user_id = '';
    /**
     * Generated from protobuf field <code>string to_user_id = 4;</code>
     */
    private $to_user_id = '';
    /**
     * Generated from protobuf field <code>string from_user_currency_id = 5;</code>
     */
    private $from_user_currency_id = '';
    /**
     * Generated from protobuf field <code>string amount = 6;</code>
     */
    private $amount = '';
    /**
     * Generated from protobuf field <code>string bank_id = 7;</code>
     */
    private $bank_id = '';
    /**
     * Generated from protobuf field <code>uint64 fee = 8;</code>
     */
    private $fee = 0;
    /**
     * Generated from protobuf field <code>string coin_type = 9;</code>
     */
    private $coin_type = '';
    /**
     * Generated from protobuf field <code>string address = 10;</code>
     */
    private $address = '';
    /**
     * Generated from protobuf field <code>string main_net = 11;</code>
     */
    private $main_net = '';
    /**
     * Generated from protobuf field <code>uint64 trade_at = 12;</code>
     */
    private $trade_at = 0;
    /**
     * Generated from protobuf field <code>string session_id = 13;</code>
     */
    private $session_id = '';
    /**
     * Generated from protobuf field <code>uint32 decimals = 14;</code>
     */
    private $decimals = 0;
    /**
     * Generated from protobuf field <code>uint32 tx_status = 15;</code>
     */
    private $tx_status = 0;
    /**
     * Generated from protobuf field <code>string code = 16;</code>
     */
    private $code = '';
    /**
     * Generated from protobuf field <code>string crypto_type = 17;</code>
     */
    private $crypto_type = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $app_id
     *     @type string $from_user_id
     *     @type string $to_user_id
     *     @type string $from_user_currency_id
     *     @type string $amount
     *     @type string $bank_id
     *     @type int|string $fee
     *     @type string $coin_type
     *     @type string $address
     *     @type string $main_net
     *     @type int|string $trade_at
     *     @type string $session_id
     *     @type int $decimals
     *     @type int $tx_status
     *     @type string $code
     *     @type string $crypto_type
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\User::initOnce();
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
     * Generated from protobuf field <code>string app_id = 2;</code>
     * @return string
     */
    public function getAppId()
    {
        return $this->app_id;
    }

    /**
     * Generated from protobuf field <code>string app_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setAppId($var)
    {
        GPBUtil::checkString($var, True);
        $this->app_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string from_user_id = 3;</code>
     * @return string
     */
    public function getFromUserId()
    {
        return $this->from_user_id;
    }

    /**
     * Generated from protobuf field <code>string from_user_id = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setFromUserId($var)
    {
        GPBUtil::checkString($var, True);
        $this->from_user_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string to_user_id = 4;</code>
     * @return string
     */
    public function getToUserId()
    {
        return $this->to_user_id;
    }

    /**
     * Generated from protobuf field <code>string to_user_id = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setToUserId($var)
    {
        GPBUtil::checkString($var, True);
        $this->to_user_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string from_user_currency_id = 5;</code>
     * @return string
     */
    public function getFromUserCurrencyId()
    {
        return $this->from_user_currency_id;
    }

    /**
     * Generated from protobuf field <code>string from_user_currency_id = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setFromUserCurrencyId($var)
    {
        GPBUtil::checkString($var, True);
        $this->from_user_currency_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string amount = 6;</code>
     * @return string
     */
    public function getAmount()
    {
        return $this->amount;
    }

    /**
     * Generated from protobuf field <code>string amount = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setAmount($var)
    {
        GPBUtil::checkString($var, True);
        $this->amount = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string bank_id = 7;</code>
     * @return string
     */
    public function getBankId()
    {
        return $this->bank_id;
    }

    /**
     * Generated from protobuf field <code>string bank_id = 7;</code>
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
     * Generated from protobuf field <code>uint64 fee = 8;</code>
     * @return int|string
     */
    public function getFee()
    {
        return $this->fee;
    }

    /**
     * Generated from protobuf field <code>uint64 fee = 8;</code>
     * @param int|string $var
     * @return $this
     */
    public function setFee($var)
    {
        GPBUtil::checkUint64($var);
        $this->fee = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string coin_type = 9;</code>
     * @return string
     */
    public function getCoinType()
    {
        return $this->coin_type;
    }

    /**
     * Generated from protobuf field <code>string coin_type = 9;</code>
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
     * Generated from protobuf field <code>string address = 10;</code>
     * @return string
     */
    public function getAddress()
    {
        return $this->address;
    }

    /**
     * Generated from protobuf field <code>string address = 10;</code>
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
     * Generated from protobuf field <code>string main_net = 11;</code>
     * @return string
     */
    public function getMainNet()
    {
        return $this->main_net;
    }

    /**
     * Generated from protobuf field <code>string main_net = 11;</code>
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
     * Generated from protobuf field <code>uint64 trade_at = 12;</code>
     * @return int|string
     */
    public function getTradeAt()
    {
        return $this->trade_at;
    }

    /**
     * Generated from protobuf field <code>uint64 trade_at = 12;</code>
     * @param int|string $var
     * @return $this
     */
    public function setTradeAt($var)
    {
        GPBUtil::checkUint64($var);
        $this->trade_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string session_id = 13;</code>
     * @return string
     */
    public function getSessionId()
    {
        return $this->session_id;
    }

    /**
     * Generated from protobuf field <code>string session_id = 13;</code>
     * @param string $var
     * @return $this
     */
    public function setSessionId($var)
    {
        GPBUtil::checkString($var, True);
        $this->session_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>uint32 decimals = 14;</code>
     * @return int
     */
    public function getDecimals()
    {
        return $this->decimals;
    }

    /**
     * Generated from protobuf field <code>uint32 decimals = 14;</code>
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
     * Generated from protobuf field <code>uint32 tx_status = 15;</code>
     * @return int
     */
    public function getTxStatus()
    {
        return $this->tx_status;
    }

    /**
     * Generated from protobuf field <code>uint32 tx_status = 15;</code>
     * @param int $var
     * @return $this
     */
    public function setTxStatus($var)
    {
        GPBUtil::checkUint32($var);
        $this->tx_status = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string code = 16;</code>
     * @return string
     */
    public function getCode()
    {
        return $this->code;
    }

    /**
     * Generated from protobuf field <code>string code = 16;</code>
     * @param string $var
     * @return $this
     */
    public function setCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string crypto_type = 17;</code>
     * @return string
     */
    public function getCryptoType()
    {
        return $this->crypto_type;
    }

    /**
     * Generated from protobuf field <code>string crypto_type = 17;</code>
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

