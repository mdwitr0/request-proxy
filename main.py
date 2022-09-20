from fastapi import Body, HTTPException, FastAPI
import requests
import json

app = FastAPI()


@app.post("/")
async def proxy_handle(payload: dict = Body(...)):
    if not payload.get("method"):
        raise HTTPException(status_code=400, detail="method must not be empty")
    if not payload.get("url"):
        raise HTTPException(status_code=400, detail="url must not be empty")

    method = payload.get("method")
    url = payload.get("url")
    data = json.dumps(payload.get("data"))
    headers = payload.get("headers")
    params = payload.get("params")

    print("method: ", method)
    print("url: ", url)
    print("params: ", params)
    print("data: ", data)
    print("headers: ", headers)

    try:
        response = requests.request(method=method, url=url, headers=headers, data=data, params=params)

        if response.status_code == 200:
            return response.json()
        else:
            raise HTTPException(status_code=400, detail="Bad Request")

    except Exception as e:
        print(e)
        return {}
