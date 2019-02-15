<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Meta;

/**
 */
class AgentClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Meta\MetaRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetInfo(\Meta\MetaRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/meta.Agent/GetInfo',
        $argument,
        ['\Meta\MetaResponse', 'decode'],
        $metadata, $options);
    }

}
