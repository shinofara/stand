import sys
import platform
import commands
import json
import os
import stat

allow_oss = ["darwin", "linux"]
local_os = platform.system().lower()

if local_os not in allow_oss:
    mes = "%s is not allow" % local_os
    sys.exit(mes)


install_pkg_name = "stand_%s_amd64" % local_os

rowdata = commands.getoutput("curl -s https://api.github.com/repos/shinofara/stand/releases/latest")
data = json.loads(rowdata)
releases = data['assets']

for release in releases:
    if release['name'] == install_pkg_name:
        print "install %s" % release['browser_download_url']
        cmd = "curl -L %s -o /usr/local/bin/stand" % release['browser_download_url']
        check = commands.getoutput(cmd)
        print check
        os.chmod('/usr/local/bin/stand', 0755)

        check = commands.getoutput("which stand")
        print check
