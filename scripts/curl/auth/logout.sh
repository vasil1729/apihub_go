#!/bin/bash
echo "Logging out..."
curl -v -X POST http://localhost:8080/api/v1/auth/logout
