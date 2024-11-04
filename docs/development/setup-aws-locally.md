# Setup Guide

## Prerequisites

1. **Docker Desktop** - Ensure Docker is installed and running.
2. **AWS CLI** - [Install AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html) to interact with AWS services locally.
3. **NoSQL Workbench for DynamoDB (Optional)**  - [Download and set up NoSQL Workbench](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/workbench.settingup.html) to manage DynamoDB tables with a graphical interface.

## Steps

1. **Start AWS Services and Redis Locally**  
   run the following command from the root directory:

   ```bash
   docker-compose -f build/package/docker-compose.yml up --remove-orphans
   ```
---
