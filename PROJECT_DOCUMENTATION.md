# wardenGo - Intrusion Detection System (IDS)

## Overview

**wardenGo** is a lightweight, real-time Intrusion Detection System (IDS) written in Go. It monitors file system activities and detects suspicious behavior based on configurable rules. The system provides continuous monitoring of directories, evaluates file system events against a rules engine, and alerts on potential security threats.

## What is wardenGo?

wardenGo is a **file system-based Intrusion Detection System** that:

- **Monitors directories** in real-time for file system events (create, write, modify, delete)
- **Evaluates events** against a customizable rules engine
- **Detects suspicious activities** such as unauthorized file access, modifications to critical system files, or unusual file operations
- **Provides real-time alerts** when rule conditions are met
- **Hot-reloads rules** without requiring system restart

## Key Features

### 1. **Real-Time File System Monitoring**
   - Continuous monitoring of specified directories
   - Detects file creation, modification, deletion, and access events
   - Low-latency event detection using Go's `fsnotify` library

### 2. **Rules-Based Detection Engine**
   - JSON-based rule configuration
   - Flexible condition matching (file paths, actions, etc.)
   - Severity levels (low, medium, high)
   - Customizable alert actions

### 3. **Hot-Reloadable Rules**
   - Rules can be updated on-the-fly without restarting the system
   - Automatic detection of rule file changes
   - Zero-downtime rule updates

### 4. **Multi-Directory Monitoring**
   - Monitor multiple directories simultaneously
   - Configurable via command-line flags
   - Independent event processing for each directory

### 5. **Structured Logging**
   - Comprehensive logging with structured fields
   - Log levels for different severity events
   - Easy integration with log aggregation systems

## Benefits

### Security Benefits
- **Early Threat Detection**: Identifies suspicious file system activities before they escalate
- **Compliance**: Helps meet security compliance requirements by monitoring critical file access
- **Forensics**: Provides detailed event logs for security incident investigation
- **Prevention**: Can be extended to block or quarantine suspicious activities

### Operational Benefits
- **Lightweight**: Minimal resource footprint, written in efficient Go
- **Easy Deployment**: Single binary, no complex dependencies
- **Flexible Configuration**: JSON-based rules make it easy to customize detection logic
- **Production-Ready**: Thread-safe implementation with proper concurrency handling

### Developer Benefits
- **Extensible Architecture**: Clean separation of concerns (scanner, engine, config)
- **Type-Safe**: Strong typing with Go's type system
- **Maintainable**: Well-structured codebase with clear module boundaries

## Selling Points

### 1. **Performance**
   - Built with Go for high performance and low latency
   - Efficient event processing with goroutines
   - Minimal memory footprint

### 2. **Reliability**
   - Thread-safe operations with proper mutex usage
   - Graceful error handling
   - Robust file watching with automatic error recovery

### 3. **Ease of Use**
   - Simple command-line interface
   - JSON-based configuration (no complex config files)
   - Clear, structured logging

### 4. **Flexibility**
   - Monitor any directory on the system
   - Customize detection rules for your specific needs
   - Extensible architecture for adding new features

### 5. **Zero-Downtime Updates**
   - Update detection rules without restarting the service
   - Continuous monitoring even during rule updates

## Use Cases

1. **Server Security Monitoring**: Monitor critical system directories (`/etc`, `/var/log`, etc.)
2. **Application Security**: Watch application directories for unauthorized changes
3. **Compliance Auditing**: Track file access for regulatory compliance
4. **Intrusion Detection**: Detect signs of system compromise through file system activity
5. **Development Environments**: Monitor project directories during development

## Architecture

The system consists of four main components:

1. **Scanner** (`internal/scanner`): Monitors file system events using `fsnotify`
2. **Rules Engine** (`internal/engine`): Evaluates events against configured rules
3. **Rule Manager** (`internal/config`): Loads and manages rule configurations
4. **Models** (`internal/models`): Defines data structures for events and rules

## Quick Start

```bash
# Build the project
go build ./cmd/wardenGo

# Run with default settings (monitors temp directory)
./wardenGo

# Monitor specific directories
./wardenGo -watch-dirs="/etc,/var/log" -rules-file="rules/sample.json"
```

## Example Rule

```json
{
    "id": "rule1",
    "name": "Unauthorized File Access",
    "conditions": {
        "file_path": "/etc/passwd",
        "action": "write"
    },
    "severity": "high",
    "action": "alert"
}
```

This rule triggers an alert when any write operation is detected on `/etc/passwd`.

## Future Enhancements

- Network activity monitoring
- Automated response actions (block, quarantine)
- Web dashboard for rule management
- Integration with SIEM systems
- Machine learning-based anomaly detection
- Performance metrics and statistics

---

**wardenGo** - Your vigilant guardian for file system security.

