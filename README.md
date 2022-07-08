Dedoub
======

Tool to generate duplicated file list for a given folder

usage ./dedoub /path/to/inspect /tmp/result.json

result file looks like

```json
{
    "8bb588aebde9909b124650b452c7d7672ba6eecffed479fb3ca959df47a8093e": [
        {
            "Size": 175,
            "LastModified": "2022-07-09T00:34:38.594351665+02:00",
            "Filename": "HEAD",
            "Path": "/home/me/sources/dedoub/.git/logs/HEAD",
            "Checksum": "8bb588aebde9909b124650b452c7d7672ba6eecffed479fb3ca959df47a8093e"
        },
        {
            "Size": 175,
            "LastModified": "2022-07-09T00:34:38.594351665+02:00",
            "Filename": "main",
            "Path": "/home/me/sources/dedoub/.git/logs/refs/heads/main",
            "Checksum": "8bb588aebde9909b124650b452c7d7672ba6eecffed479fb3ca959df47a8093e"
        }
    ]
}
```
