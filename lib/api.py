from japronto import Application

from lib.manager import Manager


class API(object):
    def __init__(self, config):
        self.config = config

        self.app = Application()

        self.manager = Manager(config)

        self.app.router.add_route('/quotes/', self.route_quotes_list, methods=['GET'])
        self.app.router.add_route('/quotes/', self.route_quote_create, methods=['PUT'])
        self.app.router.add_route('/quotes/{id}', self.route_quote_read, methods=['GET'])
        self.app.router.add_route('/quotes/{id}', self.route_quote_delete, methods=['DELETE'])

    def __form_response(self, request, code=200, message=None, data=None):
        json = {}
        if message:
            json['message'] = str(message)
        if data:
            json['data'] = data
        return request.Response(
            code=code,
            json=json
        )

    def route_quotes_list(self, request):
        limit = int(request.query.get('limit', 20))
        offset = int(request.query.get('offset', 0))

        data = [q for q in self.manager.quotes_list(limit=limit, offset=offset)]

        return self.__form_response(
            request,
            code=200,
            message='OK',
            data=data
        )

    def route_quote_create(self, request):
        text = request.json.get('text', None)
        if not text:
            return self.__form_response(
                request,
                code=400,
                message='Malformed request'
            )

        data = self.manager.quote_create(text=text)
        if not data:
            return self.__form_response(
                request,
                code=500,
                message='Error writing data'
            )

        return self.__form_response(
            request,
            code=201,
            message='Created',
            data=data
        )

    def route_quote_read(self, request):
        id = request.match_dict.get('id')
        if not id:
            return self.__form_response(
                request,
                code=400,
                message='Malformed request'
            )

        data = self.manager.quote_read(id)
        if not data:
            return self.__form_response(
                request,
                code=404,
                message='Not found'
            )

        return self.__form_response(
            request,
            code=200,
            message='OK',
            data=data
        )

    def route_quote_delete(self, request):
        # TODO add auth
        id = request.match_dict.get('id')
        if not id:
            return self.__form_response(
                request,
                code=400,
                message='Malformed request'
            )

        result = self.manager.quote_delete(id)
        if not result:
            return self.__form_response(
                request,
                code=404,
                message='Not found'
            )

        return self.__form_response(
            request,
            code=200,
            message='OK'
        )

    def run(self):
        self.app.run(
            host=self.config.api.get('host', '127.0.0.1'),
            port=self.config.api.get('port', 8080),
            debug=self.config.debug
        )
