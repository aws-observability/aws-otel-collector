#!/usr/bin/env python3
# Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License").
# You may not use this file except in compliance with the License.
# A copy of the License is located at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# or in the "license" file accompanying this file. This file is distributed
# on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
# express or implied. See the License for the specific language governing
# permissions and limitations under the License.

import sys
import boto3

if __name__ == "__main__":

    if len(sys.argv) < 5:
        print("Usage: %s package_name version s3_bucket region" % sys.argv[0])
        sys.exit()
    pkg_name = sys.argv[1]
    rel_ver = sys.argv[2]
    s3_bucket = sys.argv[3]
    region = sys.argv[4]

    client = boto3.client('ssm', region_name=region)

    response = client.list_documents(
        Filters=[
            {
                'Key': 'Owner',
                'Values': ['Self']
            },
            {
                'Key': 'DocumentType',
                'Values': ['Package']
            },
            {
                'Key': 'Name',
                'Values': [pkg_name]
            }
        ]
    )

    if len(response['DocumentIdentifiers']) == 0:
        with open("build/packages/ssm/manifest.json","r") as f:
            manifest = f.read()
        response = client.create_document(
            Content=manifest,
            Attachments=[
                {
                    'Key': 'SourceUrl',
                    'Values': ['https://s3.amazonaws.com/' + s3_bucket],
                }
            ],
            Name=pkg_name,
            VersionName=rel_ver,
            DocumentType='Package'
        )
        print("%s %s is created in %s." % (pkg_name, rel_ver, region))
    else:
        response = client.list_document_versions(Name=pkg_name)
        if len([ver for ver in response['DocumentVersions'] if ver['VersionName'] == rel_ver]) == 0:
            with open("build/packages/ssm/manifest.json","r") as f:
                manifest = f.read()
            response = client.update_document(
                Content=manifest,
                Attachments=[
                    {
                        'Key': 'SourceUrl',
                        'Values': ['https://s3.amazonaws.com/' + s3_bucket],
                    }
                ],
                Name=pkg_name,
                VersionName=rel_ver,
                DocumentVersion='$LATEST'
            )
            last_version = response['DocumentDescription']['LatestVersion']
            response = client.update_document_default_version(
                Name=pkg_name,
                DocumentVersion=last_version
            )
            print("%s is updated to %s in %s." % (pkg_name, rel_ver, region))
        else:
            print("%s %s exists in %s." % (pkg_name, rel_ver, region))
