@echo off
pushd C:\GclAgentWin
agent.exe install
net start GclMonitor
popd