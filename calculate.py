import sys
import json
import numpy as np
import copy

method = sys.argv[1]

def calculate(node):
    vector = get_vector(node, sys.argv[1])
    if node.children is not None:
        result_vector = np.array([0] * len(node.children))
        for idx, child in node.children:
            child_vector = np.array(calculate(child))
            result_vector = result_vector + vector[idx] * child_vector
        return result_vector
    else:
        return vector


def get_vector(node, method):
    if method == 'norm':
        pass
    elif method == 'gmm':
        pass
        
        
with open('ahp.json', 'r') as file:
    data=json.loads(file.read())

head = data.goal

