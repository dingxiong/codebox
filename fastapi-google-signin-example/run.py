import os
import logging

import uvicorn
from authlib.integrations.starlette_client import OAuth
from authlib.integrations.starlette_client import OAuthError
from fastapi import FastAPI
from fastapi import Request
from starlette.config import Config
from starlette.middleware.sessions import SessionMiddleware
from starlette.responses import HTMLResponse
from starlette.responses import RedirectResponse

logger = logging.getLogger("uvicorn")
# Create the APP
app = FastAPI()

# Set up OAuth
starlette_config = Config('.env')
print(starlette_config.__dict__)
oauth = OAuth(starlette_config)
oauth.register(
    name='google',
    server_metadata_url='https://accounts.google.com/.well-known/openid-configuration',
    client_kwargs={'scope': 'openid email profile'},
)

# Set up the middleware to read the request session
SECRET_KEY = os.environ.get('SECRET_KEY') or "OulLJiqkldb436-X6M11hKvr7wvLyG8TPi5PkLf4"
if SECRET_KEY is None:
    raise 'Missing SECRET_KEY'
app.add_middleware(SessionMiddleware, secret_key=SECRET_KEY)


@app.get('/')
def public(request: Request):
    logger.info("inside /")
    user = request.session.get('user')
    if user:
        name = user.get('name')
        return HTMLResponse(f'<p>Hello {name}!</p><a href=/logout>Logout</a>')
    return HTMLResponse('<a href=/login>Login</a>')


@app.route('/logout')
async def logout(request: Request):
    request.session.pop('user', None)
    return RedirectResponse(url='/')


@app.route('/login')
async def login(request: Request):
    redirect_uri = request.url_for('auth')  # This creates the url for our /auth endpoint
    return await oauth.google.authorize_redirect(request, redirect_uri)


@app.route('/auth')
async def auth(request: Request):
    try:
        access_token = await oauth.google.authorize_access_token(request)
    except OAuthError as e:
        logger.info("OAuthError %s", e)
        return RedirectResponse(url='/')
    logger.info("access_token %s", access_token)
    assert "userinfo" in access_token
    request.session['user'] = dict(access_token["userinfo"])
    return RedirectResponse(url='/')


if __name__ == '__main__':
    uvicorn.run(app, port=7000)
