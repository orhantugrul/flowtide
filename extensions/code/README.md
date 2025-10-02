# flowtide VS Code Extension

A lightweight coding activity tracker for VS Code and compatible editors. Track your coding time and analyze your development patterns.

## Quick Start

1. Install the extension from the VS Code marketplace
2. Configure your Flowtide server URL in settings (default: `http://localhost:8080`)
3. Click the Flowtide status bar item or use Command Palette: `Flowtide: Toggle Activity Tracking`

## Configuration

### Settings

| Setting                      | Type    | Default                 | Description                           |
| ---------------------------- | ------- | ----------------------- | ------------------------------------- |
| `flowtide.url`               | string  | `http://localhost:8080` | Flowtide dashboard URL                |
| `flowtide.enabled`           | boolean | `true`                  | Enable/disable tracking               |
| `flowtide.autoStart`         | boolean | `true`                  | Auto-start tracking on VS Code launch |
| `flowtide.showNotifications` | boolean | `true`                  | Show tracking notifications           |

### Commands

- `Flowtide: Start Activity Tracking` - Begin tracking
- `Flowtide: Stop Activity Tracking` - Stop tracking
- `Flowtide: Toggle Activity Tracking` - Toggle tracking state
- `Flowtide: Open Dashboard` - Open web dashboard
- `Flowtide: Show Logs` - View extension logs
- `Flowtide: Test Connection` - Test server connectivity

## How It Works

The extension monitors your coding activity by tracking:

1. **File Changes**: When you switch between files
2. **Text Edits**: When you make changes to files
3. **Language Detection**: Automatically detects the programming language
4. **Project Context**: Associates activity with your current workspace

Data is collected locally and sent to your Flowtide server in batches to minimize network overhead.

## Development

### Prerequisites

- Node.js 22+
- TypeScript 5.9+

### Setup

```bash
# Install dependencies
npm install

# Compile TypeScript
npm run compile

# Watch for changes
npm run watch

# Run tests
npm test
```

### Building

```bash
# Build for production
npm run vscode:prepublish

# Package extension
vsce package
```

## API Integration

The extension integrates with your Flowtide backend through REST APIs:

- `POST /api/editors` - Create/register editor
- `POST /api/projects` - Create/register project
- `POST /api/activities` - Create single activity
- `POST /api/activities/batch` - Create multiple activities
- `GET /api/health` - Health check

## Troubleshooting

### Common Issues

1. **Connection Failed**: Check your server URL and ensure the Flowtide server is running
2. **Tracking Not Starting**: Verify the extension is enabled in settings
3. **No Data in Dashboard**: Check the extension logs for any errors

### Debugging

1. Open Command Palette and run `Flowtide: Show Logs`
2. Check the output panel for detailed logging information
3. Use `Flowtide: Test Connection` to verify server connectivity

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

MIT License - see LICENSE file for details
