#!/usr/bin/env python

import json
import matplotlib.pyplot as plt
import numpy as np

lines = []
with open("benchmark.jsonl", "r") as f:
    for line in f:
        lines.append(json.loads(line))

data = {}


def action_start(line):
    data[line["Package"]] = {}


def action_output(line):
    pass


def action_skip(line):
    pass


def action_pass(line):
    pass


def action_run(line):
    print(line)


action_callbacks = {
    "start": action_start,
    "output": action_output,
    "skip": action_skip,
    "pass": action_pass,
    "run": action_run,
}

for line in lines:
    action_callbacks[line["Action"]](line)

print(data)
