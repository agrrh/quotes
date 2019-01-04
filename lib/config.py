import yaml


class Config(object):
    def __init__(self, path):
        self.path = path

        self.load()

    def load(self):
        with open(self.path) as fp:
            data = yaml.load(fp)
        for key, val in data.items():
            setattr(self, key, val)
