[
    {
      "name": "aoc-collector",
      "image": "${aoc_image}",
      "cpu": 10,
      "memory": 256,
      "secrets": [
        {
            "name": "AOT_CONFIG_CONTENT",
            "valueFrom": "${ssm_parameter_arn}"
        }
      ],
      "essential": true,
      "entryPoint": [],
      "command": [],
      "environment": [],
      "environmentFiles": [],
      "dependsOn": [],
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "/ecs/ecs-cwagent-sidecar-collector",
          "awslogs-region": "${region}",
          "awslogs-stream-prefix": "ecs",
          "awslogs-create-group": "True"
        }
      }
    }
]
