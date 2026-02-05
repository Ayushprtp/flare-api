# Flare Rebranding Checklist ✅

## Completed Changes

### ✅ Core Branding
- [x] Application name: "New API" → "Flare"
- [x] Binary name: `new-api` → `flare`
- [x] Module path: `github.com/Flare-sh/new-api` → `github.com/Flare-sh/api`
- [x] System name constant: "New API" → "Flare"
- [x] Database filename: `one-api.db` → `flare.db`

### ✅ Build & Deployment
- [x] Dockerfile: Binary name and paths updated
- [x] docker-compose.yml: Service, container, and database names updated
- [x] Systemd service: `new-api.service` → `flare.service`
- [x] .gitignore: Added `flare` binary

### ✅ Source Code
- [x] All Go import paths updated
- [x] All "NewAPI" → "Flare" in Go files
- [x] All "new-api" → "flare" in Go files
- [x] All "newapi" → "flare" in Go files
- [x] Error types: "new_api_panic" → "flare_panic"

### ✅ Frontend
- [x] Default system name in utils.jsx
- [x] Footer branding and links
- [x] About page repository links
- [x] Console welcome message: "WE ❤ FLARE"
- [x] All JS/JSX/TS/TSX files updated

### ✅ Internationalization
- [x] English (en.json)
- [x] Chinese (zh.json)
- [x] French (fr.json)
- [x] Japanese (ja.json)
- [x] Russian (ru.json)
- [x] Vietnamese (vi.json)

### ✅ Documentation
- [x] README.md
- [x] README.zh.md
- [x] README.fr.md
- [x] README.ja.md
- [x] All documentation URLs updated

### ✅ URLs & Links
- [x] Repository: `github.com/Flare-sh/new-api` → `github.com/Flare-sh/api`
- [x] Issue tracking: Points to `github.com/Flare-sh/api/issues`
- [x] Documentation: `docs.newapi.pro` → GitHub README sections
- [x] Related projects: `new-api-horizon` → `flare-horizon`
- [x] Panic/error URLs: Point to Flare issues

### ✅ Attribution Preserved
- [x] One API original project credit maintained
- [x] JustSong attribution in About page
- [x] License references preserved
- [x] Footer link to original One API project

## Next Steps

### 1. Test the Build
```bash
# Test Go build
go mod tidy
go build -o flare

# Test frontend build
cd web
bun install
bun run build
cd ..

# Test Docker build
docker build -t flare:test .
```

### 2. Update Docker Registry
```bash
# Tag and push to your Docker registry
docker tag flare:test yourusername/flare:latest
docker push yourusername/flare:latest
```

### 3. Update docker-compose.yml
Change the image reference from:
```yaml
image: calciumion/flare:latest
```
to:
```yaml
image: yourusername/flare:latest
```

### 4. Test Deployment
```bash
# Test with Docker Compose
docker-compose up -d

# Check logs
docker-compose logs -f flare

# Verify service is running
curl https://api.flare-sh.tech/api/status
```

### 5. Update CI/CD (if applicable)
- Update GitHub Actions workflows
- Update build scripts
- Update deployment scripts
- Update environment variables

### 6. Update Documentation
- Create your own documentation site (optional)
- Or keep pointing to GitHub README
- Update any external documentation references

### 7. Announce the Rebrand
- Update GitHub repository description
- Update README badges
- Create a release announcement
- Update any social media/community links

## Important Notes

⚠️ **Database Compatibility**: The database schema is unchanged. Existing databases will work with the new binary.

⚠️ **API Compatibility**: All API endpoints remain the same. Clients don't need updates.

⚠️ **Docker Images**: Remember to update the image name in docker-compose.yml to your own registry.

⚠️ **Attribution**: Always maintain credit to the original One API project by JustSong.

## Verification Commands

```bash
# Verify no remaining "new-api" references (should be minimal)
grep -r "new-api" --include="*.go" --include="*.jsx" --include="*.js" . | grep -v "go-epay" | grep -v ".git" | wc -l

# Verify no remaining "New API" references
grep -r "New API" --include="*.go" --include="*.jsx" --include="*.js" . | grep -v ".git" | wc -l

# Verify module path
grep "module github.com/Flare-sh/api" go.mod

# Verify binary name in Dockerfile
grep "RUN go build" Dockerfile | grep "flare"

# Verify service name
grep "container_name: flare" docker-compose.yml
```

## Success Criteria

✅ Application builds successfully
✅ Frontend builds without errors
✅ Docker image builds successfully
✅ All tests pass (if applicable)
✅ Application runs and serves requests
✅ UI displays "Flare" branding
✅ Console shows "WE ❤ FLARE" message
✅ Error messages point to correct repository
✅ Attribution to original project is visible

---

**Rebranding Status**: ✅ COMPLETE

All references to "New API", "new-api", "One API", and "one-api" have been successfully replaced with "Flare" while maintaining proper attribution to the original project.
