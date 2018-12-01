<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: developer.proto

namespace Developer;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>developer.Developer</code>
 */
class Developer extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string id = 1;</code>
     */
    private $id = '';
    /**
     * Generated from protobuf field <code>string user_id = 2;</code>
     */
    private $user_id = '';
    /**
     * Generated from protobuf field <code>string mobile = 3;</code>
     */
    private $mobile = '';
    /**
     * Generated from protobuf field <code>string email = 4;</code>
     */
    private $email = '';
    /**
     * Generated from protobuf field <code>string id_card_number = 5;</code>
     */
    private $id_card_number = '';
    /**
     * Generated from protobuf field <code>string real_name = 6;</code>
     */
    private $real_name = '';
    /**
     * Generated from protobuf field <code>string private_key = 7;</code>
     */
    private $private_key = '';
    /**
     * Generated from protobuf field <code>string public_key = 8;</code>
     */
    private $public_key = '';
    /**
     * Generated from protobuf field <code>string access_key = 9;</code>
     */
    private $access_key = '';
    /**
     * Generated from protobuf field <code>uint64 create_at = 10;</code>
     */
    private $create_at = 0;
    /**
     * Generated from protobuf field <code>uint64 update_at = 11;</code>
     */
    private $update_at = 0;
    /**
     * Generated from protobuf field <code>bool active = 12;</code>
     */
    private $active = false;
    /**
     * Generated from protobuf field <code>bool verified = 13;</code>
     */
    private $verified = false;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $id
     *     @type string $user_id
     *     @type string $mobile
     *     @type string $email
     *     @type string $id_card_number
     *     @type string $real_name
     *     @type string $private_key
     *     @type string $public_key
     *     @type string $access_key
     *     @type int|string $create_at
     *     @type int|string $update_at
     *     @type bool $active
     *     @type bool $verified
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
     * Generated from protobuf field <code>string user_id = 2;</code>
     * @return string
     */
    public function getUserId()
    {
        return $this->user_id;
    }

    /**
     * Generated from protobuf field <code>string user_id = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setUserId($var)
    {
        GPBUtil::checkString($var, True);
        $this->user_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string mobile = 3;</code>
     * @return string
     */
    public function getMobile()
    {
        return $this->mobile;
    }

    /**
     * Generated from protobuf field <code>string mobile = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setMobile($var)
    {
        GPBUtil::checkString($var, True);
        $this->mobile = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string email = 4;</code>
     * @return string
     */
    public function getEmail()
    {
        return $this->email;
    }

    /**
     * Generated from protobuf field <code>string email = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setEmail($var)
    {
        GPBUtil::checkString($var, True);
        $this->email = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string id_card_number = 5;</code>
     * @return string
     */
    public function getIdCardNumber()
    {
        return $this->id_card_number;
    }

    /**
     * Generated from protobuf field <code>string id_card_number = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setIdCardNumber($var)
    {
        GPBUtil::checkString($var, True);
        $this->id_card_number = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string real_name = 6;</code>
     * @return string
     */
    public function getRealName()
    {
        return $this->real_name;
    }

    /**
     * Generated from protobuf field <code>string real_name = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setRealName($var)
    {
        GPBUtil::checkString($var, True);
        $this->real_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string private_key = 7;</code>
     * @return string
     */
    public function getPrivateKey()
    {
        return $this->private_key;
    }

    /**
     * Generated from protobuf field <code>string private_key = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setPrivateKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->private_key = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string public_key = 8;</code>
     * @return string
     */
    public function getPublicKey()
    {
        return $this->public_key;
    }

    /**
     * Generated from protobuf field <code>string public_key = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setPublicKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->public_key = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string access_key = 9;</code>
     * @return string
     */
    public function getAccessKey()
    {
        return $this->access_key;
    }

    /**
     * Generated from protobuf field <code>string access_key = 9;</code>
     * @param string $var
     * @return $this
     */
    public function setAccessKey($var)
    {
        GPBUtil::checkString($var, True);
        $this->access_key = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>uint64 create_at = 10;</code>
     * @return int|string
     */
    public function getCreateAt()
    {
        return $this->create_at;
    }

    /**
     * Generated from protobuf field <code>uint64 create_at = 10;</code>
     * @param int|string $var
     * @return $this
     */
    public function setCreateAt($var)
    {
        GPBUtil::checkUint64($var);
        $this->create_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>uint64 update_at = 11;</code>
     * @return int|string
     */
    public function getUpdateAt()
    {
        return $this->update_at;
    }

    /**
     * Generated from protobuf field <code>uint64 update_at = 11;</code>
     * @param int|string $var
     * @return $this
     */
    public function setUpdateAt($var)
    {
        GPBUtil::checkUint64($var);
        $this->update_at = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool active = 12;</code>
     * @return bool
     */
    public function getActive()
    {
        return $this->active;
    }

    /**
     * Generated from protobuf field <code>bool active = 12;</code>
     * @param bool $var
     * @return $this
     */
    public function setActive($var)
    {
        GPBUtil::checkBool($var);
        $this->active = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool verified = 13;</code>
     * @return bool
     */
    public function getVerified()
    {
        return $this->verified;
    }

    /**
     * Generated from protobuf field <code>bool verified = 13;</code>
     * @param bool $var
     * @return $this
     */
    public function setVerified($var)
    {
        GPBUtil::checkBool($var);
        $this->verified = $var;

        return $this;
    }

}

