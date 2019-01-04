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

        data = self.storage.read(key=q.id)
        if not data:
            return False

        for key, val in data.items():
            setattr(q, key, val)

        return q.__dict__
