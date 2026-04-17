import logging

import grpc
import ollama

logger = logging.getLogger(__name__)


class RewriterServiceLogic:
    @staticmethod
    def rewrite(text: str, style: str) -> dict:
        logger.info(f"Rewrite request: style={style}, text_len={len(text)}")

        try:
            response = ollama.chat(
                model="phi3:mini",
                messages=[
                    {
                        "role": "user",
                        "content": f"Rewrite in {style} style (only result): {text}",
                    }
                ],
            )
        except ollama.ResponseError as e:
            logger.error(f"Ollama error: {e}")
            raise grpc.RpcError(grpc.StatusCode.INTERNAL, "LLM failed")
        except Exception as e:
            logger.error(f"Unexpected error: {e}")
            raise grpc.RpcError(grpc.StatusCode.UNAVAILABLE, "LLM unavailable")

        rewritten = response["message"]["content"]

        return {
            "original_text": text,
            "rewritten_text": rewritten,
            "applied_style": style,
        }
