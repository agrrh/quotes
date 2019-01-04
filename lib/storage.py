import redis
import json


class Storage(object):
    def __init__(self, host='127.0.0.1', port=6379):
        self.handler = redis.Redis(host=host, port=port)

    def read(self, key):
        data = self.handler.get(key)
        return json.loads(data) if data else {}

    def write(self, key, data):
        data = json.dumps(data)
        return self.handler.set(key, data)
