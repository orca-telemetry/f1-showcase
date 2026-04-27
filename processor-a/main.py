from orca_python import Processor

processor = Processor("Clean")


def start():
    # TODO: create and assign all the algorithms
    processor.Start()


def register():
    processor.Register()


if __name__ == "__main__":
    register()
    start()
