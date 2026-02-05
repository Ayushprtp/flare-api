# Site Configuration Update

## ✅ Changes Applied

### 1. Base URL Updated
**FROM:** `http://localhost:3000`  
**TO:** `https://api.flare-sh.tech`

**Files Updated:**
- `setting/system_setting/system_setting_old.go` - ServerAddress
- All README files (en, zh, fr, ja)
- All documentation markdown files
- QUICK_START.md
- PORT_CONFLICT_SOLUTION.md
- REBRANDING_CHECKLIST.md

### 2. Timezone Updated to IST
**FROM:** `TZ=Asia/Shanghai` (China Standard Time)  
**TO:** `TZ=Asia/Kolkata` (Indian Standard Time)

**Files Updated:**
- `docker-compose.yml`
- All README files
- QUICK_START.md
- All deployment documentation

### 3. Footer Link Verified
**Correct URL:** `https://github.com/Flare-sh/api`

All footer links are correctly pointing to the Flare-sh organization.

## 📋 Configuration Summary

### Production Settings
```yaml
Site URL: https://api.flare-sh.tech
Timezone: Asia/Kolkata (IST - UTC+5:30)
Repository: https://github.com/Flare-sh/api
Organization: Flare-sh
```

### Environment Variables
```bash
# Server Address (for system settings)
ServerAddress=https://api.flare-sh.tech

# Timezone (for Docker and system)
TZ=Asia/Kolkata

# Database (example)
SQL_DSN=postgresql://user:pass@host:5432/flare

# Redis (example)
REDIS_CONN_STRING=redis://redis:6379
```

## 🔧 Deployment Notes

### Docker Compose
The `docker-compose.yml` now uses:
```yaml
environment:
  - TZ=Asia/Kolkata
```

### README Files
All README files now reference:
- Site: `https://api.flare-sh.tech`
- Timezone: Asia/Kolkata (IST)

### System Settings
The default server address is now set to `https://api.flare-sh.tech` in the system settings.

## ✓ Verification

After rebuilding and deploying:

1. **Check timezone:**
   ```bash
   docker exec flare date
   # Should show IST time
   ```

2. **Check server address:**
   ```bash
   # In admin panel, check System Settings
   # ServerAddress should be https://api.flare-sh.tech
   ```

3. **Check footer links:**
   - Visit https://api.flare-sh.tech
   - Scroll to footer
   - Verify links point to github.com/Flare-sh/api

## 🚀 Next Steps

1. Rebuild the application:
   ```bash
   go build -o flare
   ```

2. Update Docker image (if using Docker):
   ```bash
   docker-compose down
   docker-compose build
   docker-compose up -d
   ```

3. Verify the changes:
   - Visit https://api.flare-sh.tech
   - Check footer links
   - Verify timezone in logs

All configuration now points to your production API at https://api.flare-sh.tech with IST timezone! 🎉
