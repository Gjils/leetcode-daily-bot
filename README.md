#### Bot that send leetcode daily problems to telegram chats ([@leetcode_daily_notifier_bot](https://t.me/leetcode_daily_notifier_bot))

##### Run locally
1) Clone the  repo
``` bash
git clone https://github.com/Gjils/leetcode-daily-bot
```
2) Create directory `leetcodebotdb` next to cloned repo
``` bash
mkdir leetcodebotdb
```
3) Run bot, replace YOUR_BOT_TOKEN with your token
```bash
cd leetcode-daily-bot
BOT_TOKEN="YOUR_BOT_TOKEN" docker compose up -d --build
```
