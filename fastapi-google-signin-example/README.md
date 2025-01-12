<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Contents**

- [A simple fastapi app with google signon](#a-simple-fastapi-app-with-google-signon)
  - [Steps:](#steps)
  - [sample acess_token returned from google](#sample-acess_token-returned-from-google)
  - [References](#references)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# A simple fastapi app with google signon

## Steps:

1. Create google client id and secret in google cloud console. Make sure both
   http://localhost and http://localhost:<port number> are added to authorized
   javascript origins.
2. install uvicorn, fastapi, authlib, etc
3. run the program uvicorn run:app --port 8004 --reload

## sample acess_token returned from google

```
{
  'access_token': 'xxxx',
  'expires_in': 3598,
  'scope': 'https://www.googleapis.com/auth/userinfo.profile openid https://www.googleapis.com/auth/userinfo.email',
  'token_type': 'Bearer',
  'id_token': '<a very long string>'
  'expires_at': 1655229535,
  'userinfo': {
    'iss': 'https://accounts.google.com',
    'azp': 'xxxx'
    'aud': 'xxxx'
    'sub': '117054356680777545659',
    'email': 'dingxiong203@gmail.com',
    'email_verified': True,
    'at_hash': 'xxx'
    'nonce': 'xxxx'
    'name': 'xiong ding',
    'picture': 'https://lh3.googleusercontent.com/a-/AOh14GixC-t2B81vrZYT4vTnQAOWfLZ1zZOn4-jJRH2TxQ=s96-c',
    'given_name': 'xiong',
    'family_name': 'ding',
    'locale': 'en',
    'iat': 1655225937,
    'exp': 1655229537
  }
}
```

## References

- https://blog.hanchon.live/guides/google-login-with-fastapi/
- https://developers.google.com/identity/gsi/web/guides/overview
