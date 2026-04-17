import os
from concurrent import futures

import grpc

from app.service import RewriterServiceLogic
from gen import rewrite_pb2, rewrite_pb2_grpc


class RewriterGRPC(rewrite_pb2_grpc.RewriterServicer):
    def __init__(self):
        self.logic = RewriterServiceLogic()

    def Rewrite(self, request, context):
        if not request.text or len(request.text) > 600:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("text too long or empty")
            return rewrite_pb2.RewriteResponse()

        if request.style not in ["official", "humorous", "friendly", "professional"]:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("unsupported style")
            return rewrite_pb2.RewriteResponse()

        result = self.logic.rewrite(text=request.text, style=request.style)

        return rewrite_pb2.RewriteResponse(**result)


def create_server() -> grpc.Server:
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    rewrite_pb2_grpc.add_RewriterServicer_to_server(RewriterGRPC(), server)

    port = os.getenv("GRPC_PORT", "50051")
    host = os.getenv("GRPC_HOST", "0.0.0.0")

    server.add_insecure_port(f"{host}:{port}")

    return server
