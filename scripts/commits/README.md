# Commit Scripts

This directory contains all commit scripts organized in sequential order.

## Naming Convention

Scripts follow the format: `NNN_description.sh`

- `NNN`: Three-digit serial number (001, 002, 003, etc.)
- `description`: Brief description of the commit
- Each script outputs the commit identifier with hash after execution

## Usage

```bash
# Run a specific commit script
./scripts/commits/001_initial_setup.sh

# Or from project root
bash scripts/commits/002_random_users_api.sh
```

## Commit History

| #   | Script                    | Description                                    | Commit Hash |
| --- | ------------------------- | ---------------------------------------------- | ----------- |
| 001 | `001_initial_setup.sh`    | Initial project setup with core infrastructure | a710b46     |
| 002 | `002_random_users_api.sh` | Random Users API with tests and curl script    | a4f13bd     |
| 003 | `003_random_jokes_api.sh` | Random Jokes API with tests and curl script    | 9b3be54     |

## Next Commit

The next commit script should be numbered `004_*.sh`
