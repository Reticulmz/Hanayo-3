---
title: "Installing the certificate manually"
old_id: 12
---
If you're having troubles connecting to Atoka or the switcher doesn't install the certificate properly, you can try installing it manually.

### Instructions
- First, download the certificate [by clicking here](https://atoka.pw/static/atoka.crt)
- Then, open **atoka.crt**
- Click **Install certificate...**
- Click **Next**
- Select **Place all certificates in the following store** (second option), then click **Browse...**
- A new window will pop up, select **Trusted root certification authorities** and click **Ok**
- Click **Next**
- Click **Finish**

### If the above didn't work try this
...you can try to remove all existing Atoka certificates and install the certificate again. Follow these steps:

- Press **Win+R**  
- Type `mmc certmgr.msc` in the run box and press **enter** to open the Certificate Manager  
- Select **Trusted root certification authorities** on the left  
- Select **Certificates** on the right  
- You should see **\*.ppy.sh** entries in the list. Select them, **right click** and click on **Delete**  
- Select all the positive options (Ok/Yes etc)  
- Restart the switcher, click on **Install Certificate**, then click **Connect to Atoka**
**If the above didn't fix it try running osu! as an administrator**.
