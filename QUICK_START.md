# 🚀 Flare - Quick Start Guide

## ✅ Rebranding Complete!

Your project has been successfully rebranded from "New API" to **Flare**.

## 📋 What Changed

- **Application Name**: New API → Flare
- **Binary Name**: `new-api` → `flare`
- **Repository**: `github.com/Flare-sh/new-api` → `github.com/Flare-sh/api`
- **Database**: `one-api.db` → `flare.db`
- **All URLs**: Updated to point to your Flare repository
- **Attribution**: Original One API project credit preserved

## 🔨 Build & Test

### 1. Build the Backend
```bash
# From project root
go mod tidy
go build -o flare
```

### 2. Build the Frontend
```bash
cd web
bun install
bun run build
cd ..
```

### 3. Test Locally
```bash
# Run the binary
./flare --port 3000

# Or with Docker
docker build -t flare:latest .
docker run -p 3000:3000 flare:latest
```

### 4. Deploy with Docker Compose
```bash
# Update the image name in docker-compose.yml first!
# Change: image: calciumion/flare:latest
# To:     image: yourusername/flare:latest

docker-compose up -d
```

## 🐳 Docker Registry

Before deploying, you need to:

1. **Build your image**:
   ```bash
   docker build -t yourusername/flare:latest .
   ```

2. **Push to registry**:
   ```bash
   docker push yourusername/flare:latest
   ```

3. **Update docker-compose.yml**:
   ```yaml
   services:
     flare:
       image: yourusername/flare:latest  # Change this line
   ```

## 📁 Important Files

- `flare.service` - Systemd service file
- `docker-compose.yml` - Docker Compose configuration
- `Dockerfile` - Docker build configuration
- `REBRANDING_SUMMARY.md` - Detailed change log
- `REBRANDING_CHECKLIST.md` - Complete checklist

## 🔍 Verify Installation

```bash
# Check binary
./flare --version

# Check service (if using systemd)
sudo systemctl status flare

# Check Docker container
docker ps | grep flare

# Test API
curl https://api.flare-sh.tech/api/status
```

## 📝 Database Migration

If you have an existing database:

**Option 1**: Rename your database
```bash
mv one-api.db flare.db
```

**Option 2**: Set environment variable
```bash
export SQLITE_PATH="your-existing-database.db"
./flare
```

## 🌐 Access the Application

After starting:
- **Web UI**: https://api.flare-sh.tech
- **API**: https://api.flare-sh.tech/api
- **Default credentials**: Check your configuration

## ⚙️ Environment Variables

Key environment variables (see docker-compose.yml for full list):

```bash
SQL_DSN=postgresql://user:pass@host:5432/flare
REDIS_CONN_STRING=redis://redis:6379
SESSION_SECRET=your-random-secret
TZ=Asia/Kolkata
```

## 🆘 Troubleshooting

### Build fails with import errors
```bash
go mod tidy
go clean -modcache
go build -o flare
```

### Frontend build fails
```bash
cd web
rm -rf node_modules bun.lock
bun install
bun run build
```

### Docker image not found
Update the image name in `docker-compose.yml` to your own registry.

### Database connection fails
Check your `SQL_DSN` environment variable and ensure the database exists.

## 📚 Documentation

- **GitHub**: https://github.com/Flare-sh/api
- **Issues**: https://github.com/Flare-sh/api/issues
- **README**: See README.md for full documentation

## 🙏 Attribution

This project is based on [One API](https://github.com/songquanpeng/one-api) by JustSong.

## 📄 License

See LICENSE file for details.

---

**Ready to go!** 🎉

Start with: `go build -o flare && ./flare`
