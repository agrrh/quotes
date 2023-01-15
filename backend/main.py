import os

from typing import List
from fastapi import Depends, FastAPI, HTTPException
from sqlalchemy.orm import Session
from fastapi.middleware.cors import CORSMiddleware

import crud
from database import Base
from schemas.quote import SQuoteCreate, SQuote
from database import SessionLocal, engine

Base.metadata.create_all(bind=engine)

app = FastAPI()

# CORS
origins = [
    "http://localhost",
    "http://localhost:5173",
    "http://localhost:8080",
    "https://quotes.agrrh.com",
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


# Dependency
def get_db() -> SessionLocal:
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()


@app.get("/")
def get_root() -> dict:
    return {
        "name": "quotes",
        "version": os.environ.get("APP_VERSION") or "develop",  # FIXME
    }


@app.post("/quotes/", response_model=SQuote)
def create_quote(quote: SQuoteCreate, db: Session = Depends(get_db)) -> SQuote:  # noqa: B008
    # TODO: Return code 201 and fix test case accordingly
    return crud.create_quote(db=db, quote=quote)


@app.get("/quotes/", response_model=List[SQuote])
def read_quotes(skip: int = 0, limit: int = 100, db: Session = Depends(get_db)) -> List[SQuote]:  # noqa: B008
    return crud.get_quotes(db, skip=skip, limit=limit)


@app.get("/quotes/{quote_id}", response_model=SQuote)
def read_quote(quote_id: int, db: Session = Depends(get_db)) -> SQuote:  # noqa: B008
    db_quote = crud.get_quote(db, quote_id=quote_id)
    if db_quote is None:
        raise HTTPException(status_code=404, detail="Quote not found")
    return db_quote
