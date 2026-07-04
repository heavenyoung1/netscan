from dataclasses import dataclass


@dataclass
class ScanerDevice:
    id: str         # "net:192.168.1.50:ricoh"
    name: str       # "RICOH Scanner"
    ip: str         # "192.168.1.50"
    protocol: str   # "sane-net"