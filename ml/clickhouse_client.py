from clickhouse_driver import Client
import os
from dotenv import load_dotenv

load_dotenv()

class ClickHouseClient:
    def __init__(self):
        self.client = Client(
            os.getenv('CLICKHOUSE_HOST'),
            user=os.getenv('CLICKHOUSE_USER'),
            password=os.getenv('CLICKHOUSE_PASSWORD'),
            database=os.getenv('CLICKHOUSE_DATABASE')
        )

    def execute_query(self, query):
        return self.client.execute(query)