from typing import List, Any

from sqlalchemy.orm import Session

from models.quote import MQuote
from schemas.quote import SQuoteCreate, SQuote


def __get_quotes(db: Session) -> List[SQuote]:
    return db.query(MQuote)


def __paginate(query_result: object, skip: int = 0, limit: int = 20) -> List[Any]:
    return query_result.offset(skip).limit(limit).all()


def get_quote(db: Session, quote_id: int) -> SQuote:
    return __get_quotes(db).filter(MQuote.id == quote_id).first()


def get_quotes(db: Session, skip: int = 0, limit: int = 20) -> List[SQuote]:
    return __paginate(__get_quotes(db), skip=skip, limit=limit)


def get_quotes_by_author(db: Session, author: str, skip: int = 0, limit: int = 20) -> List[SQuote]:
    return __paginate(__get_quotes(db).filter(MQuote.author == author), skip=skip, limit=limit)


def create_quote(db: Session, quote: SQuoteCreate) -> SQuote:
    db_quote = MQuote(**quote.dict())
    db.add(db_quote)
    db.commit()
    db.refresh(db_quote)
    return db_quote
