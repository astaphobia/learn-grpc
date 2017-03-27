var PROTO_PATH = __dirname + '/../../protos/customer.proto'

var _ = require('lodash');
var grpc = require('grpc');

var customer_proto = grpc.load(PROTO_PATH).customer;
var customer_list = [];

/**
 * createCustomer request handler. Creates a new Customer.
 *
 * @param {EventEmitter} call Call object for the handler to process
 * @param {function(Error, feature)} callback Response callback
 */
function createCustomer(call, callback) {
    // Add to the customer list.
    customer_list.push(call.request);

    // Create the response.
    response = {
        id: call.request.id,
        success: true
    }
    callback(null, response)
}


/**
 * getCustomers request handler. Get a filtered list of customers.
 *
 * @param {Writable} call Writable stream for responses with an additional
 *     request property for the request value.
 */
function getCustomers(call) {
    // Filter the list of customers.
    _.each(customer_list, function(customer) {
      if (call.request.keyword != "") {
        if (customer.name.indexOf(call.request.keyword) == -1) {
          return false
        }
      }

      call.write(customer);
    });

  call.end();
}

/**
 * Get a new server with the handler functions in this file bound to the methods
 * it serves.
 * @return {Server} The new server object
 */
function getServer() {
  var server = new grpc.Server();
  server.addProtoService(customer_proto.Customer.service, {
    createCustomer: createCustomer,
    getCustomers: getCustomers
  });
  return server;
}

if (require.main === module) {
  // If this is run as a script, start a server on an unused port
  var customerServer = getServer();
  customerServer.bind('0.0.0.0:50051', grpc.ServerCredentials.createInsecure());
  customerServer.start();
}

exports.getServer = getServer;
