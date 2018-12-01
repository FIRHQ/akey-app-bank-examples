<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Tsv;

/**
 */
class TSVServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * service version
     * @param \Google\Protobuf\GPBEmpty $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function versionInfo(\Google\Protobuf\GPBEmpty $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/tsv.TSVService/versionInfo',
        $argument,
        ['\VersionInfo', 'decode'],
        $metadata, $options);
    }

    /**
     *
     * get provider from provider name
     * 
     * @param \Tsv\QRRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function getProvider(\Tsv\QRRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/tsv.TSVService/getProvider',
        $argument,
        ['\Tsv\Provider', 'decode'],
        $metadata, $options);
    }

    /**
     *
     * when provider's key store is akey open platform, provider can generate and bind key , user and device.
     *
     * @param \Tsv\KeyRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function generateKey(\Tsv\KeyRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/tsv.TSVService/generateKey',
        $argument,
        ['\Tsv\KeyResponse', 'decode'],
        $metadata, $options);
    }

    /**
     *
     * bind key
     * @param \Tsv\BindKeyRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function bindKey(\Tsv\BindKeyRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/tsv.TSVService/bindKey',
        $argument,
        ['\Tsv\TwoStepVerifyStauts', 'decode'],
        $metadata, $options);
    }

    /**
     *
     * verify code, when code resource is akey open platform
     * @param \Tsv\VerifyCodeRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function verifyCode(\Tsv\VerifyCodeRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/tsv.TSVService/verifyCode',
        $argument,
        ['\Tsv\TwoStepVerifyStauts', 'decode'],
        $metadata, $options);
    }

}
