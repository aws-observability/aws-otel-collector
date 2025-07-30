import os
import json
import sys
from pathlib import Path

if __name__ == "__main__":
    which_matrix = sys.argv[1]
    testcase_json = "e2etest/testcases.json"

    ec2_matrix_1 = {"testcase": [], "testing_ami": [
        "amazonlinux2", 
        "ubuntu16", 
        "debian10", 
        "debian9", 
        "suse15"
    ]}
    ec2_matrix_2 = {"testcase": [], "testing_ami": [
        "suse12", 
        "redhat8", 
        "redhat7", 
        "centos7", 
        "windows2019"
    ]}
    ec2_matrix_3 = {"testcase": [], "testing_ami": [
        "arm_amazonlinux2",
        "arm_suse15",
        "arm_redhat8",
        "arm_redhat7",
        "arm_ubuntu16"
    ]}
    ecs_matrix = {"testcase": [], "launch_type": ["EC2", "FARGATE"]}
    
    # New EKS matrices for different versions
    eks_127_matrix = {"testcase": []}
    eks_128_matrix = {"testcase": []}
    eks_129_matrix = {"testcase": []}
    eks_130_matrix = {"testcase": []}
    eks_131_matrix = {"testcase": []}
    
    # New EKS ARM64 matrices for different versions
    eks_arm64_127_matrix = {"testcase": []}
    eks_arm64_128_matrix = {"testcase": []}
    eks_arm64_129_matrix = {"testcase": []}
    eks_arm64_130_matrix = {"testcase": []}
    eks_arm64_131_matrix = {"testcase": []}
    
    eks_fargate_matrix = {"testcase": []}
    eks_adot_operator_matrix = {"testcase": []}
    local_matrix = {"testcase": []}
    perf_matrix = {"testcase": [], "testing_ami": ["soaking_linux"], "data_rate": ["100", "1000", "5000"]}
    canary_matrix = {"testcase": [], "testing_ami": ["canary_linux", "canary_windows"]}

    matrix = {
            "ec2_matrix_1": ec2_matrix_1, 
            "ec2_matrix_2": ec2_matrix_2,
            "ec2_matrix_3": ec2_matrix_3,
            "ecs_matrix": ecs_matrix,
            "eks_127_matrix": eks_127_matrix,
            "eks_128_matrix": eks_128_matrix,
            "eks_129_matrix": eks_129_matrix,
            "eks_130_matrix": eks_130_matrix,
            "eks_131_matrix": eks_131_matrix,
            "eks_arm64_127_matrix": eks_arm64_127_matrix,
            "eks_arm64_128_matrix": eks_arm64_128_matrix,
            "eks_arm64_129_matrix": eks_arm64_129_matrix,
            "eks_arm64_130_matrix": eks_arm64_130_matrix,
            "eks_arm64_131_matrix": eks_arm64_131_matrix,
            "eks_fargate_matrix": eks_fargate_matrix,
            "eks_adot_operator_matrix": eks_adot_operator_matrix,
            "local_matrix": local_matrix,
            "perf_matrix": perf_matrix,
            "canary_matrix": canary_matrix
            }

    with open(testcase_json) as f:
        testcases = json.load(f)
        for testcase in testcases:
            if 'EC2' in testcase["platforms"]:
                ec2_matrix_1["testcase"].append(testcase["case_name"])
                ec2_matrix_2["testcase"].append(testcase["case_name"])
                ec2_matrix_3["testcase"].append(testcase["case_name"])
            if 'ECS' in testcase["platforms"]:
                ecs_matrix["testcase"].append(testcase["case_name"])
            # Handle version-specific EKS platforms
            if 'EKS_127' in testcase["platforms"]:
                eks_127_matrix["testcase"].append(testcase["case_name"])
            if 'EKS_128' in testcase["platforms"]:
                eks_128_matrix["testcase"].append(testcase["case_name"])
            if 'EKS_129' in testcase["platforms"]:
                eks_129_matrix["testcase"].append(testcase["case_name"])
            if 'EKS_130' in testcase["platforms"]:
                eks_130_matrix["testcase"].append(testcase["case_name"])
            if 'EKS_131' in testcase["platforms"]:
                eks_131_matrix["testcase"].append(testcase["case_name"])
            # Handle version-specific EKS ARM64 platforms
            if 'EKS_ARM64_127' in testcase["platforms"]:
                eks_arm64_127_matrix["testcase"].append(testcase["case_name"])
            if 'EKS_ARM64_128' in testcase["platforms"]:
                eks_arm64_128_matrix["testcase"].append(testcase["case_name"])
            if 'EKS_ARM64_129' in testcase["platforms"]:
                eks_arm64_129_matrix["testcase"].append(testcase["case_name"])
            if 'EKS_ARM64_130' in testcase["platforms"]:
                eks_arm64_130_matrix["testcase"].append(testcase["case_name"])
            if 'EKS_ARM64_131' in testcase["platforms"]:
                eks_arm64_131_matrix["testcase"].append(testcase["case_name"])
            if 'EKS_FARGATE' in testcase["platforms"]:
                eks_fargate_matrix["testcase"].append(testcase["case_name"])
            if 'EKS_ADOT_OPERATOR' in testcase["platforms"]:
                eks_adot_operator_matrix["testcase"].append(testcase["case_name"])
            if 'LOCAL' in testcase["platforms"]:
                local_matrix["testcase"].append(testcase["case_name"])
            if 'PERF' in testcase["platforms"]:
                perf_matrix["testcase"].append(testcase["case_name"])
            if 'CANARY' in testcase["platforms"]:
                canary_matrix["testcase"].append(testcase["case_name"])

    print(json.dumps(matrix[which_matrix]))
