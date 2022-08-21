from datetime import datetime

from pydantic import BaseModel


class SQuoteBase(BaseModel):
    author: str
    content: str


class SQuoteCreate(SQuoteBase):
    pass


class SQuote(SQuoteBase):
    id: str
    added_at: datetime
    approved_at: datetime

    class Config:
        orm_mode = True
