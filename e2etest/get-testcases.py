import os
import json
import sys
from pathlib import Path

if __name__ == "__main__":
    which_matrix = sys.argv[1]
    testcase_json = "e2etest/testcases.json"

    ec2_matrix = {"testcase": [], "testing_ami": ["amazonlinux2", "ubuntu16", "windows2019"]}
    ecs_matrix = {"testcase": [], "launch_type": ["EC2", "FARGATE"]}
    eks_matrix = {"testcase": []}
    matrix = {"ec2_matrix": ec2_matrix, "ecs_matrix": ecs_matrix, "eks_matrix": eks_matrix}

    with open(testcase_json) as f:
        testcases = json.load(f)
        for testcase in testcases:
            if 'EC2' in testcase["platforms"]:
                ec2_matrix["testcase"].append(testcase["case_name"])
            if 'ECS' in testcase["platforms"]:
                ecs_matrix["testcase"].append(testcase["case_name"])
            if 'EKS' in testcase["platforms"]:
                eks_matrix["testcase"].append(testcase["case_name"])

    print(json.dumps(matrix[which_matrix]))
