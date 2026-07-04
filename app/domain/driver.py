from abc import ABC, abstractmethod
from dataclasses import dataclass
from app.domain.enums import Protocol


class IScannerDriver(ABC):
    ...

@dataclass
class ScanerDriver:
    '''ДРАЙВЕР - это контракт, правила, протокол'''
    def __init__(self):
        device_string: str 

        host: str
        port: int | None = None


        protocol: Protocol


        