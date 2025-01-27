import numpy as np

class GiftRecommender:
    def __init__(self, embedding_generator, clickhouse_client, similarity_threshold=0.5):
        self.embedding_generator = embedding_generator
        self.clickhouse_client = clickhouse_client
        self.similarity_threshold = similarity_threshold

    def get_most_similar_gifts(self, description, top_n=1):
        description_embedding = self.embedding_generator.generate_embedding(description)
        query = "SELECT id, gift_embedding, gift_description FROM gifts"
        rows = self.clickhouse_client.execute_query(query)

        similarities = []
        unique_descriptions = set()
        for row in rows:
            gift_id, gift_embedding, gift_description = row

            if gift_description in unique_descriptions:
                continue

            if isinstance(gift_embedding, (list, np.ndarray)):
                gift_embedding = np.array(gift_embedding)
                similarity = np.dot(description_embedding, gift_embedding) / (
                    np.linalg.norm(description_embedding) * np.linalg.norm(gift_embedding)
                )

                if similarity >= self.similarity_threshold:
                    similarities.append((gift_id, similarity, gift_description))
                    unique_descriptions.add(gift_description)
            else:
                print("err")
                continue

        similarities.sort(key=lambda x: x[1], reverse=True)
        top_gift_ids = [gift_id for gift_id, _, _ in similarities[:top_n]]

        if top_gift_ids:
            query = f"""
                SELECT gift_name, gift_description, price, image_url
                FROM gifts
                WHERE id IN ({','.join(map(str, top_gift_ids))})
            """
            gift_details = self.clickhouse_client.execute_query(query)
        else:
            gift_details = []

        result = [
            {
                "gift_name": gift_name,
                "gift_description": gift_description,
                "price": price,
                "image_url": image_url,
            }
            for gift_name, gift_description, price, image_url in gift_details
        ]

        return result
