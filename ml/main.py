import logging
from fastapi import FastAPI, Query
from pydantic import BaseModel
from typing import List
import os
from dotenv import load_dotenv
from embedding_generator import EmbeddingGenerator
from clickhouse_client import ClickHouseClient
from gift_recommender import GiftRecommender

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

load_dotenv()

app = FastAPI()

model_name = os.getenv("MODEL_NAME")
embedding_generator = EmbeddingGenerator(model_name)
clickhouse_client = ClickHouseClient()
gift_recommender = GiftRecommender(embedding_generator, clickhouse_client)

class GiftRequest(BaseModel):
    description: str
    top_n: int = Query(1, gt=0)

class GiftResponse(BaseModel):
    gift_name: str
    gift_description: str
    price: float
    image_url: str

@app.post("/recommend", response_model=List[GiftResponse])
def recommend_gifts(request: GiftRequest):
    logger.info("Получен запрос на рекомендацию подарков")
    most_similar_gifts = gift_recommender.get_most_similar_gifts(request.description, request.top_n)
    if most_similar_gifts:
        logger.info("Рекомендации успешно найдены")
        return most_similar_gifts
    else:
        logger.warning("Не удалось найти подарки")
        return []
        
@app.get("/health")
def health():
    return {"status": "ok"}

if __name__ == "__main__":
    logger.info("Запуск FastAPI приложения")
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
