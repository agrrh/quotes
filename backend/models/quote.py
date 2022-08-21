from sqlalchemy import Column, Integer, String, DateTime
from sqlalchemy.sql import func

from database import Base

from datetime import datetime


class MQuote(Base):
    __tablename__ = "quotes"

    id = Column(Integer, primary_key=True, index=True)
    author = Column(String, index=True)
    content = Column(String)
    added_at = Column(DateTime(timezone=False), server_default=func.now())
    approved_at = Column(DateTime(timezone=False), default=datetime.fromtimestamp(0))
