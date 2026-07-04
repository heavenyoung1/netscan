from dataclasses import dataclass
from datetime import datetime
from pathlib import Path
from typing import List, Optional

from app.domain.enums import Status

@dataclass
class ScanJob:
    id: int
    status: Status
    created_at: datetime

    path: Path | None = None

