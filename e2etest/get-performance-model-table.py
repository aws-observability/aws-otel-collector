import os
import json
import sys
from pathlib import Path
import jinja2
import string

# Schema: performance_models[data_mode][tps] = [model]
performance_models = {}


def add_performance_model(model):
    """
    Adds performance model to the list of performance models.
    """
    # process model values
    model["avgCpu"] = "{:.2f}".format(model["avgCpu"])
    model["avgMem"] = "{:.2f}".format(model["avgMem"])
    model["maxCpu"] = "{:.2f}".format(model["maxCpu"])
    model["maxMem"] = "{:.2f}".format(model["maxMem"])
    model["receivers"].sort()
    model["processors"].sort()
    model["exporters"].sort()
    model["receivers"] = ", ".join(model["receivers"])
    model["processors"] = ", ".join(model["processors"])
    model["exporters"] = ", ".join(model["exporters"])

    data_mode = model["dataMode"].capitalize()
    data_rate = model["dataRate"]

    if data_mode not in performance_models:
        performance_models[data_mode] = {}
    if data_rate not in performance_models[data_mode]:
        performance_models[data_mode][data_rate] = []

    performance_models[data_mode][data_rate].append(model)


def flatten_performance_models(models):
    """
    Flattens performance model into list of grouped models where each group
    corresponds to a table in the report.
    """
    models_list = []

    for data_mode, data_rates in performance_models.items():
        for data_rate, models in data_rates.items():
            model = {}
            model["data_mode"] = data_mode
            model["data_rate"] = data_rate
            # sort by name and type
            model["models"] = sorted(
                models, key=lambda x: (x["receivers"], x["testcase"], x["dataType"]))
            models_list.append(model)

    # sort by data mode and rate
    models_list = sorted(models_list, key=lambda x: (
        x["data_mode"], x["data_rate"]))
    return models_list


if __name__ == "__main__":
    aoc_version = Path('VERSION').read_text()

    from jinja2 import Environment, PackageLoader, select_autoescape
    templateLoader = jinja2.FileSystemLoader(searchpath="e2etest/templates/")
    env = Environment(autoescape=select_autoescape(['html', 'xml', 'tpl', 'yaml', 'yml']), loader=templateLoader)

    # get performance models from artifacts
    artifacts_path = "artifacts/"
    commit_id = ""
    collection_period = None
    testing_ami = ""
    for sub_dir in os.listdir(artifacts_path):
        with open(artifacts_path + sub_dir + "/performance.json") as f:
            model = json.load(f)
            commit_id = model["commitId"]
            collection_period = model["collectionPeriod"]
            testing_ami = model["testingAmi"]
            add_performance_model(model)

    models_list = flatten_performance_models(performance_models)

    # render performance models into markdown
    template = env.get_template('performance_model.tpl')
    rendered_result = template.render({
        "aoc_version": aoc_version,
        "commit_id": commit_id,
        "collection_period": collection_period,
        "testing_ami": testing_ami,
        "models_list": models_list,
    })
    print(rendered_result)

    # write rendered result to docs/performance_model.md
    with open("docs/performance_model.md", "w") as f:
        f.write(rendered_result)
