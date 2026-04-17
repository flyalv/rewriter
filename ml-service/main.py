import signal
import sys

from app.server import create_server


def main():
    server = create_server()
    server.start()
    print("gRPC server on 0.0.0.0:50051")

    def stop_server(*_):
        print("shutting down...")
        server.stop(5)
        sys.exit(0)

    signal.signal(signal.SIGINT, stop_server)
    signal.signal(signal.SIGTERM, stop_server)

    server.wait_for_termination()


if __name__ == "__main__":
    main()
