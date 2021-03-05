
import json
import sys
from os import listdir
from os.path import isfile, join

from simpleeval import simple_eval

if __name__ == '__main__':
    inputs = sys.argv[1].split("|")
    params = {}

    results = []
    for item in inputs:
        result = {}
        key = item.split(" = ")[0]
        expression = item.split(" = ")[1]
        value = simple_eval(expression,
                names=params) 
        result['key'] = key
        result['rawvalue'] = value
        results.append(result)
        params[key] = value

    print(json.dumps(results))
