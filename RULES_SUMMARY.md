# wardenGo Rules Files - Summary

## ✅ Created Rule Files

I've created **6 rule profile files** for different use cases:

### 1. **server-security.json** (10 rules)
- Linux server security monitoring
- Monitors critical system files: `/etc/passwd`, `/etc/shadow`, SSH config, sudoers, etc.
- **Use with**: `-watch-dirs="/etc,/var/log,/usr/bin"`

### 2. **windows-server.json** (8 rules)
- Windows server security monitoring
- Monitors: System32, hosts file, registry backups, startup programs, etc.
- **Use with**: `-watch-dirs="C:\Windows\System32"`

### 3. **application-monitoring.json** (10 rules)
- Application directory monitoring
- Monitors: config files, database configs, .env files, SSL certificates, logs
- **Use with**: `-watch-dirs="/var/www/app,/app/config"`

### 4. **compliance-audit.json** (10 rules)
- Compliance and audit requirements
- Monitors: audit logs, financial data, personal data (GDPR), backups, reports
- **Use with**: `-watch-dirs="/data/audit,/data/financial"`

### 5. **web-server.json** (10 rules)
- Web server security monitoring
- Monitors: web server configs, web shells (.php, .jsp, .asp), upload directories
- **Use with**: `-watch-dirs="/var/www/html,/etc/nginx"`

### 6. **development.json** (10 rules)
- Development environment monitoring
- Monitors: source code, git config, dependencies, build artifacts, tests
- **Use with**: `-watch-dirs="./src,./config"`

## 📝 Important Notes

### ⚠️ Exact Path Matching Required
The rules engine performs **exact string matching**. You must:
- Use **full absolute paths** in rules
- Example: `"/etc/passwd"` ✅ (will match)
- Example: `"config"` ❌ (will NOT match `/app/config/file.json`)

### 📖 Documentation
See `cmd/wardenGo/rules/README.md` for:
- Detailed rule format documentation
- Usage examples for each profile
- Tips for customizing rules
- Platform-specific path examples

## 🚀 Quick Start Examples

```bash
# Linux server security
./wardenGo -watch-dirs="/etc,/var/log" -rules-file="rules/server-security.json"

# Windows server
./wardenGo -watch-dirs="C:\Windows\System32" -rules-file="rules/windows-server.json"

# Web server monitoring
./wardenGo -watch-dirs="/var/www/html" -rules-file="rules/web-server.json"

# Application monitoring
./wardenGo -watch-dirs="/app" -rules-file="rules/application-monitoring.json"
```

## 🔧 Customization

All rule files are templates. You should:
1. Copy the relevant profile file
2. Update file paths to match your actual system
3. Adjust severity levels as needed
4. Add/remove rules based on your requirements

The system will automatically reload rules when the file changes!

