<?php
// GENERATED CODE -- DO NOT EDIT!

namespace App\User;

/**
 */
class AppUserServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * sign in user
     * @param \App\User\AppUser $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function signInAppUser(\App\User\AppUser $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/app.user.AppUserService/signInAppUser',
        $argument,
        ['\App\User\AppUser', 'decode'],
        $metadata, $options);
    }

    /**
     * create tx
     * @param \App\User\AppUserTx $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function createTx(\App\User\AppUserTx $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/app.user.AppUserService/createTx',
        $argument,
        ['\App\User\AppUserTx', 'decode'],
        $metadata, $options);
    }

    /**
     * send coin
     * @param \App\User\AppUserTx $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function sendTx(\App\User\AppUserTx $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/app.user.AppUserService/sendTx',
        $argument,
        ['\App\User\AppUserTx', 'decode'],
        $metadata, $options);
    }

    /**
     * get balances
     * @param \App\User\AppUser $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function getBalance(\App\User\AppUser $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/app.user.AppUserService/getBalance',
        $argument,
        ['\App\User\AppUserBalance', 'decode'],
        $metadata, $options);
    }

    /**
     * withdraw
     * @param \App\User\AppUserWithdraw $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function withdraw(\App\User\AppUserWithdraw $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/app.user.AppUserService/withdraw',
        $argument,
        ['\App\User\AppUserWithdraw', 'decode'],
        $metadata, $options);
    }

    /**
     * create user currency address
     * @param \App\User\UserAddress $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function createUserAddress(\App\User\UserAddress $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/app.user.AppUserService/createUserAddress',
        $argument,
        ['\App\User\UserAddress', 'decode'],
        $metadata, $options);
    }

}
