"""The Python implementation of the gRPC Customer server."""

from concurrent import futures
import time

import grpc

import customer_pb2
import customer_pb2_grpc

_ONE_DAY_IN_SECONDS = 60 * 60 * 24

customers = []

class CustomerServicer(customer_pb2_grpc.CustomerServicer):
  """Provides methods that implement functionality of customer server."""

  def CreateCustomer(self, request, context):
    """Creates a new Customer."""

    customers.append(request)

    return customer_pb2.CustomerResponse(id=request.id, success=True)

  def GetCustomers(self, request, context):
    """Filter the list of customers."""

    for customer in customers:
      if request.keyword != "":
        if request.keyword not in customer.name:
          continue

      yield customer

def serve():
  """Start the server."""
  server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
  customer_pb2_grpc.add_CustomerServicer_to_server(CustomerServicer(), server)
  server.add_insecure_port('[::]:50051')
  server.start()

  try:
    while True:
      time.sleep(_ONE_DAY_IN_SECONDS)
  except KeyboardInterrupt:
    server.stop(0)

if __name__ == '__main__':
  serve()
