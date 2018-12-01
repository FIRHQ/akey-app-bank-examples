<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Developer;

/**
 */
class DeveloperServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * developer init session
     * @param \Developer\SessionRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function initSession(\Developer\SessionRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/developer.DeveloperService/initSession',
        $argument,
        ['\Developer\Session', 'decode'],
        $metadata, $options);
    }

    /**
     * get balance
     * @param \Developer\Session $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function getBalance(\Developer\Session $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/developer.DeveloperService/getBalance',
        $argument,
        ['\Developer\AppBankBalance', 'decode'],
        $metadata, $options);
    }

    /**
     * send coin from bank
     * @param \Developer\BankTx $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function sendTxFromBank(\Developer\BankTx $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/developer.DeveloperService/sendTxFromBank',
        $argument,
        ['\Developer\BankTx', 'decode'],
        $metadata, $options);
    }

}
