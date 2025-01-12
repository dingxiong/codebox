from flask import Flask, render_template
from random import randint

app = Flask(__name__)

@app.route('/')
def index():
    return render_template('index.html')


@app.route('/hello')
def hello():
    return {"hello": "world"}

@app.route('/random')
def random():
    return {"hello": randint(0, 10000)}
