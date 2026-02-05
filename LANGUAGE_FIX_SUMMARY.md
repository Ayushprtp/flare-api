# Language Configuration Fix Summary

## Issue
Error messages were appearing in Chinese instead of English when making API requests.

## Root Causes
1. **i18n Bundle Initialization**: The i18n bundle was initialized with `language.Chinese` as the default
2. **Hardcoded Error Messages**: Token validation errors in `model/token.go` were hardcoded in Chinese

## Changes Made

### 1. i18n Configuration (`i18n/i18n.go`)
Changed the default language bundle from Chinese to English:
```go
// Before:
bundle = i18n.NewBundle(language.Chinese)

// After:
bundle = i18n.NewBundle(language.English)
```

### 2. Token Validation Errors (`model/token.go`)
Replaced all hardcoded Chinese error messages with English:

| Chinese Message | English Message |
|----------------|-----------------|
| 未提供令牌 | Token not provided |
| 该令牌额度已用尽 | Token quota exhausted |
| 该令牌已过期 | This token has expired |
| 该令牌状态不可用 | This token status is unavailable |
| 无效的令牌 | Invalid token |
| 无效的令牌，数据库查询出错 | Invalid token, database query error |
| id 或 userId 为空！ | id or userId is empty! |
| id 为空！ | id is empty! |
| quota 不能为负数！ | quota cannot be negative! |
| ids 不能为空！ | ids cannot be empty! |

## Testing
```bash
# Test with invalid token
curl -s https://api.flare-sh.tech/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer sk-invalid-token-test" \
  -d '{"model": "gpt-4", "messages": [{"role": "user", "content": "test"}]}'

# Result: {"error":{"code":"","message":"Invalid token (request id: ..."}}
```

## Running Flare in Background

### Option 1: systemd (Recommended - Already Configured)
```bash
# Check status
sudo systemctl status flare

# Start service
sudo systemctl start flare

# Stop service
sudo systemctl stop flare

# Restart service
sudo systemctl restart flare

# Enable auto-start on boot
sudo systemctl enable flare

# View logs
sudo journalctl -u flare -f
```

### Option 2: nohup (Alternative)
```bash
# Start in background
nohup ./flare --port 3000 --log-dir ./logs > flare.out 2>&1 &

# Check if running
ps aux | grep flare

# Stop
pkill -f "./flare"
```

### Option 3: screen/tmux (For development)
```bash
# Using screen
screen -S flare
./flare --port 3000
# Press Ctrl+A then D to detach

# Reattach
screen -r flare

# Using tmux
tmux new -s flare
./flare --port 3000
# Press Ctrl+B then D to detach

# Reattach
tmux attach -t flare
```

## Current Status
✅ Flare is running via systemd on port 3000
✅ All error messages now display in English
✅ Service auto-starts on system boot
✅ Logs are written to `/home/ubuntu/new-api/logs/`

## Service Configuration
Location: `/etc/systemd/system/flare.service`
- User: ubuntu
- Working Directory: `/home/ubuntu/new-api`
- Port: 3000
- Log Directory: `/home/ubuntu/new-api/logs`
- Auto-restart: Enabled

## Notes
- The i18n system still supports both English and Chinese
- Users can set their preferred language in their account settings
- The Accept-Language header is respected for API requests
- Default language is now English for all new requests
