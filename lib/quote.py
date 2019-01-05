import time


class Quote(object):
    def __init__(self, id=None):
        self.id = id or int(time.time() * 1000)
        self.text = None
        self.rating = 0

    def __id_randomize(self):
        time_ = time.time()
        id = (
            'q',
            str(int(time_)),
            hex(hash(time_))[2:6],  # poor man's "random" hash
        )
        return '_'.join(id)
