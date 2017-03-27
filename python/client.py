import time

import grpc

import customer_pb2
import customer_pb2_grpc

def customer_create_customer(stub):
  """Calls the RPC method CreateCustomer of CustomerServer.

  Keyword arguments:
  stub -- The Customer service definition.
  """

  customer_one = customer_pb2.CustomerRequest(
    id=101,
    name="Shiju Varghese",
    email="shiju@xyz.com",
    phone="732-757-2923",
    addresses=[
      {
        "street": "1 Mission Street",
        "city": "San Francisco",
        "state": "CA",
        "zip": "94105",
      },
      {
        "street": "Greenfield",
        "city":"Kochi",
        "state": "KL",
        "zip": "68356",
        "isShippingAddress": True,
      }
    ]
  )
  response = stub.CreateCustomer(customer_one)
  print("A new Customer has been added with id: %d" % response.id)

  customer_two = customer_pb2.CustomerRequest(
    id=102,
    name="Irene Rose",
    email="irene@xyz.com",
    phone="732-757-2924",
    addresses=[
      {
        "street": "1 Mission Street",
        "city": "San Francisco",
        "state": "CA",
        "zip": "94105",
        "isShippingAddress": True,
      }
    ]
  )
  response = stub.CreateCustomer(customer_two)
  print("A new Customer has been added with id: %d" % response.id)

def customer_get_customers(stub):
  """Calls the RPC method GetCustomers of CustomerServer.

  Keyword arguments:
  stub -- The Customer service definition.
  """

  """Example of filtering by name:
  keyword: customer_pb2.CustomerFilter(keyword=Rose")"""
  filter = customer_pb2.CustomerFilter(keyword="")

  customers = stub.GetCustomers(filter)

  for customer in customers:
    customer_str = "Customer: id = %d, name = %s, email = %s, phone = %s, addresses = [" % (
      customer.id,
      customer.name,
      customer.email,
      customer.phone
    );

    for address in customer.addresses:
      customer_str += "(street = %s, city = %s, state = %s, zip = %s), " % (
        address.street,
        address.city,
        address.state,
        address.zip
      )

    customer_str = customer_str[0:-2]
    customer_str += "]";

    print(customer_str)

def run():
  channel = grpc.insecure_channel('localhost:50051')
  stub = customer_pb2_grpc.CustomerStub(channel)

  print("-------------- CreateCustomer --------------")
  customer_create_customer(stub)
  print("-------------- GetCustomers --------------")
  customer_get_customers(stub)

if __name__ == '__main__':
  run()
