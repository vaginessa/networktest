'''This needs to be a script as passing *.go to "go build" in the terminal on
Windows does not work.'''


import os
import subprocess
import sys


def detect_binary(binary):
    '''Windows requires a full path, so detect the full path of a binary by
    traversing $PATH.'''

    binary = '%s.exe' if os_is_windows() else binary
    delim = ';' if os_is_windows() else ':'

    paths = os.environ["PATH"].split(delim)
    for path in paths:
        if not os.path.exists(path):
            continue

        files = os.listdir(path)
        files = [file for file in files if file == binary]
        if files:
            return os.path.join(path, binary)

def invoke(args, cwd='.'):
    print("Invoking [cwd: %s] %s" % (cwd, args))
    proc = subprocess.Popen(
        args=args, cwd=cwd,
        stdout=subprocess.PIPE, stderr=subprocess.PIPE,
    )
    stdout, stderr = proc.communicate()
    return stdout.strip()

def os_is_windows():
    return sys.platform.startswith('win')


class Builder(object):
    version_module = os.path.join('src', 'app_version.go')

    def detect_version(self):
        git_exe = detect_binary(binary='git')
        args = [git_exe, 'describe']
        version = invoke(args=args)
        return version

    def set_version(self, version):
        version = version or '?'
        with open(self.version_module, 'wb') as f:
            f.write('package main\n\n\nconst appVersion = "%s"\n' % version)

    def build_on_windows(self):
        src_dir = "src"
        target = os.path.join("..", "bin", "havenet.exe")

        # Find all the sources
        files = os.listdir(src_dir)
        files.sort()
        files = [file for file in files if ".go" in file]
        files = [file for file in files if not "_test.go" in file]

        cwd = "src"
        executable = detect_binary(binary='go')
        args = [executable] + ["build"] + ["-o", target] + files

        invoke(args=args, cwd=cwd)

    def run(self):
        version = self.detect_version()
        self.set_version(version)

        # If we're not on windows we let the makefile handle the build
        if not os_is_windows():
            return

        self.build_on_windows()


if __name__ == '__main__':
    builder = Builder()
    builder.run()