import sys
import json
import numpy as np
import copy

def calculate(node):
    vector = get_vector(node, sys.argv[1])
    if node['children'] is not None and node['children'] != []:
        result_vector = np.array([0] * len(node['children']))
        for idx, child in enumerate(node['children']):
            child_vector = np.array(calculate(child))
            result_vector = result_vector + vector[idx] * child_vector
        return result_vector
    else:
        return vector

def get_vector(node, method):
    if method == 'norm':
        return norm(node)
    elif method == 'gmm':
        return gmm(node)

def gmm(node):
    return None

def norm(node):
    matrix = np.array(node['preferences'])
    matrix = matrix / matrix.sum(axis=0)
    out_vector = np.mean(matrix, axis=1)
    return out_vector / out_vector.sum()


method = sys.argv[1]

file = open('ahp.json', 'r+')
data = json.loads(file.read())

print(calculate(data['goal']))
# print(calculate(data['goal']))

