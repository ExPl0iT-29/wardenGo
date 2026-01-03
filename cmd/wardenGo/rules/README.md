# wardenGo Rules Configuration

This directory contains rule files for different security profiles.

## Rule File Format

Rules are defined in JSON format with the following structure:

```json
{
    "id": "unique-rule-id",
    "name": "Human-readable rule name",
    "conditions": {
        "file_path": "/full/path/to/file",
        "action": "write|create|remove|rename|chmod"
    },
    "severity": "low|medium|high",
    "action": "alert"
}
```

## Important: Exact Path Matching

⚠️ **The rules engine performs EXACT string matching on file paths.**

This means:
- ✅ Rule: `"file_path": "/etc/passwd"` will match exactly `/etc/passwd`
- ❌ Rule: `"file_path": "config"` will **NOT** match `/app/config/settings.json`
- ❌ Rule: `"file_path": ".env"` will **NOT** match `/app/.env`

**You must use full absolute paths in your rules.**

## Available Rule Files

### 1. `server-security.json`
**Purpose**: Linux server security monitoring  
**Use Case**: Monitor critical system directories on Linux servers  
**Example Usage**:
```bash
./wardenGo -watch-dirs="/etc,/var/log" -rules-file="rules/server-security.json"
```

### 2. `windows-server.json`
**Purpose**: Windows server security monitoring  
**Use Case**: Monitor critical system directories on Windows servers  
**Example Usage**:
```bash
./wardenGo -watch-dirs="C:\Windows\System32" -rules-file="rules/windows-server.json"
```

### 3. `application-monitoring.json`
**Purpose**: Application directory monitoring  
**Use Case**: Monitor application directories for unauthorized changes  
**Example Usage**:
```bash
./wardenGo -watch-dirs="/var/www/app" -rules-file="rules/application-monitoring.json"
```

### 4. `compliance-audit.json`
**Purpose**: Compliance and audit requirements  
**Use Case**: Track file access for regulatory compliance (GDPR, SOX, etc.)  
**Example Usage**:
```bash
./wardenGo -watch-dirs="/data/audit,/data/financial" -rules-file="rules/compliance-audit.json"
```

### 5. `web-server.json`
**Purpose**: Web server security monitoring  
**Use Case**: Monitor web server directories for web shells and unauthorized changes  
**Example Usage**:
```bash
./wardenGo -watch-dirs="/var/www/html,/etc/nginx" -rules-file="rules/web-server.json"
```

### 6. `development.json`
**Purpose**: Development environment monitoring  
**Use Case**: Monitor development directories for unusual activity  
**Example Usage**:
```bash
./wardenGo -watch-dirs="./src,./config" -rules-file="rules/development.json"
```

### 7. `sample.json`
**Purpose**: Example rule file  
**Use Case**: Template for creating custom rules

## Action Types

Currently supported actions:
- `"write"` - File write/modify operations
- `"create"` - File creation operations
- `"remove"` - File deletion operations
- `"rename"` - File rename operations
- `"chmod"` - Permission change operations

## Severity Levels

- `"low"` - Informational alerts
- `"medium"` - Moderate security concern
- `"high"` - Critical security threat

## Customizing Rules

1. Copy an existing rule file as a template
2. Modify the `file_path` values to match your actual file system paths
3. Adjust severity levels based on your security requirements
4. Save the file and wardenGo will automatically reload it

## Platform-Specific Paths

### Linux/Unix
- Use forward slashes: `/etc/passwd`
- Case-sensitive paths
- Examples: `/var/log`, `/usr/bin`, `/etc/ssh`

### Windows
- Use backslashes: `C:\Windows\System32`
- Case-insensitive (but use consistent casing)
- Examples: `C:\Windows\System32`, `C:\ProgramData`

## Tips

1. **Use absolute paths**: Always use full absolute paths in rules
2. **Test your rules**: Create test files to verify rules are working
3. **Monitor logs**: Check logrus output to see which events are detected
4. **Start simple**: Begin with a few rules and expand gradually
5. **Hot reload**: Rules are automatically reloaded when the file changes

## Example: Creating a Custom Rule

```json
[
    {
        "id": "custom-001",
        "name": "Monitor Important Config File",
        "conditions": {
            "file_path": "/home/user/important-config.json",
            "action": "write"
        },
        "severity": "high",
        "action": "alert"
    }
]
```

Save this to a file (e.g., `custom-rules.json`) and use:
```bash
./wardenGo -watch-dirs="/home/user" -rules-file="rules/custom-rules.json"
```

