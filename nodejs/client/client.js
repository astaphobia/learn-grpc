var PROTO_PATH = __dirname + '/../../protos/customer.proto'

var _ = require('lodash');
var async = require('async');
var grpc = require('grpc');

var customer_proto = grpc.load(PROTO_PATH).customer;
var client = new customer_proto.Customer('localhost:50051', grpc.credentials.createInsecure());

/**
 * runCreateCustomer - Calls the RPC method CreateCustomer of CustomerServer.
 * @param {function} callback Called when this demo is complete
 */
function runCreateCustomer(callback) {
  var next = _.after(2, callback);

  function customerCallback(error, customer) {
    if (error) {
      console.log(error)
      callback(error);
      return;
    }

    console.log('A new Customer has been added with id: ' + customer.id)

    next();
  }

  customer_one = {
    id: 101,
    name: "Shiju Varghese",
    email: "shiju@xyz.com",
    phone: "732-757-2923",
    addresses: [
      {
        street: "1 Mission Street",
        city: "San Francisco",
        state: "CA",
        zip: "94105",
      },
      {
        street: "Greenfield",
        city:"Kochi",
        state: "KL",
        zip: "68356",
        isShippingAddress: true,
      }
    ]
  }
  client.createCustomer(customer_one, customerCallback);

  customer_two = {
    id: 102,
    name: "Irene Rose",
    email: "shiju@xyz.com",
    phone: "732-757-2924",
    addresses: [
      {
        street: "Greenfield",
        city:"Kochi",
        state: "KL",
        zip: "68356",
        isShippingAddress: true,
      }
    ]
  }
  client.createCustomer(customer_two, customerCallback);
}


/**
 * runGetCustomers - Calls the RPC method GetCustomers of CustomerServer.
 * @param {function} callback Called when this demo is complete
 */
function runGetCustomers(callback) {
  var filter = {
    // Example of filtering by name:
  	// keyword: "Rose"
    keyword: ""
  };

  var call = client.getCustomers(filter);
  call.on('error', function(error) {
    console.log('An error occurred: ' + error);
  });
  call.on('data', function(customer) {
    var customer_str = 'Customer: id = ' + customer.id + ', name = ' + customer.name + ', email = ' + customer.email + ', phone = ' + customer.phone + ', addresses = [';

    _.each(customer.addresses, function(address) {
      customer_str += '(street = ' + address.street + ', city = ' + address.city + ', state = ' + address.state + ', zip = ' + address.zip + '), ';
    });

    customer_str = customer_str.substring(0, customer_str.length - 2);
    customer_str += ']';

    console.log(customer_str)
  });
  call.on('end', callback);
}

/**
 * Run all of the demos in order
 */
function main() {
  async.series([
    runCreateCustomer,
    runGetCustomers,
  ]);
}

if (require.main === module) {
  main();
}

exports.runCreateCustomer = runCreateCustomer;

// exports.runGetCustomers = runGetCustomers;
