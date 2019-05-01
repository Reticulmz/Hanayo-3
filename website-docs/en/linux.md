---
title: "How to connect to Atoka (Linux)"
old_id: 14
---
This guide is only for connecting osu! to Atoka, and not setting the game itself up. You can follow [this guide](https://gist.github.com/Francesco149/a2f796683a4e5195458f4bb171d88eb0) to set the client up.

### 1a. Modifying the hosts file, the quick way
You can use this command to append the most up to date ripple redirections to your hosts file:
```
curl -s https://pastebin.com/raw/vNrZeVMz | python -c "from __future__ import print_function; import sys, json; j = json.load(sys.stdin); print('\n# Atoka\n' + '\n'.join('{} {}'.format(k, v) for k, v in {i: ' '.join([k for k, v in j.items() if v == i]) for _, i in j.items()}.items()))" | sudo tee -a /etc/hosts > /dev/null
```

### 1b. Modifying the hosts file manually
Alternatively, you can edit your hosts file manually. To do so, run `nano /etc/hosts` as root/with sudo.

When you've got it open, paste the following at the bottom:

```
35.238.115.55 osu.ppy.sh c.ppy.sh c1.ppy.sh c2.ppy.sh c3.ppy.sh c4.ppy.sh c5.ppy.sh c6.ppy.sh ce.ppy.sh a.ppy.sh s.ppy.sh i.ppy.sh
35.238.115.55 bm6.ppy.sh
```
**CTRL+X** and then **Enter** to save the file.

### 2. Installing the certificate
Download the certificate by clicking [*here*](https://atoka.pw/static/atoka.crt)

Open the Internet Explorer configuration by running `wine control`.

Double click the *Internet Settings* icon, navigate to the *Content* tab, then click the *Certificates...* button.

Click on *Import*, then *Next*.

Click *Browse...* then select the Atoka certificate.

Click *Next*.

Select *Place all certificates in the following store*, and click *Browse*.

Select **Trusted Root Certification Authorities**, and click *Ok*.

Click *Next*, *Finish*.

You should get a message saying **The import was successful**.


After that is done, you can start the client up, and log in with your Atoka credentials.
