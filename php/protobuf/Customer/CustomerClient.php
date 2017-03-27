<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Customer {

  // The Customer service definition.
  class CustomerClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
      parent::__construct($hostname, $opts, $channel);
    }

    /**
     * Get all Customers with a filter - A server-to-client streaming RPC.
     * @param \Customer\CustomerFilter $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetCustomers(\Customer\CustomerFilter $argument,
      $metadata = [], $options = []) {
      return $this->_serverStreamRequest('/customer.Customer/GetCustomers',
      $argument,
      ['\Customer\CustomerRequest', 'decode'],
      $metadata, $options);
    }

    /**
     * Create a new Customer - A simple RPC.
     * @param \Customer\CustomerRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function CreateCustomer(\Customer\CustomerRequest $argument,
      $metadata = [], $options = []) {
      return $this->_simpleRequest('/customer.Customer/CreateCustomer',
      $argument,
      ['\Customer\CustomerResponse', 'decode'],
      $metadata, $options);
    }

  }

}
