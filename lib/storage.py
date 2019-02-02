import redis
import json


class Storage(object):
    def __init__(self, host='127.0.0.1', port=6379):
        self.handler = redis.Redis(host=host, port=port)

    def read(self, id):
        data = self.handler.zrangebyscore('quotes', id, id)
        if isinstance(data, list) and len(data) < 1:
            return {}
        return json.loads(data[0])

    def write(self, key, data):
        id = int(data.get('id'))
        data = json.dumps(data)
        return self.handler.zadd('quotes', {data: id})

    def list(self, offset=0, limit=20):
        limit = int(limit) - 1 if limit > 0 else 0
        for data in self.handler.zrange('quotes', offset, offset + limit, desc=True):
            yield(json.loads(data))

    def delete(self, id):
        return bool(self.handler.zremrangebyscore('quotes', id, id))
