@echo off
schtasks /create /tn "gclmontor" /tr C:\falcon-agent_win\AgentStart.bat /sc onstart  /ru administrator /rp 123.com