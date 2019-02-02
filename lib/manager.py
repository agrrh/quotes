from lib.storage import Storage
from lib.quote import Quote


class Manager(object):
    def __init__(self, config):
        self.storage = Storage(**config.storage)

    def quote_create(self, text=None):
        if not text:
            return False

        q = Quote()
        q.text = text

        written = self.storage.write(key=q.id, data=q.__dict__)
        if not written:
            return False

        return q.__dict__

    def quote_read(self, id):
        q = Quote(id)

        data = self.storage.read(q.id)
        if not data:
            return False

        for key, val in data.items():
            setattr(q, key, val)

        return q.__dict__

    def quotes_list(self, offset=0, limit=20):
        return self.storage.list(offset=offset, limit=limit)

    def quote_delete(self, id):
        return self.storage.delete(int(id))
