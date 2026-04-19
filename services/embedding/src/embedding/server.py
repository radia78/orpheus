from stub import embedding_pb2_grpc
from stub import embedding_pb2

import torch
from transformers import AutoModel
import torch.nn.functional as F
from torchvision import transforms
from PIL import Image

from typing import List

class Embedder(embedding_pb2_grpc.EmbedderServicer):
    def __init__(self):
        super().__init__()
        self.model = AutoModel.from_pretrained("google/tipsv2-so400m14", trust_remote_code=True)
        self.model.eval()
        self.transform = transforms.Compose([
            transforms.Resize((448, 448)),
            transforms.ToTensor(),
        ])

    def process_images(self, images: List[str]) -> torch.Tensor:
        img = list(map(Image.open, images))
        return torch.stack(list(map(self.transform, img)), dim=-1)

    def EmbedText(self, request, context):
        texts = list(request.texts)
        response = embedding_pb2.EmbeddingResponse()
        embeddings = F.normalize(self.model.encode_text(texts), dim=-1)
        for emb in embeddings:
            v_msg = embedding_pb2.Vector(values=emb.tolist())
            response.vectors.append(v_msg)

        return response

    def EmbedImage(self, request, context):
        img_links = list(request.images)
        response = embedding_pb2.EmbeddingResponse()
        img = self.process_images(img_links)
        embeddings = F.normalize(self.model.encode_image(img), dim=-1)
        for emb in embeddings:
            v_msg = embedding_pb2.Vector(values=emb.tolist())
            response.vectors.append(v_msg)

        return response