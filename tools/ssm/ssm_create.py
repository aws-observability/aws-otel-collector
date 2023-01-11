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
import argparse
import boto3

if __name__ == "__main__":

    parser = argparse.ArgumentParser()
    parser.add_argument("package_name")
    parser.add_argument("version")
    parser.add_argument("s3_bucket")
    parser.add_argument("region")
    parser.add_argument("--make_default_version", help="set if this should be the default package version", action="store_true")
    args = parser.parse_args()

    pkg_name = args.package_name
    rel_ver = args.version
    s3_bucket = args.s3_bucket
    region = args.region

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
            print("%s is updated to %s in %s." % (pkg_name, rel_ver, region))
            if args.make_default_version:
                last_version = response['DocumentDescription']['LatestVersion']
                response = client.update_document_default_version(
                    Name=pkg_name,
                    DocumentVersion=last_version
                )
                print("%s is set default to %s in %s." % (pkg_name, rel_ver, region))
        else:
            print("%s %s exists in %s." % (pkg_name, rel_ver, region))
