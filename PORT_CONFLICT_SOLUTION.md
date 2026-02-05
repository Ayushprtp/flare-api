# Port 3000 Conflict - Solution Guide

## Problem Detected
A Docker container is running the old "justsong/one-api:latest" image on port 3000.
Container name: `flare`

## ✅ Good News
- Rebranding is complete!
- Binary built successfully: `./flare`
- Version shows: v0.0.0
- One last log message fixed in main.go

## 🔧 Next Steps

### Step 1: Rebuild the binary with the latest fix
```bash
go build -o flare
```

### Step 2: Choose one of these options

#### Option 1: Stop old container and run new Flare binary
```bash
docker stop flare
./flare
```
Access at: https://api.flare-sh.tech

#### Option 2: Run Flare on a different port
```bash
./flare --port 3001
```
Access at: http://localhost:3001

#### Option 3: Stop old container and rebuild with new Flare Docker image
```bash
docker stop flare
docker rm flare
docker build -t flare:latest .
docker-compose up -d
```
Access at: https://api.flare-sh.tech

#### Option 4: Keep both running (for testing)
```bash
# Old container on 3000, new binary on 3001
./flare --port 3001
```
- Old: https://api.flare-sh.tech
- New: http://localhost:3001

## Recommended: Option 1
Stop the old container and run the new binary:
```bash
docker stop flare
go build -o flare
./flare
```

This gives you the freshly rebranded Flare running natively!
