# OwlHook - Alertmanager Webhook Adapter

OwlHook is a webhook adapter that receives alerts from Alertmanager and forwards them to various chat platforms including Lark (Feishu), DingTalk, and WeCom (Enterprise WeChat).

## Features

- Supports multiple chat platforms:
  - Lark (Feishu)
  - DingTalk
  - WeCom (Enterprise WeChat)
- Configurable through JSON configuration file
- Easy to extend for additional platforms
- Markdown message formatting
- Signature support for DingTalk

## Configuration

set up your webhook url in the `config.yaml` file:

```yaml
webhook:
  lark:
    webhookUrl: "https://open.larksuite.com/open-apis/bot/v2/hook/ba1856f8-bm4a-45bc-a191-1ad7dbd5f3c4"
    enabled: true
  dingtalk:
    webhookUrl: "https://oapi.dingtalk.com/robot/send?access_token=your-access-token"
    enabled: false
  wecom:
    webhookUrl: "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=your-webhook-key"
    enabled: false
```

## Usage

1. Build the program:
   ```bash
   go build
   ```

2. Run the program:
   ```bash
   ./OwlHook -config config.json -port 8080
   ```

3. Configure Alertmanager to send webhooks to your endpoint:
   ```yaml
   receivers:
   - name: 'lark-notifications'
     webhook_configs:
     - url: 'http://localhost:8000/webhook/lark'
   ```

## API

The webhook endpoint accepts POST requests at `/webhook/:platform` with the following path parameters:

- Path parameter: `platform`

Supported platform names:
- `lark`
- `dingtalk`
- `wecom`

## Message Format

The program formats alert messages in a readable way, including:
- Alert status (firing/resolved)
- Summary and description
- Number of alerts
- Individual alert details
- Start and end times
- Duration (for resolved alerts)
- Link to Alertmanager

## Development

To add support for a new chat platform:

1. Create a new notifier in the `notifiers` package
2. Implement the `NotifierInterface`
3. Add configuration support in the `Config` struct

## License

MIT License