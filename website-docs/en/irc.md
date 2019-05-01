---
title: "Connecting to Atoka through IRC"
old_id: 11
---
## Atoka has now IRC support!
Did you know that with the latest version of our **pep.py Bancho server**, you can connect to Atoka's chat and talk with other users from outside the game using an IRC client?  
It's super easy! Follow the instructions below to know how to do it.

### Before starting...
Please note that IRC support is still **work in progress** and might not work with **every IRC client** or might have some **bugs**.  
Our IRC server only handles **basic commands** at the moment, such as PASS, NICK, USER, PRIVMSG, NOTICE, JOIN, PART and QUIT.  
Support for more commands will come later.  
**Also, keep in mind that, for now, you can't be logged in to Atoka through Bancho and IRC at the same time.** We'll add the ability to log in from both Bancho and IRC at the same time later.

### Setting up your IRC client
As we've said above, **some IRC clients might not work**, because we are using a very basic implementation of the IRC protocol.  
We've tested our IRC server with [HexChat](https://hexchat.github.io), a free and open source IRC client, and it turned out to work pretty well.  
If you want to connect from your browser, [Kiwi IRC](https://kiwiirc.com/) works well too.
<br>
Once you've installed an IRC client, you can connect to Ripple with these settings:  

- **Host:** `35.238.115.55`  
- **Port:** `6667`  
- **Nickname and username:** Your Atoka username. <u>If you have spaces in your name, replace them with underscores. For example,</u> `- Meme -` <u>becomes</u> `-_Meme_-`  
- **Password:** Your IRC token, you can get it [here](/irc).  

### I get "Nickname is already in use" when I try to log in!
Make sure you are not connected from Bancho or another IRC client. If it still doesn't work, retry in 2 minutes and it should work.

### And now?
If you have connected successfully to Atoka's IRC server, you can start chatting in public channels by typing `/JOIN #channel` (eg: `/JOIN #osu` to join the main channel). Private messages work too. If you find any issues or your IRC client doesn't work with Atoka, let us know and we'll do what we can.
