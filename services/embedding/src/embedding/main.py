import grpc
from concurrent import futures
from stub import embedding_pb2_grpc
from server import Embedder

def serve():
    model = Embedder() 
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    embedding_pb2_grpc.add_EmbedderServicer_to_server(model, server)

    server.add_insecure_port('[::]:50051')
    print("Embedding Service started on port 50051...")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    serve()