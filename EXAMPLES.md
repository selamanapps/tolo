# Tolo Examples

This file contains practical examples of using Tolo in various scenarios.

## SSH Connections

### Simple SSH
```bash
# Save
tolo s myserver:ssh user@192.168.1.10

# Run
tolo r myserver
```

### SSH with Custom Port
```bash
tolo s myserver:ssh user@192.168.1.10 -p 2222
```

### SSH with Key File
```bash
tolo s myserver:ssh user@192.168.1.10 -i ~/.ssh/mykey.pem
```

### SSH with Multiple Options
```bash
tolo s production:ssh -i ~/.ssh/prod-key -p 2222 admin@prod.example.com
```

## Cloud Services

### Google Cloud
```bash
# Save gcloud SSH connection
tolo s gcloud-ssh:gcloud compute ssh instance-1 --zone us-central1-a --project my-project

# Save gcloud deployment
tolo s gcloud-deploy:gcloud app deploy
```

### AWS
```bash
# Save AWS SSM
tolo s aws-ssm:aws ssm start-session --target i-0123456789abcdef0

# Save AWS deployment
tolo s aws-deploy:aws lambda update-function-code --function-name my-function
```

### Azure
```bash
# Save Azure SSH
tolo s azure-vm:az vm ssh --resource-group myRG --name myVM
```

## Docker

### Start Containers
```bash
tolo s dev:docker-compose up -d
tolo s prod:docker-compose -f docker-compose.prod.yml up -d
```

### Build and Run
```bash
tolo s rebuild:docker-compose up -d --build
```

### Docker Commands
```bash
tolo s logs:docker-compose logs -f
tolo s stop:docker-compose down
```

## Development

### Git Commands
```bash
# Save complex git command
tolo s git-stats:git shortlog -sn --all

# Save git push to origin
tolo s gp:git push origin $(git branch --show-current)
```

### Build Commands
```bash
# Save Go build
tolo s go-build:go build -ldflags="-s -w" -o myapp

# Save Node build
tolo s npm-build:npm run build && npm run test
```

### Testing
```bash
# Save test command
tolo s test-all:go test ./... -race -cover
```

## System Administration

### System Update
```bash
tolo s update:sudo apt update && sudo apt upgrade -y
```

### Backup
```bash
tolo s backup:rsync -avz /home/user/ /backup/user/
```

### Log Monitoring
```bash
tolo s logs:tail -f /var/log/syslog
```

### System Check
```bash
tolo s sysinfo:htop
```

## Database

### PostgreSQL
```bash
# Save psql connection
tolo s psql-prod:psql -h prod-db.example.com -U admin -d production

# Save backup
tolo s pg-backup:pg_dump -h localhost -U admin mydb > backup.sql
```

### MongoDB
```bash
# Save mongo connection
tolo s mongo:mongo mongodb://localhost:27017/mydb

# Save backup
tolo s mongo-backup:mongodump --host localhost --db mydb --out /backup
```

### Redis
```bash
# Save redis connection
tolo s redis-cli:redis-cli -h localhost -p 6379
```

## File Operations

### Transfer Files
```bash
tolo s sync:rsync -avz --progress source/ user@remote:/destination/
```

### Archive
```bash
tolo s backup-tar:tar -czf backup-$(date +%Y%m%d).tar.gz /important/files/
```

## Automation

### Cron Jobs
```bash
# Add cron job
tolo s add-cron:crontab -e

# List cron jobs
tolo s list-cron:crontab -l
```

### Batch Processing
```bash
# Process files
tolo s process:for f in *.txt; do convert "$f" "${f%.txt}.pdf"; done
```

## Kubernetes

### Kubectl Commands
```bash
# Save kubectl get pods
tolo s k-pods:kubectl get pods -A

# Save deployment command
tolo s k-deploy:kubectl apply -f deployment.yaml
```

## Monitoring

### Health Check
```bash
tolo s health:curl -f http://localhost:8080/health || exit 1
```

### Server Status
```bash
tolo s status:systemctl status myservice
```

## Tips and Tricks

### Use Meaningful Names
```bash
# Good
tolo s prod-db-ssh:ssh admin@prod-db.example.com

# Bad
tolo s server1:ssh admin@192.168.1.10
```

### Update Aliases
```bash
# When connection details change
tolo u prod-db-ssh:ssh admin@new-prod-db.example.com
```

### Search Before Creating
```bash
# Check if alias exists
tolo se db

# Show existing alias details
tolo sh prod-db-ssh
```

### Use Shell Variables in Commands
```bash
# Save with variable
tolo s test:echo $HOME

# When you run it, it will use your current $HOME
```

## Migration from Shell Aliases

If you have existing shell aliases, migrate them to Tolo:

```bash
# In .bashrc:
alias server1='ssh user@192.168.1.10'

# Migrate to Tolo:
tolo s server1:ssh user@192.168.1.10

# Now remove from .bashrc and use Tolo instead
```

## Backup Your Aliases

```bash
# Tolo stores aliases in ~/.tolo/tolo.db.json
# Simply copy this file to backup:

cp ~/.tolo/tolo.db.json ~/backup/tolo-backup.json
```

## Restore Aliases

```bash
# Restore from backup
cp ~/backup/tolo-backup.json ~/.tolo/tolo.db.json
```
