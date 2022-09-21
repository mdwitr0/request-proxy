import json
import asyncio
import grequests
from aiohttp import web


async def handle(request):
    payload = json.loads(await request.read())
    if not payload.get("method"):
        raise "method must not be empty"
    if not payload.get("url"):
        raise "url must not be empty"

    method = payload.get("method")
    url = payload.get("url")
    data = json.dumps(payload.get("data"))
    headers = payload.get("headers")
    params = payload.get("params")

    try:
        response = grequests.request(method=method, url=url, headers=headers, data=data, params=params)
        return web.json_response(response.json())
    except Exception as e:
        print("method: ", method)
        print("url: ", url)
        print("params: ", params)
        print("data: ", data)
        print("headers: ", headers)

        print(e)
        return web.json_response({})

app = web.Application()
app.add_routes([web.post('/', handle)])

if __name__ == '__main__':
    web.run_app(app)