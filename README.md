# backup.sh

Features:

* Create encrypted backups avoiding duplicates.
* Distributed backup, ability to store files in multiple locations.
* keep history, snapshots, versioning.
* Share files or full backup via recipies.
* Hybrid cryptosystem


Available commands:

- src
- dst
- restore
- find
- list
- snapshot
- get




mkfile 1G /tmp/1GB.raw


RSA can only encrypt data blocks that are shorter than the key length so what you normally do is

1. Generate a random key of the correct length required for AES (or similar).
2. Encrypt your data using AES or similar using that key
3. Encrypt the random key using your RSA key

To decrypt

1. Decrypt the AES key using your RSA key.
2. Decrypt the data using that AES key
