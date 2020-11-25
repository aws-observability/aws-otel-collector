import os
import json
import sys
from pathlib import Path
import jinja2
import string


if __name__ == "__main__":
    from jinja2 import Environment, PackageLoader, select_autoescape
    templateLoader = jinja2.FileSystemLoader(searchpath="e2etest/templates/")
    env = Environment(loader=templateLoader)

    # get performance models from artifacts
    performance_models = []
    artifacts_path = "artifacts/"
    for sub_dir in os.listdir(artifacts_path):
        with open(artifacts_path + sub_dir + "/performance.json") as f:
            model = json.load(f)
            model["avgCpu"] = "{:.2f}".format(model["avgCpu"])
            model["avgMem"] = "{:.2f}".format(model["avgMem"])
            performance_models.append(model)

    # sort by testing_ami, name and rate
    performance_models = sorted(performance_models, key = lambda x: (x["testingAmi"], x["testcase"], x["dataRate"]))

    # render performance models into markdown
    template = env.get_template('performance_model.tpl')
    rendered_result = template.render({"performance_models": performance_models})
    print(rendered_result)

    # write rendered result to docs/performance_model.md
    with open("docs/performance_model.md", "w") as f:
        f.write(rendered_result)
