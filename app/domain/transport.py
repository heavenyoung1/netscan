from abc import ABC, abstractmethod
from app.domain.device import ScannerDevice
import asyncio


class IScannerTransport(ABC):

    @abstractmethod
    async def discover(self) -> list[ScannerDevice]:
        ...

    @abstractmethod
    async def scan(self, device_id: str, args: dict) -> bytes:
        ...


class ScannerTransport(IScannerTransport):
    async def discover(self) -> list[ScannerDevice]:
        proc = await asyncio.create_subprocess_exec(
            'scanimage',
            '-L',
            stdout=asyncio.subprocess.PIPE,
            stderr=asyncio.subprocess.PIPE,
        )