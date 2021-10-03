
class FileSystem:

    def __init__(self):
        self.rootDir = Directory("/", None)
        self.cwd = "/"

    def cd(self, path):
        return NotImplemented

    def mkdir(self, path):
        return NotImplemented

    def touch(self, path, contents):
        return NotImplemented


class FSEntry:

    def __init__(self, name):
        self.name = name


class Directory(FSEntry):

    def __init__(self, name, parent=None):
        super.__init__(self, name)
        self.parent = parent
        self.contents = []

    def ls(self):
        return [str(f) for f in self.contents]


class File(FSEntry):

    def __init__(self, name, contents: str or bytearray):
        super.__init__(self, name)
        self.contents = contents
        self.size = len(contents)


if __name__ == "__main__":
    fs = FileSystem()
