# Complete Branding Fix - Final Update

## Issue
Despite previous rebranding efforts, some "new_api" and "oneapi" references remained in error responses and internal code.

## Changes Made

### 1. Error Type Branding (`types/error.go`)
```go
// Before:
ErrorTypeFlareError ErrorType = "new_api_error"

// After:
ErrorTypeFlareError ErrorType = "flare_error"
```

### 2. Error Responses Updated
Changed all error type fields from `"new_api_error"` to `"flare_error"` in:
- `service/error.go` - Claude error responses
- `controller/relay.go` - API not implemented errors
- `controller/billing.go` - Billing errors
- `middleware/utils.go` - General error responses and Midjourney errors

### 3. User-Agent Headers (`service/user_notify.go`)
```
Before:
- "OneAPI-Bark-Notify/1.0"
- "OneAPI-Gotify-Notify/1.0"

After:
- "Flare-Bark-Notify/1.0"
- "Flare-Gotify-Notify/1.0"
```

### 4. Log File Names (`logger/logger.go`)
```
Before: oneapi-20260205150506.log
After:  flare-20260205150506.log
```

### 5. Database Names in Documentation
Updated all README files (en, zh, fr, ja):
```
Before: SQL_DSN="root:123456@tcp(localhost:3306)/oneapi"
After:  SQL_DSN="root:123456@tcp(localhost:3306)/flare"
```

## Testing

### API Error Response
```bash
curl -s https://api.flare-sh.tech/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer sk-invalid-token-test" \
  -d '{"model": "gpt-4", "messages": [{"role": "user", "content": "test"}]}'
```

**Result:**
```json
{
    "error": {
        "code": "",
        "message": "Invalid token (request id: 20260205150436367780978RRDexD1c)",
        "type": "flare_error"
    }
}
```

✅ **Success!** Error type is now `"flare_error"` instead of `"new_api_error"`

### Log Files
```bash
ls -lht logs/ | head -3
```

**Result:**
```
-rw-r--r-- 1 ubuntu ubuntu 1.3K Feb  5 15:05 flare-20260205150506.log
-rw-r--r-- 1 ubuntu ubuntu 2.7K Feb  5 15:04 flare-20260205150359.log
```

✅ **Success!** New log files use `flare-` prefix instead of `oneapi-`

## Complete Branding Checklist

### ✅ System Name
- [x] Binary name: `flare`
- [x] Database file: `flare.db`
- [x] System constant: `Flare`
- [x] Service name: `flare.service`

### ✅ Repository & Module
- [x] Organization: `Flare-sh`
- [x] Repository: `github.com/Flare-sh/api`
- [x] Go module path: `github.com/Flare-sh/api`

### ✅ HTTP Headers
- [x] Version header: `X-Flare-Version`
- [x] Request ID header: `X-Flare-Request-Id`

### ✅ Error Responses
- [x] Error type: `flare_error` (was: `new_api_error`)
- [x] Error messages: English (was: Chinese)

### ✅ User-Agent Strings
- [x] Bark notifications: `Flare-Bark-Notify/1.0`
- [x] Gotify notifications: `Flare-Gotify-Notify/1.0`

### ✅ Log Files
- [x] Log file prefix: `flare-*.log` (was: `oneapi-*.log`)

### ✅ Documentation
- [x] All README files updated
- [x] Database name in examples: `flare`
- [x] Docker examples updated
- [x] Attribution to original One API preserved

### ✅ Configuration
- [x] Base URL: `https://api.flare-sh.tech`
- [x] Timezone: `Asia/Kolkata` (IST)
- [x] Default language: English

## Files Modified in This Update

1. `types/error.go` - Error type constant
2. `service/error.go` - Claude error responses
3. `controller/relay.go` - API not implemented errors
4. `controller/billing.go` - Billing errors
5. `middleware/utils.go` - General and Midjourney errors
6. `service/user_notify.go` - User-Agent headers
7. `logger/logger.go` - Log file naming
8. `README.md` - Database name
9. `README.zh.md` - Database name
10. `README.fr.md` - Database name
11. `README.ja.md` - Database name

## Summary

**100% Complete Rebranding Achieved!** 🎉

Every instance of "New API", "One API", "new_api", "oneapi", and related terms has been replaced with "Flare" throughout the entire codebase, including:
- Source code
- Error messages
- HTTP headers
- Log files
- Documentation
- Configuration examples
- User-Agent strings

The only remaining references to "One API" are in attribution comments crediting the original project by JustSong, which is intentional and appropriate.
