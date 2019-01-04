from lib.config import Config
from lib.api import API


if __name__ == '__main__':
    config = Config('config/config.yml')
    print(config.__dict__)
    api = API(config)
    api.run()
