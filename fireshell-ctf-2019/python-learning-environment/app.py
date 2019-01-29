import sys
import os
from cStringIO import StringIO
from flask import Flask, render_template, request

app = Flask(__name__)

LIMIT = 950

def attack(payload):
    payload = payload[:LIMIT]
    words = ['import', 'subprocess', 'os', 'imp', 'subp', 'rocess', 'subpr', 'ocess', 'bpr', 'catch_warnings', 'file', 'formatter', 'code', 'func', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '[',']', 'getitem', 'getattribute']
    for word in words:
        payload = payload.replace(word, "")
    scope = {"__builtins__" : {}}
    exec(payload) in scope
    return payload


@app.route('/', methods=['GET', 'POST'])
def index():
    code = request.form.get('code', """print('a')
test = 'a' + 'b'
print test""")
    raw_code = request.form.get('code', code)
    old_stdout = sys.stdout
    old_stderr = sys.stderr
    redirected_output = sys.stdout = StringIO()
    redirected_err = sys.stderr = StringIO()
    try:
        attack(raw_code)
    except Exception, e:
        print e
    finally: # !
        sys.stdout = old_stdout
        sys.stderr = old_stderr
    output = redirected_output.getvalue() + redirected_err.getvalue()
    new_output = ''
    for line in output.splitlines():
        new_output += ('> '+line+'\n')
    context = dict(raw_code=code, output=new_output)
    return render_template('index.html', **context)
