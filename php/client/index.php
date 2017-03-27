<?php
/**
 * @file
 * Test PHP client for gRPC.
 */

use Customer\CustomerClient;
use Customer\CustomerFilter;
use Customer\CustomerRequest;
use Customer\CustomerRequest_Address;
use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;

require dirname(__FILE__) . '/vendor/autoload.php';

/**
 * Calls the RPC method CreateCustomer of CustomerServer.
 *
 * @param CustomerClient $client
 *   The protocol buffer client.
 * @param CustomerRequest $customer
 *   The request message to create the customer.
 */
function createCustomer(CustomerClient $client, CustomerRequest $customer) {
    list($response, $status) = $client->CreateCustomer($customer)->wait();

    print sprintf("Created customer: id = %d\n", $response->getId());
}

/**
 * Calls the RPC method GetCustomers of CustomerServer.
 *
 * @param CustomerClient $client
 *   The protocol buffer client.
 * @param CustomerFilter $filter
 *   Filter the customers list by this keyword.
 */
function getCustomers(CustomerClient $client, CustomerFilter $filter) {
    // Start the server streaming call.
    $call = $client->GetCustomers($filter);
    $customers = $call->responses();

    foreach ($customers as $customer) {
        $customerStr = sprintf("Customer: id = %d, name = %s, email = %s, phone = %s, addresses = [",
            $customer->getId(),
            $customer->getName(),
            $customer->getEmail(),
            $customer->getPhone());

        foreach ($customer->getAddresses() as $address) {
            $customerStr .= sprintf("(street = %s, city = %s, state = %s, zip = %s), ",
                $address->getStreet(),
                $address->getCity(),
                $address->getState(),
                $address->getZip());
        }

        $customerStr = substr($customerStr, 0, strlen($customerStr) - 2);
        $customerStr .= "]\n";
        print $customerStr;
    }
}

// Create a new client.
$client = new CustomerClient('localhost:50051', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

// Build an address.
$addressOne = new CustomerRequest_Address();
$addressOne->setStreet("1 Mission Street");
$addressOne->setCity("San Francisco");
$addressOne->setState("CA");
$addressOne->setZip("94105");

// Build another address.
$addressTwo = new CustomerRequest_Address();
$addressTwo->setStreet("Greenfield");
$addressTwo->setCity("Kochi");
$addressTwo->setState("KL");
$addressTwo->setZip("68356");
$addressTwo->setIsShippingAddress(true);

// Build the first customer.
$customerOne = new CustomerRequest();
$customerOne->setId(101);
$customerOne->setName("Shiju Varghese");
$customerOne->setEmail("shiju@xyz.com");
$customerOne->setPhone("732-757-2923");

// Set the addresses.
$customerOneAddresses = new RepeatedField(GPBType::MESSAGE, CustomerRequest_Address::class);
$customerOneAddresses[] = $addressOne;
$customerOneAddresses[] = $addressTwo;
$customerOne->setAddresses($customerOneAddresses);

// Create the customer.
createCustomer($client, $customerOne);

// Build a second customer.
$customerTwo = new CustomerRequest();
$customerTwo->setId(102);
$customerTwo->setName("Irene Rose");
$customerTwo->setEmail("irene@xyz.com");
$customerTwo->setPhone("732-757-2924");

// Set the addresses.
$customerTwoAddresses = new RepeatedField(GPBType::MESSAGE, CustomerRequest_Address::class);
$customerTwoAddresses[] = $addressOne;
$customerTwo->setAddresses($customerTwoAddresses);

// Create the customer.
createCustomer($client, $customerTwo);

// Filter with an empty Keyword
$filter = new CustomerFilter();

// Example filter by surname:
// $filter->setKeyword("Rose");
$filter->setKeyword("");

// Get the filtered list of customers.
getCustomers($client, $filter);