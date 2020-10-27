import os
import json
import sys
from pathlib import Path

if __name__ == "__main__":
    which_matrix = sys.argv[1]
    root_path = "e2etest/testcases"

    ec2_matrix = {"testing_suite": [], "testing_ami": ["amazonlinux2", "ubuntu16", "windows2019"]}
    ecs_matrix = {"testing_suite": [], "launch_type": ["EC2", "FARGATE"]}
    eks_matrix = {"testing_suite": []}
    matrix = {"ec2_matrix": ec2_matrix, "ecs_matrix": ecs_matrix, "eks_matrix": eks_matrix}

    # read the testcase directory
    testcase_dirs = os.listdir(root_path)
    testing_suites = testcase_dirs

    for testcase_dir in testcase_dirs:
        supported_platform_file = Path(root_path + "/" + testcase_dir + "/supported_platforms")
        if not supported_platform_file.exists():
            # which means all the platforms are supported for this test case
            ec2_matrix["testing_suite"].append(testcase_dir)
            ecs_matrix["testing_suite"].append(testcase_dir)
            eks_matrix["testing_suite"].append(testcase_dir)
            continue

        with open(root_path + "/" + testcase_dir + "/supported_platforms") as supported_platform_file:
            file_content = supported_platform_file.read()
            if 'EC2' in file_content:
                ec2_matrix["testing_suite"].append(testcase_dir)
            if 'ECS' in file_content:
                ecs_matrix["testing_suite"].append(testcase_dir)
            if 'EKS' in file_content:
                eks_matrix["testing_suite"].append(testcase_dir)

    print(json.dumps(matrix[which_matrix]))
                    
