
class ScannerError(Exception):
    pass


class ScannerNotFoundError(ScannerError):
    pass


class ScannerBusyError(ScannerError):
    pass


class ScannerDisconnectedError(ScannerError):
    pass