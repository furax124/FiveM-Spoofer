# Purpose 
This program blocks the outbound and inbound connections from FiveM so they can't check your HWID via their authentication server. Essentially, it allows you to play FiveM on a HWID banned computer.

# How to Use:

1. **Clean Traces**:
   - Before enabling the bypass, clean any leftover traces of **FiveM** and Rockstar to avoid detection of the HWID ban. This step is done automatically by the program when you enter the paths.
   
2. **Enable Bypass**:
   - Enter the paths for your **FiveM** installation and subprocess directories when prompted. The program will automatically block FiveM and its associated processes from making outbound and inbound connections to the server.
   
3. **Log in to a New Rockstar Account**:
   - Use a new Rockstar account to log into FiveM. Ensure that the account is unused, as FiveM will ban new accounts if you try to use them without the bypass on a HWID-banned computer.

4. **Join a Server**:
   - Try joining a server. If it returns an error, disable the bypass and try again. If the bypass was not active, FiveM would have detected your HWID ban, causing the error.

5. **Important: Don't Open FiveM Without the Bypass**:
   - Never open FiveM without the bypass enabled. Doing so will result in your Rockstar account being banned for using it on a HWID-banned computer.

6. **Enable the Bypass Before Leaving the Server**:
   - Always ensure the bypass is still active when leaving the server. This is crucial because FiveM may send your HWID (and other identifying data) to the server when you exit, which could lead to your account being flagged. If needed, check that the bypass is still running before exiting.

# Notes:
- The bypass will block outgoing and incoming connections while FiveM is running. Once FiveM is closed, the bypass will automatically be removed.

# DISCLAIMER:

I Never test it and idk if it work but try it and open a issue if it doesnt work or add on discord (axel0277)
It should normally work since i saw a doc on fivem about it

# TODO

Figure out why it always disconect rockstar account after the bypass
