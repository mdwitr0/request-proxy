from fastapi import Body, HTTPException, FastAPI
import requests

app = FastAPI()


@app.post("/")
async def proxy_handle(payload: dict = Body(...)):
    if not payload.get("method"):
        raise HTTPException(status_code=400, detail="method must not be empty")
    if not payload.get("url"):
        raise HTTPException(status_code=400, detail="url must not be empty")

    method = payload.get("method")
    url = payload.get("url")
    data = payload.get("data") if payload.get("data") else {}
    headers = payload.get("headers") if payload.get("headers") else {}
    params = payload.get("params") if payload.get("params") else {}

    response = requests.request(method, url, headers, data, params)

    return response.json()
