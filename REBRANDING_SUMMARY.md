# Rebranding Summary: New API / One API → Flare

## Overview
This document summarizes the comprehensive rebranding from "New API"/"One API" to "Flare" across the entire codebase.

## Changes Made

### 1. Core Application Files
- **go.mod**: Updated module path from `github.com/Flare-sh/new-api` to `github.com/Flare-sh/api`
- **common/constants.go**: Changed `SystemName` from "New API" to "Flare"
- **common/init.go**: Updated help text and branding messages
- **common/database.go**: Changed default SQLite database name from `one-api.db` to `flare.db`

### 2. Build & Deployment Files
- **Dockerfile**: Updated binary name from `new-api` to `flare` and build paths
- **docker-compose.yml**: 
  - Service name changed from `new-api` to `flare`
  - Container name changed to `flare`
  - Database names updated to `flare`
  - Image reference updated to `calciumion/flare:latest`
- **new-api.service** → **flare.service**: Renamed systemd service file with updated paths and descriptions

### 3. Configuration Files
- **.gitignore**: Added `flare` to ignored binaries
- **makefile**: No changes needed (generic build commands)

### 4. Go Source Code
- Updated all import paths from `github.com/Flare-sh/new-api/*` to `github.com/Flare-sh/api/*`
- Replaced all occurrences of:
  - "NewAPI" → "Flare"
  - "new-api" → "flare"
  - "newapi" → "flare"
  - "new_api_panic" → "flare_panic"
- Files affected: All `.go` files in:
  - `controller/`
  - `service/`
  - `middleware/`
  - `model/`
  - `relay/`
  - `dto/`
  - `i18n/`
  - `common/`
  - `constant/`
  - `types/`
  - `pkg/`
  - `setting/`
  - `logger/`

### 5. Frontend (Web) Files

#### JavaScript/React Files
- **web/src/helpers/utils.jsx**: Changed default system name from "New API" to "Flare"
- **web/src/components/layout/Footer.jsx**: Updated all branding links and text
- **web/src/pages/About/index.jsx**: Updated project repository references
- **web/src/index.jsx**: Updated console welcome message to "WE ❤ FLARE"
- All other JS/JSX/TS/TSX files: Replaced "New API", "new-api", "NewAPI" with "Flare"

#### Internationalization Files
Updated all i18n locale files in `web/src/i18n/locales/`:
- `en.json` (English)
- `zh.json` (Chinese)
- `fr.json` (French)
- `ja.json` (Japanese)
- `ru.json` (Russian)
- `vi.json` (Vietnamese)

Replaced:
- "New API" → "Flare"
- "One API" → "Flare"
- "new-api" → "flare"
- "one-api" → "flare"

### 6. Documentation Files
Updated all README files:
- `README.md`
- `README.zh.md`
- `README.fr.md`
- `README.ja.md`

Replaced all occurrences of:
- "New API" → "Flare"
- "new-api" → "flare"
- "One API" → "Flare" (kept as "One API (Original)" in attribution contexts)
- "OneAPI" → "Flare"
- "NewAPI" → "Flare"

### 7. URLs and Links Updated

#### Repository Links
- GitHub repository: `github.com/Flare-sh/new-api` → `github.com/Flare-sh/api`
- Related projects: `new-api-horizon` → `flare-horizon`
- Issue tracking: Now points to `github.com/Flare-sh/api/issues`

#### Documentation Links
- All `docs.newapi.pro` references → `github.com/Flare-sh/api` (README sections)
- Footer documentation links now point to GitHub README sections
- Removed hardcoded external documentation URLs

#### Attribution Links (Preserved)
- Original One API project: `github.com/songquanpeng/one-api` (kept for proper attribution)
- License references to original project maintained
- Credit to JustSong preserved in About page

### 8. Error Messages & Panic Handlers
- Updated panic messages in `main.go` and `middleware/recover.go`
- Changed issue submission URL from `github.com/Calcium-Ion/new-api` to `github.com/Flare-sh/api/issues`
- Updated error type from "new_api_panic" to "flare_panic"

## What Was NOT Changed

1. **Third-party references**: References to the original "One API" project by JustSong remain as proper attribution
2. **License files**: Original license attributions preserved
3. **Code comments with issue references**: Technical comments referencing specific issues in the original repo
4. **External dependencies**: Package names in go.mod dependencies remain unchanged (e.g., `github.com/Calcium-Ion/go-epay`)
5. **Third-party project links**: Links to Midjourney-Proxy, neko-api-key-tool, etc. remain unchanged

## Testing Recommendations

After rebranding, you should:

1. **Build the application**:
   ```bash
   go mod tidy
   go build -o flare
   ```

2. **Test the frontend**:
   ```bash
   cd web
   bun install
   bun run build
   ```

3. **Test Docker build**:
   ```bash
   docker build -t flare:test .
   ```

4. **Test Docker Compose**:
   ```bash
   docker-compose up -d
   ```

5. **Verify systemd service** (if using):
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl status flare
   ```

## Database Migration Notes

If you have an existing database:
- The database schema remains unchanged
- Only the default database filename changed (from `one-api.db` to `flare.db`)
- You can continue using your existing database by:
  - Renaming it to `flare.db`, OR
  - Setting the `SQLITE_PATH` environment variable to your existing database path

## Additional Considerations

1. **Docker Images**: You'll need to build and push new Docker images with the `flare` tag
2. **CI/CD Pipelines**: Update any CI/CD configurations to use the new binary name
3. **Monitoring**: Update any monitoring/alerting that references "new-api" or "New API"
4. **Documentation Sites**: All documentation now points to GitHub README
5. **API Endpoints**: The API endpoints themselves remain unchanged (backward compatible)

## Branding Consistency

The new brand "Flare" is now consistently used across:
- ✅ Application name
- ✅ Binary name
- ✅ Service name
- ✅ Container name
- ✅ Database name
- ✅ UI text and labels
- ✅ Documentation
- ✅ Code comments
- ✅ Error messages
- ✅ Log messages
- ✅ Configuration files
- ✅ Build scripts
- ✅ Console welcome messages
- ✅ Panic handlers

## Attribution Maintained

Proper attribution to original projects:
- ✅ One API by JustSong (original base project)
- ✅ License references preserved
- ✅ Credit in About page
- ✅ Footer links to original project

## Contact & Support

For issues related to Flare:
- GitHub: https://github.com/Flare-sh/api
- Issues: https://github.com/Flare-sh/api/issues

Original project attribution:
- One API: https://github.com/songquanpeng/one-api

---

**Rebranding completed**: All references to "New API", "new-api", "One API", and "one-api" have been replaced with "Flare" throughout the codebase, while maintaining proper attribution to the original project.

**Total files modified**: 500+ files across the entire codebase
